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

// These hundle function send a slice of subCategories who has a parent categorie
func GetSubCategorie(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id_teacher := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id_teacher)
	if err != nil {
		// Write someone comming to the getSubCategorie route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id_teacher, err))
		file.Close()
		return
	}

	if r.Method != http.MethodPost {
		// Write someone comming to the getSubCategorie route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()
		return
	}

	// Get the datas from the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Write someone comming to the getSubCategorie route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	defer r.Body.Close()

	var categorieId string
	err = json.Unmarshal(body, &categorieId)
	if err != nil {
		// Write someone comming to the getSubCategorie route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Write someone comming to the getSubCategorie route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s - dats: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, categorieId))
	file.Close()

	// Get the subCategories from the BDD
	subCategories, err := controller.SelectSubCategories("subCategories", []string{"idCategorie"}, categorieId)
	if err != nil {
		// Write someone comming to the getSubCategorie route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Send as a JSON to the user
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(subCategories)
	if err != nil {
		// Write someone comming to the getSubCategorie route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetSubCategorie fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
