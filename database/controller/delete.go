package controller

import (
	initdatabase "dossier_scolaire/database/initDatabase"
	"errors"
	"fmt"
)

// These function remove a line of a table of the BDD
//
// These function takes at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These function return an error or nil
func Delete(tab string, rows []string, datas ...string) error {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string and prepare them
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("DELETE FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("DELETE FROM `%s`", tab)
	}

	stmt, err := db.Prepare(request)
	if err != nil {
		return err
	}

	// Execute the SQL request
	_, err = stmt.Exec()
	return err
}
