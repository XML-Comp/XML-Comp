package comparer

import (
	"os"
	"testing"
)

func TestCompareFolder(t *testing.T) {
	//Testing for non-existing directories
	instance := Data{PathA: "fakeDir1", PathB: "fakeDir2"}
	if instance.CompareContainingFoldersAndFiles() == nil {
		//If error = nil, then the function failed to notice fakeDirectories
		t.Error("Oops! looks like error in existence of folders")
	}
	//Testing for blank values
	instance = Data{}
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
	instance = Data{PathA: pathA, PathB: pathB}
	t.Log(instance.CompareContainingFoldersAndFiles())
	//Logging as output is in form of error
}
