package utility

import (
	"dossier_scolaire/structure"
	"strings"
)

// These function take a slice of category and add in a new array each category who can be used by a teacher before return it
func FilterCategoriesForTeacher(categories []structure.Categorie, teacher structure.Teacher) []structure.Categorie {
	var res []structure.Categorie

	for _, v := range categories {
		if Contains(strings.Split(v.Class, " | "), strings.Split(teacher.Class, " | ")) {
			res = append(res, v)
		}
	}

	return res
}
