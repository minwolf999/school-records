package fetch

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// These hundle function send remove a link between a student and a competence and send a boolean
func RemoveLink(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id_teacher := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id_teacher)
	if err != nil {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id_teacher, err))
		file.Close()
		return
	}

	if r.Method != http.MethodPost {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()
		return
	}

	// Get the datas from the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	defer r.Body.Close()

	var datas map[string]string
	err = json.Unmarshal(body, &datas)
	if err != nil {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Write someone comming to the removeLink route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, datas))
	file.Close()

	// Stock the student id from the map to a variable
	studentId, isOk := datas["studentId"]
	if !isOk {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "forgotten student Id"))
		file.Close()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(false)
		return
	}

	// Stock the competence id from the map to a variable
	competenceId, isOk := datas["competenceId"]
	if !isOk {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "forgotten competence Id"))
		file.Close()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(false)
		return
	}

	// Delente the link between the student and the competence
	err = controller.Delete("linkCompetenceEleve", []string{"studentId", "competenceId"}, studentId, competenceId)
	if err != nil {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(false)
		return
	}

	// Send as a JSON to the user the result of the function
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(true)
	if err != nil {
		// Write someone comming to the removeLink route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - RemoveLink fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
