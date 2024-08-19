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

// These hundle function send a slice of competence who has a parent categorie and who has been created by a user in the BDD
func GetCompetences(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the getCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	if r.Method != http.MethodPost {
		// Write someone comming to the getCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()
		return
	}

	// Get the categorie id from the body of the request
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		// Write someone comming to the getCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	var datas map[string]string
	err = json.Unmarshal(body, &datas)
	if err != nil {
		// Write someone comming to the getCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Write someone comming to the getCompetence route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, datas))
	file.Close()

	var competences []structure.Competence
	// Get the competences who his link to this categorie id
	_, isOk := datas["categories"]
	if isOk {
		competences, err = controller.SelectCompetences("competences", []string{"categorieId", "teacherId"}, datas["categories"], teacher.Id)
		if err != nil {
			// Write someone comming to the getCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
			return
		}
	}

	_, isOk = datas["subcategories"]
	if isOk {
		competences, err = controller.SelectCompetences("competences", []string{"subCategorieId"}, datas["subcategories"])
		if err != nil {
			// Write someone comming to the getCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
			return
		}
	}

	// Send as a JSON to the user
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(competences)
	if err != nil {
		// Write someone comming to the getCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
