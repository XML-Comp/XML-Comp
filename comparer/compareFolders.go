package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
)

// FoldersAndFiles looks to two different directories,
// and creates a file named "missingFolders.txt" or "missingFiles.txt"
// with the missing files and folders
func FoldersAndFiles(original, translation string) error {
	missFiles, missFolders, err := diff(original, translation)
	if err != nil {
		return err
	}
	if missFolders != nil {
		err = createOutuputFile(translation, "", "missingFolders.txt", missFolders)
		if err != nil {
			return err
		}
	}
	if missFiles != nil {
		err = createOutuputFile(translation, "", "missingFolders.txt", missFiles)
		if err != nil {
			return err
		}
	}
	return nil
}

func diff(A, B string) (missFiles, missFolders []string, err error) {
	filesInfo, err := ioutil.ReadDir(A)
	if err != nil {
		return nil, nil, err
	}
	dirOri, filesOri := appendFileOrDir(filesInfo)
	filesInfo, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, nil, err
	}
	dirTrans, filesTrans := appendFileOrDir(filesInfo)
	missingFolders := findMissing(dirOri, dirTrans)
	missingFiles := findMissing(filesOri, filesTrans)
	return missingFiles, missingFolders, nil
}

func appendFileOrDir(filesInfo []os.FileInfo) ([]string, []string) {
	var dir, refFile []string
	for _, v := range filesInfo {
		if v.IsDir() {
			dir = append(dir, v.Name())
		} else {
			refFile = append(refFile, v.Name())
		}
	}
	return dir, refFile
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
			}
		}
	}
	return fileFolderA
}

// createOutuputFile create the file with the missing files and folders
func createOutuputFile(path, prefix, name string, miss []string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s%s", path, prefix, name))
	defer file.Close()
	if err != nil {
		return err
	}
	for _, v := range miss {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	return nil
}
