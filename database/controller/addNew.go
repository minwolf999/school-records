package controller

import (
	initdatabase "dossier_scolaire/database/initDatabase"
	"errors"
	"fmt"
	"strings"
)

// These function add a new line in a table of the BDD
//
// These function takes at minimum 2 arguments:
//
// - the tab name
// 
// - the information to add in the table
// 
// These function return an error or nil
func AddNew(tab string, datas ...string) error {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return err
	}

	// If the datas variable have a length of 0 then we got not enought arguments to call the INSERT in the SQL request
	if len(datas) == 0 {
		return errors.New("there isn't data to add in the BDD")
	}

	// Create the request in a string and prepare them
	request := fmt.Sprintf("INSERT INTO `%s` VALUES(\"%s\")", tab, strings.Join(datas, "\", \""))
	stmt, err := db.Prepare(request)
	if err != nil {
		return err
	}

	// Execute the SQL request
	_, err = stmt.Exec()
	return err
}
