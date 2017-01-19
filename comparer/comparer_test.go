package comparer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var wd, _ = os.Getwd()

var tests = []struct {
	name     string
	PathA    string
	PathB    string
	Expected error
}{
	{
		name:     "no dir",
		PathA:    "fakeDir1",
		PathB:    "fakeDir2",
		Expected: fmt.Errorf("chdir fakeDir1: no such file or directory"),
	},
	// {
	// 	name:     "testPaths",
	// 	PathA:    filepath.Join(wd, "testPaths", "Original"),
	// 	PathB:    filepath.Join(wd, "testPaths", "Translation"),
	// 	Expected: nil,
	// },
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

func Test_readFile(t *testing.T) {
	type args struct {
		file string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "read file",
			args: args{
				file: "File01.xml",
				path: filepath.Join(wd, "testPaths", "Original"),
			},
			want:    []string{"<linha 1>", "<linha 2>", "<linha 3>"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.file, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMissing(t *testing.T) {
	type args struct {
		fileFolderA []string
		fileFolderB []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "find missing Tags",
			args: args{
				fileFolderA: []string{"<linha 1>", "<linha 2>", "<linha 3>"},
				fileFolderB: []string{"<linha 1>", "<linha 2>"},
			},
			want: []string{"<linha 3>"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissing(tt.args.fileFolderA, tt.args.fileFolderB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissing() = %v, want %v", got, tt.want)
			}
		})
	}
}
