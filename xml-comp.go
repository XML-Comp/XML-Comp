package main

import (
	"flag"
	"fmt"
	"os"

	c "github.com/ArxdSilva/XML-Comp/comparer"
)

const (
	Version = "0.0.21"
)

var original = flag.String("original", "", "Full path directory of your RimWorld English folder (required)")
var translation = flag.String("translation", "", "Full path directory of your RimWorld Translation folder (required)")

func main() {
	flag.Parse()
	args := os.Args
	// If we do not have enough params or help requested
	if len(args) < 2 || args[1] == "-h" {
		flag.Usage()
		os.Exit(1)
	}
	// If either original or translation not provided exit
	if len(*original) == 0 || len(*translation) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Creating instance ...")
	fmt.Println("Output:-")
	fmt.Println(c.CompareContainingFoldersAndFiles(*original, *translation))
}
