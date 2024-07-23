package utility

// These function look if a string is in a slice.
// 
// If the string is in the slice he return false else true
func IsNotInTab(tab []string, str string) bool {
	for _, v := range tab {
		if v == str {
			return false
		}
	}

	return true
}