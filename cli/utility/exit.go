package utility

import (
	"fmt"
	"os"
)

// Exit will exit the code in cli with an error
func Exit(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}

// ExitText will exit the code in cli with an error using text error message
func ExitText(text string, code int) {
	Exit(fmt.Errorf(text), code)
}

// ConditionalExit will exit the code in cli if thier is an error
func ConditionalExit(err error, code int) {
	if err != nil {
		Exit(err, code)
	}
}
