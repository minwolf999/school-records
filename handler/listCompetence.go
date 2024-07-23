package handler

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
	"fmt"
	"net/http"
	"os"
	"time"
)

// These hundle function the listCompetence page
func ListCompetenceHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the listCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Write someone comming to the listCompetence route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listCompetence route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
	file.Close()

	// Get all the categories and remove the categorie not accessible to the teacher
	categories, err := controller.SelectCategories("categories", []string{})
	if err != nil {
		// Write someone comming to the listCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	categories = utility.FilterCategoriesForTeacher(categories, teacher)

	// Get all the students of a teacher and Get all the different years
	students, err := controller.SelectStudents("students", []string{"teacherId"}, teacher.Id)
	if err != nil {
		// Write someone comming to the listCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	yearList := utility.YearListFromStudent(students)

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		err = structure.Tpl.ExecuteTemplate(w, "listCompetence.html", structure.Result{Categories: categories, YearList: yearList})
		if err != nil {
			// Write someone comming to the listCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
			return
		}
		return
	}
}
