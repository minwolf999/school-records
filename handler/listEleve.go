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

// These hundle function the listEleve page
func ListEleve(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the listEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Write someone comming to the listEleve route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listEleve route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
	file.Close()

	// Get all the students of a teacher and Get all the different years
	students, err := controller.SelectStudents("students", []string{"teacherId"}, teacher.Id)
	if err != nil {
		// Write someone comming to the listEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}
	yearList := utility.YearListFromStudent(students)

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		err = structure.Tpl.ExecuteTemplate(w, "listEleve.html", yearList)
		if err != nil {
			// Write someone comming to the listEleve route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/listEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
			file.Close()
			return
		}
		return
	}
}
