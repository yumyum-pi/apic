package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yumyum-pi/apic/cli/args"
	"github.com/yumyum-pi/apic/cli/utility"
)

func main() {
	args.Init()
	// create a new flag
	createEndPoint := flag.NewFlagSet("endPoint", flag.ExitOnError)
	createServer := flag.NewFlagSet("server", flag.ExitOnError)
	testEndPoint := flag.NewFlagSet("test", flag.ExitOnError)

	// createRoute subcommand flag pointers
	name := createEndPoint.String("n", "", "name of the route")
	if len(os.Args) == 1 {
		fmt.Println("use --help to find info on options")
		return
	}

	switch os.Args[1] {
	case "server":
		createServer.Parse(os.Args[2:])
	case "endPoint":
		createEndPoint.Parse(os.Args[2:])
	case "test":
		testEndPoint.Parse(os.Args[2:])
	default:
		utility.Exit(fmt.Errorf("%q is not valid command", os.Args[1]), 2)
	}

	if createServer.Parsed() {
		args.CreateServer()
	} else if createEndPoint.Parsed() {
		if *name != "" {
			args.CreateEndPoint(*name)
		} else {
			utility.Exit(fmt.Errorf("-n cannot be empty"), 2)
		}
	} else if testEndPoint.Parsed() {
		args.Test()
	}
}
