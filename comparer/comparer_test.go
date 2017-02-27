package comparer

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var wd, _ = os.Getwd()

func TestCompareFolder(t *testing.T) {
	tests := []struct {
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
	}
	for _, test := range tests {
		err := Compare(test.PathA, test.PathB)
		if !reflect.DeepEqual(err.Error(), test.Expected.Error()) {
			t.Errorf("Wanted error %v, got %v", test.Expected, err)
		}
	}
	err := Compare(filepath.Join(wd, "testPaths", "Original"), filepath.Join(wd, "testPaths", "Translation"))
	if err != nil {
		t.Errorf("Wanted error %v, got %v", nil, err)
	}
}

func Test_readFile(t *testing.T) {
	DocType = "xml"
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

func Test_readFiles(t *testing.T) {
	tests := []struct {
		name    string
		orgF    string
		trltF   string
		wantErr bool
	}{
		{
			name:    "testing equal files",
			orgF:    "testPaths/Original/File01.xml",
			trltF:   "testPaths/Translation/File01.xml",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readFiles(filepath.Join(wd, tt.orgF), filepath.Join(wd, tt.trltF)); (err != nil) != tt.wantErr {
				t.Errorf("readFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	err := readFiles(filepath.Join(wd, "testPaths/Original/File01.xml"), filepath.Join(wd, "testPaths/Translation/File03.xml"))
	if err != nil {
		t.Errorf("readFiles() error = %v, wantErr %v", err, nil)
	}
}

func Test_checkTransDir(t *testing.T) {
	os.Chdir("..")
	wd, _ := os.Getwd()
	tests := []struct {
		d, t    string
		wantErr bool
	}{
		{
			d:       filepath.Join(wd, "English", "Dir1"),
			t:       filepath.Join(wd, "Translation"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := checkTransDirExists(tt.d, tt.t); (err != nil) != tt.wantErr {
			t.Errorf("checkTransDir() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
	err := os.Remove(filepath.Join(wd, "Translation", "Dir1"))
	if err != nil {
		t.Errorf("Error = %v", err)
	}
}
