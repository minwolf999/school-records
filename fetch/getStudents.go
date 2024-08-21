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

// These hundle function send a slice of student who has a given year and who has been created by a user in the BDD
func GetStudents(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id_teacher := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id_teacher)
	if err != nil {
		// Write someone comming to the getStudents route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id_teacher, err))
		file.Close()
		return
	}

	if r.Method != http.MethodPost {
		// Write someone comming to the getStudents route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()
		return
	}

	// Get the year from the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Write someone comming to the getStudents route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
	defer r.Body.Close()

	var datas map[string]string
	err = json.Unmarshal(body, &datas)
	if err != nil {
		// Write someone comming to the getStudents route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Write someone comming to the getStudents route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, datas))
	file.Close()

	// Get the teacher id in the Context and get all the datas with this id
	students, err := controller.SelectStudents("students", []string{"teacherId", "year"}, id_teacher, datas["year"])
	if err != nil {
		// Write someone comming to the getStudents route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Get all the categorie assigned to a student
	for i := range students {
		students[i].Competences, err = controller.SelectCompetenceLinked("linkCompetenceEleve", []string{"studentId"}, students[i].Id)
		if err != nil {
			// Write someone comming to the getStudents route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
			return
		}

		for y := range students[i].Competences {
			tmp, err := controller.SelectCompetences("competences", []string{"id"}, students[i].Competences[y].Id)
			if err != nil {
				// Write someone comming to the getStudents route in the log file
				file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
				file.Close()
				return
			}

			students[i].Competences[y] = tmp[0]
		}
	}

	students = utility.FilterByClassAndName(students)

	// Send as a JSON to the user
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		// Write someone comming to the getStudents route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - GetStudents fetch - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
