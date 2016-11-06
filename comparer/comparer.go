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
	dirOri, filesOri := isItFileOrFiler(filesInfo)
	filesInfo, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, nil, err
	}
	dirTrans, filesTrans := isItFileOrFiler(filesInfo)
	missingFolders := findMissing(dirOri, dirTrans)
	missingFiles := findMissing(filesOri, filesTrans)
	return missingFiles, missingFolders, nil
}

// isItFileOrFiler recieves all the content from the given directory and
// separates files from folders
func isItFileOrFiler(filesInfo []os.FileInfo) ([]string, []string) {
	var folders, files []string
	for _, v := range filesInfo {
		if v.IsDir() {
			folders = append(folders, v.Name())
		} else {
			files = append(files, v.Name())
		}
	}
	return folders, files
}

// findMissing takes two repos and checks If B has different files from A
// If B is missing something, It will remove from sliceA similar files or folders
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
func createOutuputFile(path, prefix, name string, missing []string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s%s", path, prefix, name))
	defer file.Close()
	if err != nil {
		return err
	}
	for _, v := range missing {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	return nil
}
