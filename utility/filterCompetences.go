package utility

import (
	"dossier_scolaire/structure"
	"sort"
	"strings"
)

// These function used a bubble sort to sort a slice of competence alphabetically
func FilterCompetences(arr []structure.Competence) []structure.Competence {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Categorie.Name != arr[j].Categorie.Name {
			return RemoveAccent(arr[i].Categorie.Name) < RemoveAccent(arr[j].Categorie.Name)
		}
		if arr[i].SubCategorie.Name != arr[j].SubCategorie.Name {
			return RemoveAccent(arr[i].SubCategorie.Name) < RemoveAccent(arr[j].SubCategorie.Name)
		}
		return RemoveAccent(arr[i].Name) < RemoveAccent(arr[j].Name)
	})

	return arr
}

// These function remove all frensh accent in a string
func RemoveAccent(s string) string {
	accentMap := map[string]string{
		"à": "a", "â": "a", "æ": "a", "á": "a",
		"ç": "c",	
		"é": "e", "è": "e", "ê": "e", "ë": "e",
		"î": "i", "í": "i", "ì": "i", "ï": "i",
		"ñ": "n",
		"ô": "o", "ó": "o", "ò": "o", "ö": "o", "õ": "o", "ø": "o", "œ": "o",
		"š": "s",
		"ù": "u", "û": "u", "ú": "u", "ü": "u",
		"ý": "y", "ÿ": "y",
		"ž": "z",
	}

	for i, v := range accentMap {
		s = strings.ReplaceAll(s, i, v)
	}

	return s
}
