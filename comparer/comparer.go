package comparer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

var (
	translationName     string
	docs, lines, inNeed int
)

func Compare(original, translation string, b bool) error {
	if b {
		populateTranslationName(translation)
	}
	originalDir, err := readDir(original)
	if err != nil {
		return err
	}
	for _, f := range originalDir {
		if f.IsDir() {
			checkTransDirExists(f.Name())
			err = Compare(filepath.Join(original, f.Name()), filepath.Join(translation, f.Name()), false)
		} else {
			docs += 2
			err = readFiles(filepath.Join(original, f.Name()), filepath.Join(translation, f.Name()))
		}
		if err != nil {
			return err
		}
	}
	fmt.Printf("Documents scanned: %v | Lines scanned: %v | Translations needed: %v\n", docs, lines, inNeed)
	return nil
}

func readDir(path string) ([]os.FileInfo, error) {
	err := os.Chdir(path)
	if err != nil {
		return nil, err
	}
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	file, err := fi.Readdir(0)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	return file, nil
}

func readFiles(orgF, trltF string) error {
	err := os.Chdir(filepath.Dir(orgF))
	if err != nil {
		return err
	}
	fName := strings.Split(orgF, "/")
	fileName := fName[len(fName)-1]
	orgTags, err := readFile(fileName, filepath.Dir(orgF))
	if err != nil {
		return err
	}
	fName = strings.Split(trltF, "/")
	fileName = fName[len(fName)-1]
	trltTags, err := readFile(fileName, filepath.Dir(trltF))
	if err != nil {
		err = os.Chdir(filepath.Dir(trltF))
		if err != nil {
			return err
		}
		file, errCreate := os.Create(fileName)
		defer file.Close()
		if errCreate != nil {
			return errCreate
		}
		return nil
	}
	if trltTags == nil {
		return nil
	}
	missingTags := findMissing(orgTags, trltTags)
	if missingTags == nil {
		return nil
	}
	inNeed += len(missingTags)
	f, err := os.OpenFile(trltF, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, t := range missingTags {
		if (strings.Compare(string(t[:3]), "<!-") == 0) || (strings.Compare(string(t[:3]), "<--") == 0) {
			if _, err = f.WriteString(fmt.Sprintf("\n%s", t)); err != nil {
				return err
			}
			continue
		}
		if _, err = f.WriteString(fmt.Sprintf("\n%sAdd your translation here%s/%s", t, t[:1], t[1:])); err != nil {
			return err
		}
	}
	return nil
}

func readFile(file, path string) ([]string, error) {
	if file[len(file)-3:] != "xml" {
		return nil, nil
	}
	inFile, err := os.Open(filepath.Join(path, file))
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	tags := []string{}
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		indexStart := strings.Index(line, "<")
		indexEnd := strings.Index(line, ">")
		tag := line[indexStart : indexEnd+1]
		if string(tag[0]) == "/" {
			continue
		}
		if (indexStart != -1) && (indexEnd != -1) {
			tags = append(tags, tag)
		}
	}
	return tags, nil
}

// More info: https://gist.github.com/ArxdSilva/7392013cbba7a7090cbcd120b7f5ca31
func findMissing(fileFolderA, fileFolderB []string) []string {
	sort.Strings(fileFolderA)
	sort.Strings(fileFolderB)
	if reflect.DeepEqual(fileFolderA, fileFolderB) {
		return nil
	}
	for i := len(fileFolderA) - 1; i >= 0; i-- {
		for _, vD := range fileFolderB {
			if fileFolderA[i] == vD {
				fileFolderA = append(fileFolderA[:i], fileFolderA[i+1:]...)
				break
			}
		}
	}
	return fileFolderA
}

func checkTransDirExists(dir string) error {
	dir = strings.Replace(dir, "English", translationName, 1)
	_, err := os.Open(dir)
	if err != nil {
		splitedDirectory := strings.Split(dir, "/")
		parentDirFromSplit := dir[:len(dir)-len(splitedDirectory[len(splitedDirectory)-1])]
		os.Chdir(parentDirFromSplit)
		errMkdir := os.Mkdir(splitedDirectory[len(splitedDirectory)-1], 0700)
		if err != nil {
			return errMkdir
		}
	}
	return nil
}

func populateTranslationName(t string) {
	var splitTransPath []string
	slash := strings.Contains("/", t)
	if slash {
		splitTransPath = strings.Split(t, "/")
	} else {
		splitTransPath = strings.Split(t, "\\")
	}
	translationName = splitTransPath[len(splitTransPath)-1]
	return
}
