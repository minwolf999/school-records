package initdatabase

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// These function create a save of the BDD in the folder backup
func SaveBDD() {
	command := "cp"
	dbpath := "./database/database.sqlite3"
	bakpath := "./database/backup/backup.sqlite3"

	// If the OS is on windows we modify the bash command to make it work for Windows
	os := runtime.GOOS
	if os == "windows" {
		command = "copy"
		dbpath = strings.ReplaceAll(dbpath, "/", "\\")
		bakpath = strings.ReplaceAll(bakpath, "/", "\\")
	}

	// Execute the command
	cmd := exec.Command(command, dbpath, bakpath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Errorf("dbBackup failed : %s : %v", string(out), err))
		return
	}

	fmt.Println("Backup success")
}
