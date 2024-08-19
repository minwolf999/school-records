package utility

import (
	"dossier_scolaire/structure"
	"strings"
)

// This function filter the students by class and by name
func FilterByClassAndName(students []structure.Student) []structure.Student {
	for y := len(students) - 1; y > 0; y-- {
		for x := 0; x < y; x++ {
			if strings.ToLower(students[y].Class) > strings.ToLower(students[x].Class) ||
				strings.EqualFold(students[y].Class, students[x].Class) &&
					strings.ToLower(students[y].Name) < strings.ToLower(students[x].Name) {
				students[y], students[x] = students[x], students[y]
			}
		}
	}

	return students
}
