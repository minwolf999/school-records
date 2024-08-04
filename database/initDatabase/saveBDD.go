package initdatabase

import (
	"fmt"
	"io"
	"os"
)

// These function create a save of the BDD in the folder backup
func SaveBDD() {
	src := "./database/database.sqlite3"
	dst := "./database/backup/backup.sqlite3"

	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println(1, err)
		return
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(2, err)
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println(3, err)
	}

	fmt.Println("Backup success")
}
