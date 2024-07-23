package utility

import "dossier_scolaire/structure"

func GenerateString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[structure.SeededRand.Intn(len(charset))]
	}
	return string(b)
}
