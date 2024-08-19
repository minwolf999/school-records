package fetch

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// These hundle function send a slice of competence who has a parent categorie
func GetStudentsByCompetence(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id_teacher := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id_teacher)
	if err != nil {
		// Write someone comming to the getStudentsyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id_teacher, err))
		file.Close()
		return
	}

	if r.Method != http.MethodPost {
		// Write someone comming to the getStudentsyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()
		return
	}

	// Get the datas from the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Write someone comming to the getStudentsyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	defer r.Body.Close()

	var datas []string
	err = json.Unmarshal(body, &datas)
	if err != nil {
		// Write someone comming to the getStudentsyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Write someone comming to the getStudentsyCompetence route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, datas))
	file.Close()

	// Get the all competences from a categorie
	competences, err := controller.SelectCompetences("competences", []string{"categorieId"}, datas[0])
	if err != nil {
		// Write someone comming to the getStudentsyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Get all students linked to the competence
	for i := range competences {
		competences[i].Students, err = controller.SelectStudentLinked("linkCompetenceEleve", []string{"competenceId"}, competences[i].Id)
		if err != nil {
			// Write someone comming to the getStudentsyCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
			return
		}

		for y := range competences[i].Students {
			tmp, err := controller.SelectStudents("students", []string{"id", "year"}, competences[i].Students[y].Id, datas[1])
			if err != nil {
				// Write someone comming to the getStudentsyCompetence route in the log file
				file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
				file.Close()
				return
			}

			competences[i].Students[y] = tmp[0]
		}
	}

	competences = utility.FilterCompetencesByCategorieAndName(competences)

	// Send as a JSON to the user
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(competences)
	if err != nil {
		// Write someone comming to the getStudentsyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudentsByCompetence fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
