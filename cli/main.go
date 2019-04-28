package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yumyum-pi/apic/cli/args"
	"github.com/yumyum-pi/apic/cli/utility"
)

const serverFileName = "main"
const serverPath = "./"

var serverFilePath = fmt.Sprintf("%v%v.go", serverPath, serverFileName)

func main() {
	// create a new flag
	createServer := flag.NewFlagSet("createServer", flag.ExitOnError)

	//argument for createServer
	port := createServer.String("p", "8080", "Port no. for the server")
	IP := createServer.String("ip", "", "IP address for the server")

	if len(os.Args) == 1 {
		fmt.Println("use --help to find info on options")
		return
	}

	switch os.Args[1] {
	case "createServer":
		createServer.Parse(os.Args[2:])
	default:
		utility.Exit(fmt.Errorf("%q is not valid command", os.Args[1]), 2)
	}

	if createServer.Parsed() {
		args.CreateServer(serverFilePath, port, IP)
	}
}
