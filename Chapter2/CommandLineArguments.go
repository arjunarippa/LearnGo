/*
File:	HelloWorld.go
Author:	Arjun Arippa
Desc:	The below program describes how to parse the command line arugments in GoLang. The package - OS - is used in parsing the command
		line arguments.

 */
package main

import (
	"os"
	"fmt"
)

func main() {
	var comment, delimiter string
	for i:=1; i < len(os.Args); i++{
		comment+=delimiter+os.Args[i]
		delimiter=" "
	}
	fmt.Println(" The command line arguments are ", comment)
}
