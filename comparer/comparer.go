package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
)

// FoldersAndFiles looks to two different directories,
// and creates a file named "missingFolders.txt" and/or "missingFiles.txt" and/or "<fileName>MissingTags.txt"
// with the missing files, folders and tags on each line of the file
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
		err = createOutuputFile(translation, "", "missingFiles.txt", missFiles)
		if err != nil {
			return err
		}
	}
	err = readFiles(original, translation)
	if err != nil {
		return err
	}
	return nil
}

func diff(original, translation string) (missingFiles, missingFolders []string, err error) {
	dirOri, filesOri, err := directoriesAndFiles(original)
	if err != nil {
		return nil, nil, err
	}
	dirTrans, filesTrans, err := directoriesAndFiles(translation)
	if err != nil {
		return nil, nil, err
	}
	missingFolders = findMissing(dirOri, dirTrans)
	missingFiles = findMissing(filesOri, filesTrans)
	return missingFiles, missingFolders, nil
}

// isItFileOrFolder receives all the content from the given directory and
// separates files from folders
func isItFileOrFolder(filesInfo []os.FileInfo) ([]string, []string) {
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

func directoriesAndFiles(language string) ([]string, []string, error) {
	filesInfo, err := ioutil.ReadDir(language)
	if err != nil {
		return nil, nil, err
	}
	dir, files := isItFileOrFolder(filesInfo)
	return dir, files, nil
}
