package args

import (
	"fmt"
	"strings"

	"github.com/yumyum-pi/apic/cli/utility"
)

const funcTmpPath = "cli/template/model/func.tmp"
const routeTmpPath = "cli/template/route/main.tmp"

// structVal is an structure stores Name and type of a single structure
type structVal struct {
	Name string
	Type string
}

func (s *structVal) toString() string {
	return fmt.Sprintf("%v %v `json:%q`;", s.Name, s.Type, strings.ToLower(s.Name))
}

type structVals []structVal

// add structVal to array
func (s *structVals) Add(name, kind string) {
	*s = append(*s, structVal{Name: name, Type: kind})
}

func (s *structVals) toString() (text string) {
	for _, aStructVar := range *s {
		text += aStructVar.toString()
	}
	return text
}

// Find the given string in array
func (s *structVals) Contains(x string) bool {
	for _, n := range *s {
		if x == n.Name {
			return true
		}
	}
	return false
}

// store a structure
type structure struct {
	Name string
	Vals structVals
	Func bool
}

func (s *structure) toString() (model, route string) {
	Name := strings.Title((*s).Name)
	newStruct := fmt.Sprintf("// %v need to explain\ntype %v struct {%v}", Name, Name, (*s).Vals.toString())
	var funcTmp, routeTmp string
	if (*s).Func {
		var varList, varListEq, mVarList, mVarList7, eq string
		for _, n := range (*s).Vals {
			if n.Name != "id" {
				varList = fmt.Sprintf("%v, %v", varList, n.Name)
				if n.Type == "string" {
					eq = fmt.Sprintf("%v, '%%s'", eq)
					varListEq = fmt.Sprintf("%v, %v='%%s'", varListEq, n.Name)
				} else {
					varListEq = fmt.Sprintf("%v, %v=%%d", varListEq, n.Name)
					eq = fmt.Sprintf("%v, '%%d'", eq)
				}
				mVarList = fmt.Sprintf("%v, m.%v", mVarList, n.Name)
				mVarList7 = fmt.Sprintf("%v, &m.%v", mVarList7, n.Name)
			}
		}
		varList = strings.Trim(varList, "! ,")
		varListEq = strings.Trim(varListEq, "! ,")
		mVarList = strings.Trim(mVarList, "! ,")
		mVarList7 = strings.Trim(mVarList7, "! ,")
		eq = strings.Trim(eq, "! ,")

		funcTmp = utility.ReadFile(funcTmpPath)
		routeTmp = utility.ReadFile(routeTmpPath)

		var modelTempVars utility.TempVars
		modelTempVars = append(modelTempVars, utility.NewTempVar("modelName", (*s).Name))
		modelTempVars = append(modelTempVars, utility.NewTempVar("ModelName", Name))
		modelTempVars = append(modelTempVars, utility.NewTempVar("varList", varList))
		modelTempVars = append(modelTempVars, utility.NewTempVar("varListEq", varListEq))
		modelTempVars = append(modelTempVars, utility.NewTempVar("mVarList", mVarList))
		modelTempVars = append(modelTempVars, utility.NewTempVar("&mVarList", mVarList7))
		modelTempVars = append(modelTempVars, utility.NewTempVar("eq", eq))
		modelTempVars = append(modelTempVars, utility.NewTempVar("MYPATH", MYPATH))

		utility.Template(&funcTmp, modelTempVars)
		utility.Template(&routeTmp, modelTempVars)

	}
	return fmt.Sprintf("\n%v\n%v\n", newStruct, funcTmp), routeTmp
}

type structures []structure

// Add structure to the array
func (s *structures) Add(name string, vals structVals, funcBool bool) {
	*s = append(*s, structure{Name: name, Vals: vals, Func: funcBool})
}

// Find the given string in array
func (s *structures) Contains(x string) bool {
	for _, n := range *s {
		if x == n.Name {
			return true
		}
	}
	return false
}

// Find the given string in array
func (s *structures) Find(x string) structure {
	var newStructure structure
	for i, n := range *s {
		if x == n.Name {
			newStructure = (*s)[i]
		}
	}
	return newStructure
}

// Template returns text file to be saved
func (s *structures) Template(name string) (model, route, mainRoute string) {
	// loop throught structures
	for _, aStructure := range *s {
		modelNFunc, routeNRouter := aStructure.toString()
		model += modelNFunc
		routeArry := strings.Split(routeNRouter, "//--->")
		fmt.Printf("> route = %d \n", len(routeArry))
		if len(routeArry) == 2 {
			route += routeArry[0]
			mainRoute += routeArry[1]
		}
	}

	return fmt.Sprintf("package model\nimport ( %q; %q )\n %v", "database/sql", "fmt", model), route, mainRoute
}
