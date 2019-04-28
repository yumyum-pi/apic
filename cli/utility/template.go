package utility

import (
	"fmt"
	"strings"
)

// TempVar is type which stores key value pair for template
type TempVar struct {
	key   string
	value string
}

// NewTempVar Create a new TempVar
func NewTempVar(key, value string) TempVar {
	return TempVar{
		key:   key,
		value: value,
	}
}

// TempVars is an arry of Template variables
type TempVars []TempVar

const tempVarComment = "//--->"

// return a valid tempVarName
func tempVarName(varName string) string {
	return fmt.Sprintf("%v%v", tempVarComment, varName)
}

// Template function generates string from template file and template variables
func Template(template *string, tempVars TempVars) {
	// check if all the variable are not empty
	if *template == "" {
		ExitText("template is empty", 1)
	}

	if len(tempVars) == 0 {
		ExitText("tempVars is empty", 1)
	}

	// loop through all the items in the arry and replace each commnet into variable
	// @add check wheather all tempVarsName exit in the template text
	//	if not through an error
	for _, tempVar := range tempVars {
		*template = strings.Replace(*template, tempVarName(tempVar.key), tempVar.value, -1)
	}

}
