package utility

import (
	"dossier_scolaire/structure"
)

// These function used a bubble sort to sort a slice of competence alphabetically
func FilterCompetences(arr []structure.Competence) []structure.Competence {
	for i := len(arr) - 1; i > 0; i-- {
		for y := 0; y < i; y++ {
			if arr[i].Categorie.Name < arr[y].Categorie.Name && arr[i].Name < arr[y].Name && arr[i].SubCategorie.Name < arr[y].SubCategorie.Name {
				arr[i], arr[y] = arr[y], arr[i]
			}
		}
	}

	return arr
}
