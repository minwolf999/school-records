package handler

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"fmt"
	"net/http"
	"os"
	"time"
)

// These hundle function the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the home route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /home route - %s - error %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Write someone comming to the home route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /home route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
	file.Close()

	// Execute the template
	err = structure.Tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		// Write someone comming to the home route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /home route - %s - error %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}
}
