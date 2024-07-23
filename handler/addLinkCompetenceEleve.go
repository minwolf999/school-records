package handler

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// These hundle function the addLinkCompetenceEleve page
func AddLinkCompetenceEleve(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the addLinCompetenceEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Get all the students of a teacher and Get all the different years
	studentsForYear, err := controller.SelectStudents("students", []string{"teacherId"}, teacher.Id)
	if err != nil {
		// Write someone comming to the addLinCompetenceEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	yearList := utility.YearListFromStudent(studentsForYear)

	// Get all the categories and remove the categorie not accessible to the teacher
	categories, err := controller.SelectCategories("categories", []string{})
	if err != nil {
		// Write someone comming to the addLinCompetenceEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	categories = utility.FilterCategoriesForTeacher(categories, teacher)

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the addLinCompetenceEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addLinkCompetenceEleve.html", structure.Result{
			YearList:   yearList,
			Categories: categories,
		})
		return
	}

	// Get the value from the form
	r.ParseForm()
	students := r.Form["students"]
	competences := r.FormValue("competences")

	// Write someone comming to the addLinCompetenceEleve route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - datas: [%s, %s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, strings.Join(students, ", "), competences))
	file.Close()

	// Verify if there isn't an empty field
	if len(students) == 0 || competences == "" {
		// Write someone comming to the addLinCompetenceEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "empty field"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addLinkCompetenceEleve.html", structure.Result{
			Error:      "Aucun élève ou compétence n'ont été sélectionné !",
			YearList:   yearList,
			Categories: categories,
		})
		return
	}

	// Add a link between all students selected and a competence
	for _, v := range students {
		studentsVerification, err := controller.SelectStudents("students", []string{"teacherId", "id"}, teacher.Id, v)
		if err != nil {
			// Write someone comming to the addLinCompetenceEleve route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
			return
		}

		if studentsVerification[0].Id == "" {
			// Write someone comming to the addLinCompetenceEleve route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "student invalid"))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "addLinkCompetenceEleve.html", structure.Result{
				Error:      fmt.Sprintf("Un étudiant n'appartient pas à %s", teacher.Username),
				YearList:   yearList,
				Categories: categories,
			})
			return
		}

		// Verify if there is already a link between a student and a competence
		tmp, _ := controller.SelectCompetenceLinked("linkCompetenceEleve", []string{"competenceId", "studentId"}, competences, v)
		if len(tmp) != 0 {
			continue
		}

		// Add the link in the BDD
		err = controller.AddNew("linkCompetenceEleve", utility.NewUUID(), studentsVerification[0].Id, competences)
		if err != nil {
			// Write someone comming to the addLinCompetenceEleve route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "addLinkCompetenceEleve.html", structure.Result{
				Error:      err.Error(),
				YearList:   yearList,
				Categories: categories,
			})
			return
		}
	}

	// Execute the Template with a success message
	err = structure.Tpl.ExecuteTemplate(w, "addLinkCompetenceEleve.html", structure.Result{
		Success:    "La compétence a bien été attribuer aux élèves sélectionnées",
		YearList:   yearList,
		Categories: categories,
	})

	if err != nil {
		// Write someone comming to the addLinCompetenceEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addLinCompetenceEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
