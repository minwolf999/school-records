package main

import (
	initdatabase "dossier_scolaire/database/initDatabase"
	"dossier_scolaire/routes"
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Listen for the CTRL+c command to save the BDD and close the program
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		// Write the server is closing in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("[%s] Closing Server", time.Now().Format("2006-01-02 15:04:05")))
		file.Close()

		fmt.Println()
		initdatabase.SaveBDD()
		os.Exit(0)
	}()

	// The setInterval function is called asynchronally and save the BDD each minute
	go utility.SetInterval(24*time.Hour, initdatabase.SaveBDD)

	// Set in the global variable the port used by the application and the local ipv4 address we need to used for access to the web site
	structure.App = utility.NewApplication("8080")

	// Parse all the template files in the server (html, css, images)
	structure.Tpl = template.Must(template.New("").ParseGlob("template/html/*.html"))
	fs := structure.JustFilesFilesystem{Fs: http.Dir("template")}

	http.Handle("/template/", http.StripPrefix("/template", http.FileServer(fs)))
}

func main() {
	fmt.Println("Server started on http://" + structure.App.Ip + ":" + structure.App.Port)
	fmt.Println("To close the server, use CTRL+c")

	// Set all the accessible routes for the site
	routes.Routes()

	// Open in the default browser when the server is starting
	utility.Open("http://" + structure.App.Ip + ":" + structure.App.Port)

	// Open the server
	if err := http.ListenAndServe(":"+structure.App.Port, nil); err != nil {
		fmt.Println(err)
	}
}
