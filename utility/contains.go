package utility

// This function look if two strings has an common element
// 
// These functions take 2 slice of string has arguments and return a boolean (true if there is a common element else false)
func Contains(s []string, e []string) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}

	return false
}