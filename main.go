package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/XML-Comp/XML-Comp/comparer"
	"os/user"
	"runtime"
)

const ver = "v0.4"

func main() {
	var (
		original            = flag.String("original", getOriginalDir(), "Full path directory of your RimWorld English folder (required)")
		translation         = flag.String("translation", "", "Full path directory of your RimWorld Translation folder (required)")
		docType             = flag.String("doc", "xml", "Type of the Doc that you want to compare (not required)")
		multipleMsg         = "Considers the translation flag as a collection of translations, enabling 1:N comparison"
		multiple            = flag.Bool("multiple", false, multipleMsg)
		version             = flag.Bool("version", false, "Prints current version")
		docs, lines, inNeed int
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
	if *multiple {
		dir, err := comparer.ReadDir(*translation)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range dir {
			if d.IsDir() {
				err := comparer.Compare(*original, filepath.Join(*translation, d.Name()))
				if err != nil {
					log.Fatal(err)
				}
			}
			docs += comparer.Docs
			lines += comparer.Lines
			inNeed += comparer.InNeed
		}
	} else {
		err := comparer.Compare(*original, *translation)
		if err != nil {
			log.Fatal(err)
		}
		docs = comparer.Docs
		lines = comparer.Lines
		inNeed = comparer.InNeed
	}
	fmt.Println("Docs comparisons are DONE!")
	fmt.Printf("Documents scanned: %v | Lines scanned: %v | Translations needed: %v\n", docs, lines, inNeed)
	os.Exit(0)
}

// getOriginalDir gets the default Steam installation directory.
// Returns empty string if directory is not found.
func getOriginalDir() string {
	rimWorldDir := ""

	// https://support.steampowered.com/kb_article.php?ref=7710-tdlc-0426
	if runtime.GOOS == "windows" && runtime.GOARCH == "386" {
		rimWorldDir = `C:\Program Files\Steam\steamapps\common\RimWorld`
	} else if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		rimWorldDir = `C:\Program Files (x86)\Steam\steamapps\common\RimWorld`
	} else if runtime.GOOS == "darwin" {
		// https://gaming.stackexchange.com/a/219537
		currentUser, err := user.Current()
		if err != nil {
			return ""
		}

		// safe to use *nix path sep since we are in the darwin runtime
		rimWorldDir = filepath.Join(currentUser.HomeDir,
			`/Library/Application Support/Steam/steamapps/common/RimWorld/RimWorldMac.app`)
	}

	// we're in all runtimes here, use os dependent separator
	englishTrans := filepath.Join(rimWorldDir, "Mods", "Core", "Languages", "English")

	// do one final check to make sure the English dir exists
	_, err := os.Stat(englishTrans)
	if err != nil {
		englishTrans = ""
	}

	return englishTrans
}
