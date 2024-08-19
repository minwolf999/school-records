package utility

import (
	"dossier_scolaire/structure"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// These function initialise the variable with the server setting
func NewApplication(port string) structure.Application {
	app := new(structure.Application)

	app.Port = port
	app.Ip = GetOutboundIP()
	app.LogPath = "/log/" + time.Now().Format("2006-01-02__15-04-05") + ".log"

	if runtime.GOOS == "windows" {
		app.LogPath = strings.ReplaceAll(app.LogPath, "/", "\\")
	}

	file, _ := os.Create(app.LogPath[1:])
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s Starting Server\n", time.Now().Format("2006-01-02 15:04:05")))

	return *app
}
