package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	operation, err := operationFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		project := projectName()
		paths := getLibPath()
		switch operation {
		case "set":
			set(paths, project)
		case "show":
			show(paths, project)
		case "status":
			status(paths, project)
		case "reset":
			reset(project)
		case "unset":
			unset(paths, project)
		default:
			fmt.Println("Invalid operation")
			os.Exit(1)
		}
	}
}

func operationFromCommandLine() (operation string, err error) {
	if len(os.Args) == 1 ||
		(len(os.Args) > 1 &&
			(os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help")) {
		err = fmt.Errorf("usage: %s set|show|status|reset|unset", filepath.Base(os.Args[0]))
		return "", err
	}
	if len(os.Args) > 1 {
		operation = os.Args[1]
	}
	return operation, nil
}
