package XML

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

// These hundle function remove a competence create by a user in the BDD
func DeleteCompetence(w http.ResponseWriter, r *http.Request) {
	// Get the teacher id in the Context and get all the datas with this id
	teacherId := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, teacherId)
	if err != nil {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacherId, err))
		file.Close()
		return
	}

	if r.Method != http.MethodPost {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()
		return
	}

	// Get the competence id from the body of the request
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	var id string
	err = json.Unmarshal(body, &id)
	if err != nil {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacherId, err))
		file.Close()
		return
	}

	// Write someone comming to the deleteCompetence route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, id))
	file.Close()

	// Get the informations of the competences to remove the image in the image folder
	competence, err := controller.SelectCompetences("competences", []string{"id", "teacherId"}, id, teacher.Id)
	if err != nil {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacherId, err))
		file.Close()
		return
	}

	err = os.Remove(competence[0].ImagePath[1:])
	if err != nil {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacherId, err))
		file.Close()
		return
	}
	
	// Remove the competence from the BDD
	err = controller.Delete("competences", []string{"id", "teacherId"}, id, teacher.Id)
	if err != nil {
		// Write someone comming to the deleteCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - DeleteCompetence XML - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacherId, err))
		file.Close()
		return
	}
}
