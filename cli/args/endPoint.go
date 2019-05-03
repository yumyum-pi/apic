package args

import (
	"fmt"

	"github.com/yumyum-pi/apic/cli/utility"
)

const modelPath = "./model/"
const routePath = "./route/"

// CreateEndPoint will create API end-point
// which include creating a model
// and creating a route
func CreateEndPoint(name string) {
	modelFilePath := fmt.Sprintf("%v%v.go", modelPath, name)
	routeFilePath := fmt.Sprintf("%v%v.go", routePath, name)

	fmt.Println("Creating the folowing files:")
	fmt.Printf("> Model of '%v' in '%v' \n", name, modelFilePath)
	// check if file already exist
	if utility.Exists(modelFilePath) {
		// exit the cli command
		fmt.Printf("# [Error]: File exists :%v \n", modelFilePath)
		utility.ExitText("> [Solution]: Please change the -n flag name try again", 2)
	}

	fmt.Printf("> Route of '%v' in '%v' \n", name, routeFilePath)
	// check if file already exist
	if utility.Exists(routeFilePath) {
		// exit the cli command
		fmt.Printf("# [Error]: File exists :%v \n", routeFilePath)
		utility.ExitText("> [Solution]: Please change the -n flag name try again", 2)
	}
	// Create Model
	newStructures := createModel(name)
	//fmt.Printf("%#v\n", newStructures)

	// Create Route
	// -- Add Function from struct
	model, route, mainRoute := newStructures.Template(name)
	// Save file
	// -- Save Model
	utility.SaveFile(model, modelFilePath)

	// -- Save Routes
	utility.SaveFile(route, routeFilePath)

	// -- Update Main.go
	mainTemp := utility.ReadFile(serverFilePath)
	var tempVars utility.TempVars
	tempVars = append(tempVars, utility.NewTempVar("route", mainRoute))
	utility.Template(&mainTemp, tempVars)

	// save file
	utility.SaveFile(mainTemp, serverFilePath)
}
