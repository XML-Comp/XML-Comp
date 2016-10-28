package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
)

var missingFolder, missingFiles []string

// CompareContainingFoldersAndFiles looks to two different directories given called PathA & PathB, and creates a file named
// "missingFolders.txt"
func CompareContainingFoldersAndFiles(PathA, PathB string) error {
	if (PathA == "") || (PathB == "") {
		return fmt.Errorf("Empty path")
	}
	missFiles, missFolders, err := missing(PathA, PathB)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/missingFolders.txt", PathB))
	if err != nil {
		return err
	}
	defer file.Close()
	mfldr := []byte("<--Missing Folders-->\n")
	file.Write(mfldr)
	for _, v := range missFolders {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	file, err = os.Create(fmt.Sprintf("%s/missingFiles.txt", PathB))
	if err != nil {
		return err
	}
	mfls := []byte("<--Missing Files-->\n")
	file.Write(mfls)
	for _, v := range missFiles {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	return nil
}

func missing(A, B string) ([]string, []string, error) {
	var dirA, dirB, filesA, filesB []string
	filesInfo, err := ioutil.ReadDir(A)
	if err != nil {
		return nil, nil, err
	}
	dirA, filesA = appendFileOrDir(filesInfo, dirA, filesA)
	filesInfo, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, nil, err
	}
	dirB, filesB = appendFileOrDir(filesInfo, dirB, filesB)
	missingFolder = findMissing(dirA, dirB)
	missingFiles = findMissing(filesA, filesB)
	if (len(missingFolder) > 0) || (len(missingFiles) > 0) {
		return missingFiles, missingFolder, nil
	}
	return nil, nil, fmt.Errorf("No folders or files missing in PathB!")
}

func appendFileOrDir(filesInfo []os.FileInfo, dir, refFile []string) ([]string, []string) {
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
	diff := fileFolderA
	for i, v := range diff {
		for _, vD := range fileFolderB {
			if v == vD {
				diff = append(diff[:i], diff[i+1:]...)
				break
			}
		}
	}
	return diff
}
