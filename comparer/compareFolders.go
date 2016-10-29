package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
)

// CompareContainingFoldersAndFiles looks to two different directories given called Original & Translation,
// and creates a file named "missingFolders.txt" or "missingFiles.txt" or both depending on It's differences
func CompareContainingFoldersAndFiles(Original, Translation string) error {
	missFiles, missFolders, err := missing(Original, Translation)
	if err != nil {
		return err
	}
	if missFolders != nil {
		file, err := os.Create(fmt.Sprintf("%s/missingFolders.txt", Translation))
		if err != nil {
			return err
		}
		defer file.Close()
		for _, v := range missFolders {
			d := []byte(fmt.Sprintf("- %s\n", v))
			file.Write(d)
		}
	}
	if missFiles != nil {
		file, err := os.Create(fmt.Sprintf("%s/missingFiles.txt", Translation))
		if err != nil {
			return err
		}
		for _, v := range missFiles {
			d := []byte(fmt.Sprintf("- %s\n", v))
			file.Write(d)
		}
	}
	return nil
}

func missing(A, B string) ([]string, []string, error) {
	// var dirOri, dirTrans, filesOri, filesTrans []string
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
	missingFolder := findMissing(dirOri, dirTrans)
	missingFiles := findMissing(filesOri, filesTrans)
	if (len(missingFolder) > 0) || (len(missingFiles) > 0) {
		return missingFiles, missingFolder, nil
	}
	return nil, nil, nil
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
	if reflect.DeepEqual(fileFolderA, fileFolderB) == true {
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
