package utility

import "dossier_scolaire/structure"

// These function return an array of student
func YearListFromStudent(students []structure.Student) []string {
	var res []string

	for _, v := range students {
		if IsNotInTab(res, v.Year) {
			res = append(res, v.Year)
		}
	}

	// Sort the result slice
	for i := len(res)-1; i > 0; i-- {
		for y := 0; y < i; y++ {
			if res[i] < res[y] {
				res[i], res[y] = res[y], res[i]
			}
		}
	}

	return res
}
