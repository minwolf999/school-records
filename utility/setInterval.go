package utility

import (
	"time"
)

// This asynchrone function is a remake of the javascript setInterval function.
// 
// These function take 2 arguments:
// 
// - the interval delay between each function call
// 
// - the function to execute
func SetInterval(delai time.Duration, f func()) {
	startTime := time.Now()

	for {
		dif := time.Since(startTime)

		if dif >= delai {
			startTime = time.Now()

			f()
		}
	}
}
