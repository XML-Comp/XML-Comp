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
	missFiles, missFolders, _, _, err := diff(Original, Translation)
	if err != nil {
		return err
	}
	if missFolders != nil {
		mFol := missingFile{
			name: "missingFolders.txt",
			path: Translation,
		}
		err = mFol.fileCreator(missFolders)
		if err != nil {
			return err
		}
	}
	if missFiles != nil {
		mFil := missingFile{
			name: "missingFiles.txt",
			path: Translation,
		}
		err = mFil.fileCreator(missFiles)
		if err != nil {
			return err
		}
	}
	return nil
}

func diff(A, B string) ([]string, []string, []string, []string, error) {
	// var dirOri, dirTrans, filesOri, filesTrans []string
	filesInfo, err := ioutil.ReadDir(A)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	dirOri, filesOri := appendFileOrDir(filesInfo)
	filesInfo, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	dirTrans, filesTrans := appendFileOrDir(filesInfo)
	missingFolders, equalFolders := findMissing(dirOri, dirTrans)
	missingFiles, equalFiles := findMissing(filesOri, filesTrans)
	if (len(missingFolders) > 0) || (len(missingFiles) > 0) {
		return missingFiles, missingFolders, equalFiles, equalFolders, nil
	}
	return nil, nil, nil, nil, nil
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
				break
			}
		}
	}
	return fileFolderA, equalFiles
}

// missingFile defines the characteristics of a file to be created on your machine
// It is used to create "missingFiles.txt", "missingFolders.txt" or whatever
// kind of file you want
type missingFile struct {
	name   string
	path   string
	prefix string
}

func (f missingFile) fileCreator(miss []string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s%s", f.path, f.prefix, f.name))
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
