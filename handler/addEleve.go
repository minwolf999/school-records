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

// These hundle function the addEleve page
func AddEleve(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the addEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the addEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | ")})
		return
	}

	// Get the value from the form
	id = utility.NewUUID()
	class := r.FormValue("class")
	name := r.FormValue("name")
	year := r.FormValue("year")

	// Write someone comming to the addEleve route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s - datas: [%s, %s, %s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, name, class, year))
	file.Close()

	// Verify if there isn't an empty field
	if name == "" || year == "" || (class != "Ps" && class != "Ms" && class != "Gs") {
		// Write someone comming to the addEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "empty field"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Error: "Il y a un champ vide !"})
		return
	}

	students, _ := controller.SelectStudents("students", []string{"name", "class", "year", "teacherId"}, name, class, year, teacher.Id)
	if len(students) != 0 {
		// Write someone comming to the addEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "Il y a déjà un élève portant ce nom dans cette classe"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Error: "Il y a déjà un élève portant ce nom dans cette classe"})
		return
	}

	// Add the new student in the BDD
	err = controller.AddNew("students", id, name, class, year, teacher.Id)
	if err != nil {
		// Write someone comming to the addEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Error: err.Error()})
		return
	}

	// Execute the template with a success message
	err = structure.Tpl.ExecuteTemplate(w, "addEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Success: fmt.Sprintf("L'élève `%s` a été créé avec succès!", name)})
	if err != nil {
		// Write someone comming to the addEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}
}
