package comparer

import (
	"fmt"
	"os"
)

// reader recieves a path and a file and returns It's tags and an error
func reader(file, path string) ([]string, error) {
	var tags []string

	return tags, nil
}

func readFiles(files []string, pathA, pathB string) error {
	for _, v := range files {
		tagsA, err := reader(v, pathA)
		if err != nil {
			return err
		}
		tagsB, err := reader(v, pathA)
		if err != nil {
			return err
		}
		missingTags, _ := findMissing(tagsA, tagsB)
		// Create "nameMissing.txt"
		file, err := os.Create(fmt.Sprintf("%s/%sMissingTags.txt", pathB, v))
		if err != nil {
			return err
		}
		for _, v := range missingTags {
			d := []byte(fmt.Sprintf("- %s\n", v))
			file.Write(d)
		}
		defer file.Close()
	}
	return nil
}
