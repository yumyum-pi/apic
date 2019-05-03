package args

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yumyum-pi/apic/cli/utility"
)

// server file info
const serverFileName = "main"
const serverPath = "./"

const routeUtilityPath = "./route/utility/response.go"
const routeUtilityTmp = "./cli/template/route/response.tmp"

var serverFilePath = fmt.Sprintf("%v%v.go", serverPath, serverFileName)

// database file info
const dbFileName = "db"
const dbPath = "./model/db/"

var dbFilePath = fmt.Sprintf("%v%v.go", dbPath, dbFileName)

// MYPATH contains the absulute path of the project
var MYPATH string

// Init Initialization of the CLI Tool
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in main")
	}

	MYPATH = os.Getenv("MYPATH")
}

// CreateServer creates new server
func CreateServer() {

	// check if essential
	if !(utility.Exists("./model")) {
		os.MkdirAll("./model/db", os.ModePerm)
	}

	if !(utility.Exists("./route/utility")) {
		os.MkdirAll("./route/utility", os.ModePerm)
	}
	// check if main file exist
	if utility.Exists(serverFilePath) {

		fmt.Printf(`# "%v.go" file exist in "%v".`, serverFileName, serverPath)

		bo := utility.YesNo(`Do you want to remove the currrent file`)

		if bo {
			// delete file
			err := os.Remove(serverFilePath)
			if err != nil {
				utility.Exit(err, 2)
			}
			fmt.Println("-- File is deleted")
		} else {
			utility.ExitText("Process Cancled", 2)
		}
	}

	// check if database file exist
	if utility.Exists(dbFilePath) {

		fmt.Printf(`# "%v.go" file exist in "%v".`, dbFileName, dbPath)

		bo := utility.YesNo(`Do you want to remove the currrent file`)

		if bo {
			// delete file
			err := os.Remove(dbFilePath)
			if err != nil {
				utility.Exit(err, 2)
			}
			fmt.Println("-- File is deleted")
		} else {
			utility.ExitText("Process Cancled", 2)
		}
	}
	// read file
	dbTemplate := utility.ReadFile("cli/template/model/db.tmp")

	// save file
	utility.SaveFile(dbTemplate, dbFilePath)
	fmt.Printf("-- Created Database @ '%v' \n", dbFilePath)

	// make server file
	// read file
	serverTemplate := utility.ReadFile("cli/template/server.tmp")

	// check if the file exists
	if !utility.Exists(routeUtilityPath) {
		utility.SaveFile(utility.ReadFile(routeUtilityTmp), routeUtilityPath)
	}
	// change file
	var tempVars utility.TempVars
	tempVars = append(tempVars, utility.NewTempVar("MYPATH", MYPATH))
	utility.Template(&serverTemplate, tempVars)

	// save file
	utility.SaveFile(serverTemplate, serverFilePath)
	fmt.Printf("-- Created Server @ '%v' \n", serverFilePath)
}
