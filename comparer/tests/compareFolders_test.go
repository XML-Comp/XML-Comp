package test

import (
	"XML-Comp/comparer"
	"os"
	"testing"
)

func TestCompareFolder(t *testing.T) {
	//Testing for non-existing directories
	instance := comparer.Data{PathA: "fakeDir1", PathB: "fakeDir2"}
	if instance.CompareContainingFoldersAndFiles() == nil {
		//If error = nil, then the function failed to notice fakeDirectories
		t.Error("Oops! looks like error in existance of folders")
	}

	//Testing for blank values
	instance = comparer.Data{}
	if instance.CompareContainingFoldersAndFiles() == nil {
		//If error nil, then the validation was somehow bypassed
		t.Error("Oops! Error in Path validation")
	}

	//Testing for actual folders
	//the code will create 2 directories with files
	tmpDir := os.TempDir()
	pathA := tmpDir + "/_Dir1/subDir/superSub"
	err := os.MkdirAll(pathA, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", pathA, err)
	}

	pathB := tmpDir + "/_Dir2/subDir"
	err = os.MkdirAll(pathB, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", pathB, err)
	}
	//Temporary directories created

	instance = comparer.Data{PathA: pathA, PathB: pathB}
	t.Log(instance.CompareContainingFoldersAndFiles())
	//Logging as output is in form of error
}
