package comparer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	pathSep = string(os.PathSeparator)
)

var (
	// DocType is required If you want to use the package, so don't
	// forget to instantiate It before using the Compare function
	DocType string
	// Docs , Lines and InNeed are `metrics` of how the program is running
	Docs   int
	Lines  int
	InNeed int
)

// Compare is the function that takes two paths to comparable
// files and directories and builds It's differences into new
// files or new lines in the translated file
// getTranslationName determines If you
func Compare(original, translation string) error {
	originalDir, err := ReadDir(original)
	if err != nil {
		return err
	}
	for _, f := range originalDir {
		if f.IsDir() {
			errDirExists := checkTransDirExists(f.Name(), translation)
			if errDirExists != nil {
				return errDirExists
			}
			errCompare := Compare(filepath.Join(original, f.Name()), filepath.Join(translation, f.Name()))
			if errCompare != nil {
				return errCompare
			}
		} else {
			Docs += 2
			errRead := readFiles(filepath.Join(original, f.Name()), filepath.Join(translation, f.Name()))
			if errRead != nil {
				return errRead
			}
		}
	}
	return nil
}

func ReadDir(path string) ([]os.FileInfo, error) {
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
	fName := strings.Split(orgF, pathSep)
	fileName := fName[len(fName)-1]
	orgTags, err := readFile(fileName, filepath.Dir(orgF))
	if err != nil {
		return err
	}
	fName = strings.Split(trltF, pathSep)
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
	}
	missingTags := findMissing(orgTags, trltTags)
	if missingTags == nil {
		return nil
	}
	f, err := os.OpenFile(trltF, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	for k, v := range missingTags {
		if string(k[1]) ==pathSep {
			continue
		}
		InNeed++
		if (isFormatFile(k, "<!-")) || (isFormatFile(k, "<--")) || isFormatFile(k, "<?"+DocType) {
			if _, err = f.WriteString(fmt.Sprintf("\n%s", k)); err != nil {
				return err
			}
			continue
		}
		if _, err = f.WriteString(fmt.Sprintf("\n%s%s%s/%s", k, v, k[:1], k[1:])); err != nil {
			return err
		}
	}
	return nil
}

func isFormatFile(str, s string) bool {
	return strings.Contains(s, str)
}

func readFile(file, path string) (map[string]string, error) {
	splittedFileName := strings.Split(file, ".")
	if splittedFileName[len(splittedFileName)-1] != DocType {
		return nil, nil
	}
	inFile, err := os.Open(filepath.Join(path, file))
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	tags := map[string]string{}
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		Lines++
		line := scanner.Text()
		indexStart := strings.Index(line, "<")
		indexEnd := strings.Index(line, ">")
		if (len(line) == 0) || indexStart < 0 || indexEnd < 0 {
			continue
		}
		tag := line[indexStart: indexEnd+1]
		markers := strings.Split(tag, " ")
		if string(tag[0]) == pathSep {
			continue
		}
		tag = markers[0]
		valEnd := strings.LastIndex(line, "<")
		if valEnd < indexEnd {
			continue
		}
		translationValue := line[indexEnd+1: valEnd]
		if (indexStart != -1) && (indexEnd != -1) {
			tags[tag] = translationValue
		}
	}
	return tags, nil
}

func findMissing(original, translation map[string]string) map[string]string {
	missing := make(map[string]string)
	if reflect.DeepEqual(original, translation) {
		return nil
	}
	for k, v := range original {
		if _, ok := translation[k]; !ok {
			missing[k] = v
		}
	}
	return missing
}

func checkTransDirExists(dir, translation string) error {
	splitDir := strings.Split(dir, pathSep)
	dir = filepath.Join(translation, splitDir[len(splitDir)-1])
	_, err := os.Open(dir)
	if err != nil {
		splitedDirectory := strings.Split(dir, pathSep)
		parentDirFromSplit := dir[:len(dir)-len(splitedDirectory[len(splitedDirectory)-1])-1]
		os.Chdir(parentDirFromSplit)
		errMkdir := os.Mkdir(splitedDirectory[len(splitedDirectory)-1], 0700)
		if err != nil {
			return errMkdir
		}
	}
	return nil
}
