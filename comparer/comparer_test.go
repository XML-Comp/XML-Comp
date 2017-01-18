package comparer

import (
	"fmt"
	"io/ioutil"
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
		err := Compare(test.PathA, test.PathB)
		if !reflect.DeepEqual(err.Error(), test.Expected.Error()) {
			t.Errorf("Wanted error %v, got %v", test.Expected, err)
		}
	}
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("TempDir %q: %s", tmpDir, err)
	}
	PathA := tmpDir + "_Dir1/subDir/superSub"
	err = os.MkdirAll(PathA, 0777)

	if err != nil {
		t.Fatalf("MkdirAll %q: %s", PathA, err)
	}
	PathB := tmpDir + "_Dir2/subDir"
	err = os.MkdirAll(PathB, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", PathB, err)
	}
	err = Compare(PathA, PathB)
	if err != nil {
		t.Errorf("Wanted error <nil>, got %v", err)
	}
}
