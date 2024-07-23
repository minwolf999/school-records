package handler

import (
	"dossier_scolaire/cookie"
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// These hundle function the register page
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the register route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /register route\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "register.html", nil)
		return
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Get the value from the form
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")
	class := strings.Join(r.Form["class"], " | ")
	key := utility.GenerateString(43, charset)

	id := utility.NewUUID()
	cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	// Write someone comming to the register route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /register route - datas: [%s, %s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, username, class))
	file.Close()

	// Verify if there isn't an empty field
	if username == "" || password == "" || confirm == "" || len(r.Form["class"]) == 0 {
		// Write someone comming to the register route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /register route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "empty field"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "register.html", "Il y a une zone vide !")
		return
	}

	// Verify if password and the confirmation of the password match
	if password != confirm {
		// Write someone comming to the register route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /register route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "password and confirm password don't match"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "register.html", "Le mot de passe et la confirmation ne correspondent pas !")
		return
	}

	// Verify if there isn't to much class selected
	if len(r.Form["class"]) > 2 {
		// Write someone comming to the register route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /register route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "to many class selected"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "register.html", "Trop de classes ont été sélectionné !")
		return
	}

	// Add a new teacher to the BDD
	err := controller.AddNew("teachers", id, username, string(cryptPassword), class, "", key)
	if err != nil {
		// Write someone comming to the register route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /register route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, err))
		file.Close()

		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			structure.Tpl.ExecuteTemplate(w, "register.html", "Le nom d'utilisateur voulu est déjà utilisé par un autre utilisateur")
		} else {
			structure.Tpl.ExecuteTemplate(w, "register.html", err.Error())
		}

		return
	}

	// Set a cookie and redirect to the home page
	cookie.AddCookies(w, id)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
