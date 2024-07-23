package handler

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// These hundle function the modifyEleve page
func ModifyEleve(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Get the id of the student selected to modify these informations
	ck, err := r.Cookie("studentId")
	if err != nil {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "no student match"))
		file.Close()

		http.Redirect(w, r, "/saisir/listEleve", http.StatusSeeOther)
		return
	}
	studentId := ck.Value

	// Get the detail of the student selected
	students, err := controller.SelectStudents("students", []string{"id", "teacherId"}, studentId, teacher.Id)
	if err != nil {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "modifyEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Student: students[0]})
		return
	}

	// Get the value from the form
	class := r.FormValue("class")
	name := r.FormValue("name")
	year := r.FormValue("year")

	// Write someone comming to the modifyEleve route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - datas: [%s, %s, %s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, name, class, year))
	file.Close()

	// Verify if there isn't an empty field
	if name == "" || year == "" || (class != "Ps" && class != "Ms" && class != "Gs") {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "empty field"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Student: students[0], Error: "Il y a un champ vide !"})
		return
	}

	// Update the informations of the students
	err = controller.Update("students", map[string]string{"name": name, "class": class, "year": year}, []string{"id", "teacherId"}, studentId, teacher.Id)
	if err != nil {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "modifyEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Student: students[0], Error: err.Error()})
		return
	}

	students[0].Name = name
	students[0].Class = class
	students[0].Year = year

	// Execute the template with a success message
	err = structure.Tpl.ExecuteTemplate(w, "modifyEleve.html", structure.Result{Class: strings.Split(teacher.Class, " | "), Student: students[0], Success: fmt.Sprintf("L'élève `%s` a été mis à jour avec succès!", name)})
	if err != nil {
		// Write someone comming to the modifyEleve route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyEleve route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}
}
