package main

import (
	"fmt"
	"os"

	"github.com/ArxdSilva/XML-Comp/comparer"
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
		"-h": `Yo
To Use Folder Verification, usage: xml-comp -fd pathA pathB
u need two paths that we call pathA & pathB, which are described bellow:
	pathA: /home/user/folder1
	pathB: /home/user/folder2

To get current version, usage: xml-comp -v`,
		//version
		"version": "Current version is " + Version,
		"-v":      "Current version is " + Version,
		//Folder and File function index
		"-fd": "folder",
		"-fl": "file",
	}
	//Help option
	if len(os.Args) == 1 { // xml-comp
		fmt.Println("Kindly mention options")
		fmt.Println("Eg: xml-comp help")
	} else {
		//Flags

		pathA := os.Args[1]
		pathB := os.Args[2]
		fmt.Println("Creating instance ...")
		instance := comparer.Data{PathA: pathA, PathB: pathB}
		fmt.Println("Output:-")
		fmt.Println(instance.CompareContainingFoldersAndFiles())
	}
}
