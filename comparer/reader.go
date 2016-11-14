package comparer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readFile recieves a complete path of a file and returns It's tags & an error
func readFile(file, path string) []string {
	inFile, _ := os.Open(path + file)
	defer inFile.Close()
	tags := []string{}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		indexStart := strings.Index(line, "<")
		indexEnd := strings.Index(line, ">")
		if (indexStart != -1) && (indexEnd != -1) {
			tags = append(tags, line[indexStart:indexEnd])
		}
	}
	return tags
}

func readFiles(files []string, pathA, pathB string) error {
	for _, file := range files {
		tagsA := readFile(file, pathA)
		if tagsA == nil {
			return fmt.Errorf("Expected tagsA not nil, received: %v", tagsA)
		}
		tagsB := readFile(file, pathA)
		if tagsB == nil {
			return fmt.Errorf("Expected tagsB not nil, received: %v", tagsB)
		}
		missing := findMissing(tagsA, tagsB)
		err := createOutuputFile(pathB, file, "MissingTags.txt", missing)
		if err != nil {
			return err
		}
	}
	return nil
}
