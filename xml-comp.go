package main

import (
	"fmt"
	"os"

	c "github.com/ArxdSilva/XML-Comp/comparer"
)

const (
	Version = "0.0.2"
)

func main() {
	//Help option
	if len(os.Args) == 1 { // xml-comp
		fmt.Println("Kindly mention options")
		fmt.Println("Eg: xml-comp help")
	} else if len(os.Args) == 2 {
		//Help output
		if os.Args[1] == "help" {
			fmt.Println("You need two paths that we call pathA & pathB, which are described bellow:")
			fmt.Println("	pathA: /home/user/folder1")
			fmt.Println("	pathB: /home/user/folder2")
			fmt.Println("Use: xml-comp pathA pathB")
		}
	} else { //xml-comp pathA pathB
		pathA := os.Args[1]
		pathB := os.Args[2]
		fmt.Println("Creating instance ...")
		fmt.Println("Output:-")
		fmt.Println(c.CompareContainingFoldersAndFiles(pathA, pathB))
	}
}
