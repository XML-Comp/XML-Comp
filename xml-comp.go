package main

import (
	"fmt"
	"os"

	"XML-Comp/comparer"
)

const (
	Version = "0.0.2"
)

func main() {
	//map for options
	options := map[string]string{
		//help
		"help": `
To Use Folder Verification, usage: xml-comp -fd pathA pathB
You need two paths that we call pathA & pathB, which are described bellow:
	pathA: /home/user/folder1
	pathB: /home/user/folder2

To get current version, usage: xml-comp -v`,
		"-h": `
To Use Folder Verification, usage: xml-comp -fd pathA pathB
u need two paths that we call pathA & pathB, which are described bellow:
	pathA: /home/user/folder1
	pathB: /home/user/folder2

To get current version, usage: xml-comp -v`,
		//version
		"version": "Current version is " + Version,
		"-v":      "Current version is " + Version,
		//Folder and File function index
		"-fd": "Mention paths for comparing, usage: xml-comp -fd pathA pathB",
		"-fl": "Coming soon!",
	}
	//Help option
	if len(os.Args) == 1 { // xml-comp
		fmt.Println("Kindly mention options")
		fmt.Println("Eg: xml-comp help")
	} else if len(os.Args) == 2 {
		//Flags
		val := os.Args[1]
		//Existence of flags
		if len(options[val]) > 0 {
			//Flag exists
			fmt.Println(options[val])
		} else {
			//Doesn't exist
			fmt.Println("Seems like you chose wrong option, see $ xml-comp help")
		}
	} else {
		if os.Args[1] == "-fd" {
			//proceeding with action for folder comparer
			pathA := os.Args[1]
			pathB := os.Args[2]
			fmt.Println("Creating instance ...")
			instance := comparer.Data{PathA: pathA, PathB: pathB}
			fmt.Println("Output:-")
			fmt.Println(instance.CompareContainingFoldersAndFiles())
		} else {
			fmt.Println(options[os.Args[1]])
		}
	}
}
