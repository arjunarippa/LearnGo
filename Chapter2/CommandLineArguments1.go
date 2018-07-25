/*
File:	CommandLineArguments1.go
Author:	Arjun Arippa
Desc:	The below program describes how to parse the command line arugments in GoLang. The package - OS - is used in parsing the command
		line arguments. This is a best and efficient solution for joining/parsing commandline arguments.

 */
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
}
