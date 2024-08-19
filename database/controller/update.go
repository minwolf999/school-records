package controller

import (
	initdatabase "dossier_scolaire/database/initDatabase"
	"errors"
	"fmt"
)

// These function update the datas of a line of a table of the BDD
//
// These function takes at minimum 4 arguments:
//
// - the tab name
//
// - a map with the information and the rows to update (rows name has key and rows new value has value)
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These function return an error or nil
func Update(tab string, toUpdate map[string]string, rows []string, datas ...string) error {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	modify := ""
	count := 0
	for i, v := range toUpdate {
		count++

		modify += fmt.Sprintf("%s=\"%s\"", i, v)

		if count < len(toUpdate) {
			modify += ", "
		}
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("UPDATE `%s` SET %s WHERE %s", tab, modify, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("UPDATE `%s` SET %s", tab, modify)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return err
	}

	// Execute the SQL request
	_, err = stmt.Exec()
	return err
}
