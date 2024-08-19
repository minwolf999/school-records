package utility

import (
	"dossier_scolaire/structure"
	"strings"
)

// Filter the subCategories by name
func FilterSubCategoriesByName(subCategories []structure.SubCategorie) []structure.SubCategorie {
	for i := len(subCategories) - 1; i > 0; i-- {
		for y := 0; y < i; y++ {
			if strings.ToLower(subCategories[i].Name) > strings.ToLower(subCategories[y].Name) {
				subCategories[i], subCategories[y] = subCategories[y], subCategories[i]
			}
		}
	}

	return subCategories
}
