package test

import (
	"XML-Comp/comparer"
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
}
