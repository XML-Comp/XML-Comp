package comparer

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var tests = []struct {
	PathA    string
	PathB    string
	Expected error
}{
	{
		PathA:    "fakeDir1",
		PathB:    "fakeDir2",
		Expected: fmt.Errorf("open fakeDir1: no such file or directory"),
	},
}

func TestCompareFolder(t *testing.T) {
	for _, test := range tests {
		err := FoldersAndFiles(test.PathA, test.PathB)
		if !reflect.DeepEqual(err.Error(), test.Expected.Error()) {
			t.Errorf("Wanted error %v, got %v", test.Expected, err)
		}
	}
	tmpDir := os.TempDir()
	PathA := tmpDir + "_Dir1/subDir/superSub"
	err := os.MkdirAll(PathA, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", PathA, err)
	}
	PathB := tmpDir + "_Dir2/subDir"
	err = os.MkdirAll(PathB, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", PathB, err)
	}
	err = FoldersAndFiles(PathA, PathB)
	if err != nil {
		t.Errorf("Wanted error <nil>, got %v", err)
	}
}
