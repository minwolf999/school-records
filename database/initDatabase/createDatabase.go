package initdatabase

import (
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
)

var categories = []structure.Categorie{
	{
		Id:    utility.NewUUID(),
		Name:  "Mobiliser le langage dans toutes ses dimensions",
		Class: "Ps | Ms | Gs",
		SubCategories: []structure.SubCategorie{
			{
				Id: "1",
			},
		},
	},

	{
		Id:    utility.NewUUID(),
		Name:  "Agir, s'exprimer,comprendre à travers l'éducation physique",
		Class: "Ps | Ms | Gs",
	},

	{
		Id:    utility.NewUUID(),
		Name:  "Agir, s'exprimer, comprendre à travers l'éducation artistique",
		Class: "Ps | Ms | Gs",
		SubCategories: []structure.SubCategorie{
			{
				Id: "1",
			},
		},
	},

	{
		Id:    utility.NewUUID(),
		Name:  "Construire les premiers outils pour structurer sa pensée",
		Class: "Ps | Ms | Gs",
		SubCategories: []structure.SubCategorie{
			{
				Id: "1",
			},
		},
	},

	{
		Id:    utility.NewUUID(),
		Name:  "Explorer le monde",
		Class: "Ps | Ms | Gs",
		SubCategories: []structure.SubCategorie{
			{
				Id: "1",
			},
		},
	},

	{
		Id:    utility.NewUUID(),
		Name:  "Devenir élève",
		Class: "Ps | Ms | Gs",
	},
}

var subCategories = []structure.SubCategorie{
	{
		Id:        utility.NewUUID(),
		Name:      "L'oral",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[0].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "L'écrit",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[0].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "Les productions plastiques et visuelles",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[2].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "L'univers sonore",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[2].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "Découvrir les nombres et leur utilisation",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[3].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "Explorer des formes, des grandeurs, des suites",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[3].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "Se repérer dans le temps et l'espaces",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[4].Id},
	},

	{
		Id:        utility.NewUUID(),
		Name:      "Explorer le monde du vivant, des objets et de la matière",
		Class:     "Ps | Ms | Gs",
		Categorie: structure.Categorie{Id: categories[4].Id},
	},
}

// These function create all the table of the BDD and pre-filled the table categories and subCategories
func CreateDatabase() error {
	db, err := OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "teachers" (
  		"id" TEXT UNIQUE NOT NULL,
  		"username" TEXT UNIQUE NOT NULL,
  		"password" TEXT NOT NULL,
		"class" TEXT NOT NULL,
		"signingPath" TEXT,
		"key" TEXT NOT NULL
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "students" (
		"id" TEXT UNIQUE NOT NULL,
		"name" TEXT NOT NULL,
		"class" TEXT NOT NULL,
		"year" TEXT NOT NULL,
		"teacherId" TEXT NOT NULL REFERENCES"teachers"("id")
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "categories" (
		"id" TEXT UNIQUE NOT NULL,
		"name" TEXT NOT NULL,
		"class" TEXT NOT NULL
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "competences" (
		"id" TEXT UNIQUE NOT NULL,
		"name" TEXT NOT NULL,
		"imagePath" TEXT NOT NULL,
		"categorieId" TEXT NOT NULL REFERENCES"categories"("id"),
		"subCategorieId" TEXT NOT NULL REFERENCES"subCategories"("id"),
		"teacherId" TEXT NOT NULL REFERENCES"teachers"("id")
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "linkCompetenceEleve" (
		"id" TEXT UNIQUE NOT NULL,
		"studentId" TEXT NOT NULL REFERENCES"students"("id"),
		"competenceId" TEXT NOT NULL REFERENCES"competences"("id")
	)`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "subCategories" (
		"id" TEXT UNIQUE NOT NULL,
		"name" TEXT NOT NULL,
		"class" TEXT NOT NULL,
		"idCategorie" TEXT NOT NULL
	)`)
	if err != nil {
		return err
	}

	for _, v := range categories {
		_, err = db.Exec("INSERT INTO categories VALUES(?,?,?);", v.Id, v.Name, v.Class)
		if err != nil {
			return err
		}
	}

	for _, v := range subCategories {
		_, err = db.Exec("INSERT INTO subCategories VALUES(?,?,?,?)", v.Id, v.Name, v.Class, v.Categorie.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
