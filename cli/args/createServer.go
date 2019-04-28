package args

import (
	"fmt"
	"os"

	"github.com/yumyum-pi/apic/cli/utility"
)

// CreateServer creates new server
func CreateServer(serverFilePath string, port, IP *string) {
	// check if main file exist
	if utility.Exists(serverFilePath) {

		fmt.Println(`"main.go" file exist in the root folder.`)

		var bo string
		fmt.Printf(`Do you want to remove the currrent "main.go" file. (Y/N) ? `)
		_, err := fmt.Scan(&bo)

		if err != nil {
			utility.Exit(err, 2)
		}

		if bo == "y" || bo == "yes" || bo == "Y" || bo == "YES" {
			// delete file
			err := os.Remove(serverFilePath)
			if err != nil {
				utility.Exit(err, 2)
			}
			fmt.Println("the File is deleted")
		} else if bo == "n" || bo == "no" || bo == "N" || bo == "NO" {
			fmt.Println("Process Cancled")
			os.Exit(2)
		} else {
			utility.ExitText("Undefined Input. Please restart", 2)
		}
	}

	// read file
	template := utility.ReadFile("cli/template/server.tmp")

	p := fmt.Sprintf(`const port = "%v"`, *port)
	ip := fmt.Sprintf(`const ip = "%v"`, *IP)
	// change file
	var tempVars utility.TempVars
	tempVars = append(tempVars, utility.NewTempVar("port", p))
	tempVars = append(tempVars, utility.NewTempVar("ip", ip))
	utility.Template(&template, tempVars)

	// save file
	utility.SaveFile(template, serverFilePath)
	fmt.Printf("You have created server @ '%v' \n", serverFilePath)
}
