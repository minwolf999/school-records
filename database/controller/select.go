package controller

import (
	initdatabase "dossier_scolaire/database/initDatabase"
	"dossier_scolaire/structure"
	"errors"
	"fmt"
)

// These function take the datas of a table of the BDD
//
// These function takes at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These function return an user and an error
func SelectTeacher(tab string, rows []string, datas ...string) (structure.Teacher, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return structure.Teacher{}, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return structure.Teacher{}, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return structure.Teacher{}, err
	}

	// Execute the SQL request and fill the variable user
	var user structure.Teacher
	return user, stmt.QueryRow().Scan(&user.Id, &user.Username, &user.Password, &user.Class, &user.SigningUpPath, &user.Key)
}

// These functions take the datas of a table of the BDD
//
// These functions take at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These functions return a slice of student and an error
func SelectStudents(tab string, rows []string, datas ...string) ([]structure.Student, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return nil, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return nil, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return nil, err
	}

	// Execute the SQL request
	data, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Fill the students slice with the valut get from the SQL request
	var students []structure.Student
	for data.Next() {
		var student structure.Student
		err = data.Scan(&student.Id, &student.Name, &student.Class, &student.Year, &student.Teacher.Id)
		if err != nil {
			return nil, err
		}

		// Fill the teacher zone of the student structure
		student.Teacher, err = SelectTeacher("teachers", []string{"id"}, student.Teacher.Id)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

// These functions take the datas of a table of the BDD
//
// These functions take at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These functions return a slice of categorie and an error
func SelectCategories(tab string, rows []string, datas ...string) ([]structure.Categorie, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return nil, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return nil, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return nil, err
	}

	// Execute the SQL request
	data, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Fill the categories slice with the valut get from the SQL request
	var categories []structure.Categorie
	for data.Next() {
		var categorie structure.Categorie
		err = data.Scan(&categorie.Id, &categorie.Name, &categorie.Class)
		if err != nil {
			return nil, err
		}

		// Fill the subCategorie zone of the categorie structure
		tmp, err := SelectSubCategories("subCategories", []string{"idCategorie"}, categorie.Id)
		if err != nil {
			return nil, err
		}

		categorie.SubCategories = tmp

		categories = append(categories, categorie)
	}

	return categories, nil
}

// These functions take the datas of a table of the BDD
//
// These functions take at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These functions return a slice of competence and an error
func SelectCompetences(tab string, rows []string, datas ...string) ([]structure.Competence, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return nil, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return nil, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return nil, err
	}

	// Execute the SQL request
	data, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Fill the competences slice with the valut get from the SQL request
	var competences []structure.Competence
	for data.Next() {
		var competence structure.Competence
		var useles string
		err = data.Scan(&competence.Id, &competence.Name, &competence.ImagePath, &competence.Categorie.Id, &competence.SubCategorie.Id, &useles)
		if err != nil {
			return nil, err
		}

		// Fill the categorie zone of the competence structure
		tmp, err := SelectCategories("categories", []string{"id"}, competence.Categorie.Id)
		if err != nil {
			return nil, err
		}

		if len(tmp) > 0 {
			competence.Categorie = tmp[0]
		}

		tmp2, err := SelectSubCategories("subCategories", []string{"id"}, competence.SubCategorie.Id)
		if err != nil {
			return nil, err
		}

		if len(tmp2) > 0 {
			competence.SubCategorie = tmp2[0]
		}

		competences = append(competences, competence)
	}

	return competences, nil
}

// These functions take the datas of a table of the BDD
//
// These functions take at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These functions return a slice of competence and an error
func SelectCompetenceLinked(tab string, rows []string, datas ...string) ([]structure.Competence, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return nil, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return nil, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return nil, err
	}

	// Execute the SQL request
	data, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Fill the competences slice with the valut get from the SQL request
	var competences []structure.Competence
	for data.Next() {
		var competence structure.Competence
		var useless string
		err = data.Scan(&useless, &useless, &competence.Id)

		if err != nil {
			return nil, err
		}

		competences = append(competences, competence)
	}

	return competences, nil
}

// These functions take the datas of a table of the BDD
//
// These functions take at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These functions return a slice of student and an error
func SelectStudentLinked(tab string, rows []string, datas ...string) ([]structure.Student, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return nil, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return nil, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return nil, err
	}

	// Execute the SQL request
	data, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Fill the students slice with the valut get from the SQL request
	var students []structure.Student
	for data.Next() {
		var student structure.Student
		var useless string
		err = data.Scan(&useless, &student.Id, &useless)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

// These functions take the datas of a table of the BDD
//
// These functions take at minimum 3 arguments:
//
// - the tab name
//
// - a slice with the name of the rows
//
// - the information who the table of rows need to contain
//
// These functions return a slice of subcategorie and an error
func SelectSubCategories(tab string, rows []string, datas ...string) ([]structure.SubCategorie, error) {
	// Open a connexion to the BDD
	db, err := initdatabase.OpenBDD()
	if err != nil {
		return nil, err
	}

	// If the datas variable haven't the same length as rows variable then we got not enought arguments to call the INSERT in the SQL request
	if len(rows) != len(datas) {
		return nil, errors.New("there is not the same amount of rows and datas")
	}

	// Create the request in a string
	where := ""
	for i := 0; i < len(rows); i++ {
		where += fmt.Sprintf("%s='%s' AND ", rows[i], datas[i])
	}

	request := ""
	if len(rows) > 0 {
		request = fmt.Sprintf("SELECT * FROM `%s` WHERE %s", tab, where[0:len(where)-5])
	} else {
		request = fmt.Sprintf("SELECT * FROM `%s`", tab)
	}

	// Prepare the SQL request
	stmt, err := db.Prepare(request)
	if err != nil {
		return nil, err
	}

	// Execute the SQL request
	data, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Fill the subCategories slice with the valut get from the SQL request
	var subCategories []structure.SubCategorie
	for data.Next() {
		var subCategorie structure.SubCategorie
		err = data.Scan(&subCategorie.Id, &subCategorie.Name, &subCategorie.Class, &subCategorie.Categorie.Id)

		if err != nil {
			return nil, err
		}

		subCategories = append(subCategories, subCategorie)
	}

	return subCategories, nil
}
