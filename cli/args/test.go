package args

import (
	"fmt"

	"github.com/yumyum-pi/apic/cli/utility"
)

var test = structures{
	structure{
		Name: "rand",
		Vals: structVals{
			structVal{Name: "Rand1", Type: "string"},
			structVal{Name: "Rand2", Type: "string"},
			structVal{Name: "Rand3", Type: "string"},
			structVal{Name: "Rand4", Type: "int"},
		},
		Func: false,
	},
	structure{
		Name: "users",
		Vals: structVals{
			structVal{Name: "Name", Type: "string"},
			structVal{Name: "Gender", Type: "string"},
			structVal{Name: "Age", Type: "int"},
			structVal{Name: "ID", Type: "int"},
		},
		Func: true,
	},
}

// Test will run test
func Test() {
	//fmt.Println("this is a test")

	_, _, mainRoute := test.Template("users")

	// -- Update Main.go
	mainTemp := utility.ReadFile(serverFilePath)
	//fmt.Println(model)
	var tempVars utility.TempVars
	tempVars = append(tempVars, utility.NewTempVar("route", mainRoute))
	utility.Template(&mainTemp, tempVars)

	fmt.Println(mainTemp)

	// test.Template("user")
	// test[0].toString()
	// fmt.Println(user.toString())
}
