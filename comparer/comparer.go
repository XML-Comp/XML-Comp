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
func CompareContainingFoldersAndFiles(original, translation string) error {
	missFiles, missFolders, _, _, err := diff(original, translation)
	if err != nil {
		return err
	}
	if missFolders != nil {
		mFol := MissingFile{
			name: "missingFolders.txt",
			path: translation,
		}
		err = mFol.fileCreator(missFolders)
		if err != nil {
			return err
		}
	}
	if missFiles != nil {
		mFil := MissingFile{
			name: "missingFiles.txt",
			path: translation,
		}
		err = mFil.fileCreator(missFiles)
		if err != nil {
			return err
		}
	}
	return nil
}

// diff compares the given directories and returns their missingFiles, missingFolders & an error
func diff(A, B string) ([]string, []string, []string, []string, error) {
	filesInfo, err := ioutil.ReadDir(A)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	dirOri, filesOri := isItFileOrFiler(filesInfo)
	filesInfo, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	dirTrans, filesTrans := isItFileOrFiler(filesInfo)
	missingFolders, equalFolders := findMissing(dirOri, dirTrans)
	missingFiles, equalFiles := findMissing(filesOri, filesTrans)
	if (len(missingFolders) > 0) || (len(missingFiles) > 0) {
		return missingFiles, missingFolders, equalFiles, equalFolders, nil
	}
	return nil, nil, nil, nil, nil
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
func findMissing(fileFolderA, fileFolderB []string) ([]string, []string) {
	sort.Strings(fileFolderA)
	sort.Strings(fileFolderB)
	var equalFiles []string
	if reflect.DeepEqual(fileFolderA, fileFolderB) == true {
		return nil, nil
	}
	for i := len(fileFolderA) - 1; i >= 0; i-- {
		for _, vD := range fileFolderB {
			if fileFolderA[i] == vD {
				fileFolderA = append(fileFolderA[:i], fileFolderA[i+1:]...)
				equalFiles = append(equalFiles, vD)
			}
		}
	}
	return fileFolderA, equalFiles
}

// MissingFile defines the characteristics of a file to be created on your machine
// We use this struct to create "missingFiles.txt", "missingFolders.txt" and "missingTags.txt"
type MissingFile struct {
	name   string
	path   string
	prefix string
}

// fileCreator recieves parameters to write into the given path It's files.
// It's a way to create files and write content in It
func (f MissingFile) fileCreator(missing []string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s%s", f.path, f.prefix, f.name))
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
