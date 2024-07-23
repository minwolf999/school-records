package utility

import (
	"dossier_scolaire/structure"
	"fmt"
	"os"
	"time"
)

// These function initialise the variable with the server setting
func NewApplication(port string) structure.Application {
	app := new(structure.Application)

	app.Port = port
	app.Ip = GetOutboundIP()
	app.LogPath = "/template/log/" + time.Now().Format("2006-01-02_15:04:05") + ".log"

	file, _ := os.Create(app.LogPath[1:])
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s Starting Server\n", time.Now().Format("2006-01-02 15:04:05")))

	return *app
}
