package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ArxdSilva/XML-Comp/comparer"
)

var pathA = flag.String("original", "", "Full path directory of your RimWorld English folder (required)")
var pathB = flag.String("translation", "", "Full path directory of your RimWorld Translation folder (required)")

func main() {
	flag.Parse()
	args := os.Args
	// If we do not have enough params or help requested
	if len(args) < 2 || args[1] == "-h" {
		flag.Usage()
		os.Exit(1)
	}
	// If either pathA or pathB not provided exit
	if len(*pathA) == 0 || len(*pathB) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Creating instance ...")
	instance := comparer.Data{PathA: *pathA, PathB: *pathB}
	fmt.Println("Output:-")
	fmt.Println(instance.CompareContainingFoldersAndFiles())
}
