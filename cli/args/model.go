package args

import (
	"fmt"
	"strings"

	"github.com/yumyum-pi/apic/cli/utility"
)

const modelTemp = "cli/template/model/main.tmp"

// create a struct
func addStructVar(modelName string, varArry *structVals) {
	var variables = " "
	fmt.Println("-- Enter Variables. { [VarName1]:[VarName2]:...=[VarType] }")
	fmt.Print(">> ")
	fmt.Scan(&variables)

	// seperate keys from types
	someVar := strings.Split(variables, "=")

	// Check for valid inputs
	if len(someVar) != 2 {
		// send error type not written
		fmt.Println("# [Error]: Varible not properly formated.")
		utility.ExitText("> [Solution]: Use the following pattern : [VarName1]:[VarName2]:...=[VarType]", 2)
	} else if someVar[0] == "" {
		// send error that key cannot be empty
		utility.ExitText("# [Error]: VarName cannot be empty", 2)
	} else if someVar[1] == "" {
		// send error that type cannot be empty
		utility.ExitText("# [Error]: VarType cannot be empty", 2)
	}

	names := strings.Split(someVar[0], ":")

	// assign each value of arry with the same type
	for _, name := range names {
		// storing variable in array
		if name == "" {
			utility.ExitText("# [Error]: VarName cannot be empty", 2)
		}

		varArry.Add(strings.Title(name), someVar[1])
	}
}

// createStructLoop loop the structor creation function
func addStructVarLoop(modelName string) structVals {
	var ifLoop = true
	var varArry structVals
	for ifLoop {
		addStructVar(modelName, &varArry)

		// add an yes or no input
		// for adding more variables
		ifLoop = utility.YesNo("Want to add more variables")
	}

	// print the variables in the struct
	fmt.Println("\n-- Following are the name and type of varibale :")

	var vars string
	for _, Var := range varArry {
		fmt.Printf(" > %v \n", Var)
		vars = fmt.Sprintf("%v \n	%v	%v", vars, Var.Name, Var.Type)
	}

	return varArry
}

//create struct
func createStruct(name string, newStructures *structures) {
	var structName = ""
	var addID bool
	//check if name exist in newStructores
	if newStructures.Contains(name) {
		fmt.Println("-- Enter Struct Name")
		fmt.Print(">> ")

		fmt.Scan(&structName)
	} else if utility.YesNo(fmt.Sprintf(`Current Struct Name is "%v". Do you want to change it`, name)) {
		fmt.Println("-- Enter Struct Name")
		fmt.Print(">> ")

		fmt.Scan(&structName)
	} else {
		structName = name
		addID = true
	}
	// make struct
	newStruct := addStructVarLoop(structName)
	if addID {
		newStruct.Add("ID", "int")
	}

	funcBool := utility.YesNo("Want to AddFunctions")
	newStructures.Add(structName, newStruct, funcBool)
	// -- Add function for struct ? loop
	//   -- default: GgCcUuDd
	//     -- GetMany, GetSingle, CreateMany, Create, UpdateMany Update, DeleteMany, Delete
	// 	return structrure with a functions
	//model := fmt.Sprintf("type %v struct {\n %v\n}", strings.Title(structName), vars)
	//fmt.Println("model")
}

func createStructLoop(name string) structures {
	var ifLoop = true
	var newStructures structures
	for ifLoop {
		createStruct(name, &newStructures)

		// add an yes or no input
		// for adding more structs
		ifLoop = utility.YesNo("Want to add more structs")
	}
	if !newStructures.Contains(name) {
		utility.Exit(fmt.Errorf(`# [Error]: The model does not contain a "%v" struct`, name), 1)
	}
	return newStructures
}

// createModel create a model for route
func createModel(name string) structures {
	// -- Create Struct
	return createStructLoop(name)
}
