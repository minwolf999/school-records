package handler

import (
	"dossier_scolaire/cookie"
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// These hundle function the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the login route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /login route\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "login.html", nil)
		return
	}

	// Get the value from the form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Write someone comming to the login route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /login route - datas: [%s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, username))
	file.Close()

	// Verify if there isn't an empty field
	if username == "" || password == "" {
		// Write someone comming to the login route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /login route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "empty field"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "login.html", "Il y a une zone vide !")
		return
	}

	// Get the teacher datas by his username
	teacher, err := controller.SelectTeacher("teachers", []string{"username"}, username)
	if err != nil {
		// Write someone comming to the login route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /login route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, err))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "login.html", "Nom d'utilisateur ou mot de passe incorrect !")
		return
	}

	// Look if the datas is empty or if the password is incorrect
	err = bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password))
	if teacher.Id == "" || err != nil {
		if err != nil {
			// Write someone comming to the login route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /login route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, err))
			file.Close()
		} else {
			// Write someone comming to the login route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /login route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "no teacher match"))
			file.Close()
		}

		structure.Tpl.ExecuteTemplate(w, "login.html", "Nom d'utilisateur ou mot de passe incorrect !")
		return
	}

	// Add a cookie and redirect to th home page
	cookie.AddCookies(w, teacher.Id)
	http.Redirect(w, r, "/saisir", http.StatusSeeOther)
}
