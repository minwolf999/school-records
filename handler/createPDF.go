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

// These hundle function the createPDF page
func CreatePDF(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the createPDF route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Get the id of the student selected to create the pdf
	ck, err := r.Cookie("studentId")
	if err != nil {
		// Write someone comming to the createPDF route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		http.Redirect(w, r, "/observer", http.StatusSeeOther)
		return
	}

	studentId := ck.Value

	// Write someone comming to the createPDF route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, studentId))
	file.Close()

	// Get the student informations
	students, err := controller.SelectStudents("students", []string{"id", "teacherId"}, studentId, teacher.Id)
	if err != nil || len(students) == 0 {
		if err != nil {
			// Write someone comming to the createPDF route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
			file.Close()
		} else {
			// Write someone comming to the createPDF route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "student didn't exist"))
			file.Close()
		}

		http.Redirect(w, r, "/observer", http.StatusSeeOther)
		return
	}

	student := students[0]

	// Get the id of the competences linked to the student
	student.Competences, err = controller.SelectCompetenceLinked("linkCompetenceEleve", []string{"studentId"}, student.Id)
	if err != nil {
		// Write someone comming to the createPDF route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		http.Redirect(w, r, "/observer", http.StatusSeeOther)
		return
	}

	// Get the detail of each competence
	for y := range student.Competences {
		competences, err := controller.SelectCompetences("competences", []string{"id"}, student.Competences[y].Id)
		if err != nil {
			// Write someone comming to the createPDF route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
			file.Close()

			http.Redirect(w, r, "/observer", http.StatusSeeOther)
			return
		}

		student.Competences[y] = competences[0]
	}

	// Filter the competences
	student.Competences = utility.FilterCompetences(student.Competences)

	// generate the pdf
	pdf := utility.GeneratePDF(student, teacher)
	err = pdf.Output(w)
	if err != nil {
		// Write someone comming to the createPDF route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/createPDF route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
	}
}
