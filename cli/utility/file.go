package utility

import (
	"io/ioutil"
	"os"
)

// ReadFile read the template files
func ReadFile(path string) (text string) {
	file, err := ioutil.ReadFile(path)
	ConditionalExit(err, 1)

	//converting to strings
	text = string(file)
	return
}

// SaveFile is a function to save the file in desired path
func SaveFile(text, path string) {
	err := ioutil.WriteFile(path, []byte(text), 0777)
	ConditionalExit(err, 1)
	return
}

// Exists check if a file exits or not
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
