package utility

import (
	"dossier_scolaire/structure"
	"strings"
)

// This function filter the competences by categorie and by name (and filter the student linked to this competence by class and name)
func FilterCompetencesByCategorieAndName(competences []structure.Competence) []structure.Competence {
	for i := len(competences) - 1; i > 0; i-- {
		for y := 0; y < i; y++ {
			if strings.ToLower(competences[i].Categorie.Name) < strings.ToLower(competences[y].Categorie.Name) ||
				strings.EqualFold(competences[i].Categorie.Name, competences[y].Categorie.Name) && strings.ToLower(competences[i].SubCategorie.Name) < strings.ToLower(competences[y].SubCategorie.Name) ||
				strings.EqualFold(competences[i].Categorie.Name, competences[y].Categorie.Name) && strings.EqualFold(competences[i].SubCategorie.Name, competences[y].SubCategorie.Name) && strings.ToLower(competences[i].Name) < strings.ToLower(competences[y].Name) {
				competences[i], competences[y] = competences[y], competences[i]
			}
		}
	}

	for i := 0; i < len(competences); i++ {
		for y := len(competences[i].Students) - 1; y > 0; y-- {
			for x := 0; x < y; x++ {
				if strings.ToLower(competences[i].Students[y].Class) > strings.ToLower(competences[i].Students[x].Class) ||
					strings.EqualFold(competences[i].Students[y].Class, competences[i].Students[x].Class) &&
						strings.ToLower(competences[i].Students[y].Name) < strings.ToLower(competences[i].Students[x].Name) {
					competences[i].Students[y], competences[i].Students[x] = competences[i].Students[x], competences[i].Students[y]
				}
			}
		}
	}

	return competences
}
