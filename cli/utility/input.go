package utility

import "fmt"

// YesNo function ask question in termial and reads user's input
// user has to write yes or no
// then the function returns bool according to the user's input
func YesNo(text string) bool {
	var value string
	fmt.Printf("\n>> %v? (Y/N) : ", text)
	_, err := fmt.Scan(&value)
	if err != nil {
		Exit(err, 2)
	}

	if value == "y" || value == "yes" || value == "Y" || value == "YES" {
		return true
	} else if value == "n" || value == "no" || value == "N" || value == "NO" {
		return false
	} else {
		ExitText("Undefined Input. Please restart", 2)
	}

	return false
}
