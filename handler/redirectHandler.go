package handler

import (
	"dossier_scolaire/structure"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// These hundle function redirect to the good page
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/saisir/") {
		// Write someone comming to the redirect route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - / Route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "redirect to saisir route"))
		file.Close()

		SaisirHandler(w, r)
		return
	}

	// Write someone comming to the redirect route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - / Route - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, "redirect to login route"))
	file.Close()

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
