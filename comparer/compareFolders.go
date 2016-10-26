package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CompareContainingFoldersAndFiles looks to two different directories given called d.PathA & d.PathB, and creates a file named
// "missingFolders.txt"
func (d *Data) CompareContainingFoldersAndFiles() error {
	if (d.PathA == "") || (d.PathB == "") {
		return fmt.Errorf("Empty path")
	}
	missing, err := d.missingFolders(d.PathA, d.PathB)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/missingFolders.txt", d.PathB))
	if err != nil {
		return err
	}
	defer file.Close()
	mfldr := []byte("<--Missing Folders-->\n")
	file.Write(mfldr)
	for _, v := range missing.missingFolder {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	file, err = os.Create(fmt.Sprintf("%s/missingFiles.txt", d.PathB))
	if err != nil {
		return err
	}
	mfls := []byte("<--Missing Files-->\n")
	file.Write(mfls)
	for _, v := range missing.missingFiles {
		d := []byte(fmt.Sprintf("- %s\n", v))
		file.Write(d)
	}
	return nil
}

func (d *Data) missingFolders(A, B string) (*Data, error) {
	files, err := ioutil.ReadDir(A)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		if v.IsDir() {
			d.dirA = append(d.dirA, v.Name())
		}
		d.fileA = append(d.fileA, v.Name())
	}
	files, err = ioutil.ReadDir(B)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		if v.IsDir() {
			d.dirB = append(d.dirB, v.Name())
		}
		d.fileB = append(d.fileB, v.Name())
	}
	for _, vA := range d.dirA {
		for iB, vB := range d.dirB {
			if (vA != vB) && (iB == len(d.dirB)-1) {
				d.missingFolder = append(d.missingFolder, vA)
			}
		}
	}
	for _, vA := range d.dirA {
		for iB, vB := range d.dirB {
			if (vA != vB) && (iB == len(d.fileB)-1) {
				d.missingFolder = append(d.missingFiles, vA)
			}
		}
	}
	if (len(d.missingFolder) > 0) || (len(d.missingFiles) > 0) {
		return d, nil
	}
	return nil, fmt.Errorf("No folders or files missing in PathB!")
}
