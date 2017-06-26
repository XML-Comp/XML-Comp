package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/XML-Comp/XML-Comp/comparer"
)

const ver = "v0.24"

func main() {
	var (
		original    = flag.String("original", "", "Full path directory of your RimWorld English folder (required)")
		translation = flag.String("translation", "", "Full path directory of your RimWorld Translation folder (required)")
		docType     = flag.String("doc", "xml", "Type of the Doc that you want to compare (not required)")
		version     = flag.Bool("version", false, "Prints current version")
	)
	flag.Parse()
	args := os.Args
	switch {
	case len(args) < 2 || args[1] == "-h":
		flag.Usage()
		os.Exit(1)
	case *version:
		fmt.Println(ver)
		os.Exit(0)
	case len(*original) == 0 || len(*translation) == 0:
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println("Creating instance ...")
	fmt.Print("Output:- ")
	comparer.DocType = *docType
	err := comparer.Compare(*original, *translation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Docs comparisons are DONE!")
	fmt.Printf("Documents scanned: %v | Lines scanned: %v | Translations needed: %v\n", comparer.Docs, comparer.Lines, comparer.InNeed)
	os.Exit(0)
}
