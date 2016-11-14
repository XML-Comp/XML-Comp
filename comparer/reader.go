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
	fmt.Println(path + file)
	defer inFile.Close()
	tags := []string{}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		indexStart := strings.Index(line, "<")
		indexEnd := strings.Index(line, ">")
		if (indexStart != -1) && (indexEnd != -1) {
			tags = append(tags, line[indexStart:indexEnd+1])
		}
	}
	return tags
}

func readFiles(files []string, pathA, pathB string) error {
	for _, file := range files {
		tagsA, err := tags(file, pathA)
		if err != nil {
			return err
		}
		tagsB, err := tags(file, pathB)
		if err != nil {
			return err
		}
		fileSplited := strings.Split(file, ".")
		fileName := fileSplited[0] + "_"
		missing := findMissing(tagsA, tagsB)
		fmt.Print(missing)
		err = createOutuputFile(pathB, fileName, "MissingTags.txt", missing)
		if err != nil {
			return err
		}
	}
	return nil
}

func tags(file, path string) ([]string, error) {
	tag := readFile(file, path)
	if tag == nil {
		return nil, fmt.Errorf("Expected tagsB not nil, received: %v", tag)
	}
	return tag, nil
}
