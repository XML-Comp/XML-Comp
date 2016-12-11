package comparer

import (
	"bufio"
	"os"
	"strings"
)

// readFile receives a complete path of a file and returns It's tags & an error
func readFile(file, path string) ([]string, error) {
	inFile, err := os.Open(path + file)
	if err != nil {
		return nil, err
	}
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
	return tags, nil
}

func readFiles(original, translation string) error {
	_, filesOri, err := directoriesAndFiles(original)
	if err != nil {
		return err
	}
	original = lastChar(original)
	translation = lastChar(translation)
	for _, file := range filesOri {
		tagsA, err := readFile(file, original)
		if err != nil {
			return err
		}
		tagsB, err := readFile(file, translation)
		if (err != nil) || (tagsB == nil) {
			continue
		}
		fileSplited := strings.Split(file, ".")
		fileName := fileSplited[0] + "_"
		missing := findMissing(tagsA, tagsB)
		if err := createOutuputFile(translation, fileName, "MissingTags.txt", missing); err != nil {
			return err
		}
	}
	return nil
}

func lastChar(s string) string {
	if s[len(s)-1:] != "/" {
		s = s + "/"
	}
	return s
}
