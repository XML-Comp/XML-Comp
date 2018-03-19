package comparer

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var wd, _ = os.Getwd()

func TestCompareFolder(t *testing.T) {
	fileMissing := "no such file or directory"
	if runtime.GOOS == "windows" {
		fileMissing = "The system cannot find the file specified."
	}
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
			Expected: fmt.Errorf("chdir fakeDir1: %s", fileMissing),
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

func TestReadFile(t *testing.T) {
	DocType = "xml"
	type args struct {
		file string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "read file",
			args: args{
				file: "File01.xml",
				path: filepath.Join(wd, "testPaths", "Original"),
			},
			want:    map[string]string{"<linha1>": "TEste", "<linha2>": "teste", "<linha3>": "teste"},
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

func TestFindMissing(t *testing.T) {
	type args struct {
		original    map[string]string
		translation map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "find missing Tags",
			args: args{
				original:    map[string]string{"<linha 1>": "", "<linha 2>": "", "<linha 3>": "teste"},
				translation: map[string]string{"<linha 1>": "", "<linha 2>": ""},
			},
			want: map[string]string{"<linha 3>": "teste"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissing(tt.args.original, tt.args.translation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFiles(t *testing.T) {
	tests := []struct {
		name    string
		orgF    string
		trltF   string
		wantErr bool
	}{
		{
			name:    "testing equal files",
			orgF:    filepath.Join("testPaths", "Original", "File01.xml"),
			trltF:   filepath.Join("testPaths", "Translation", "File01.xml"),
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
}

func TestReadFilesWithOutdatedTags(t *testing.T) {
	tests := []struct {
		name    string
		orgF    string
		trltF   string
		wantErr bool
	}{
		{
			name:    "testing equal files",
			orgF:    filepath.Join("testPaths", "Original", "File03.xml"),
			trltF:   filepath.Join("testPaths", "Translation", "File03.xml"),
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
	f, err := os.Open(filepath.Join(wd, tests[0].trltF))
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var outdated bool
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "OUTDATED") {
			outdated = true
		}
	}
	if !outdated {
		t.Error("File File03.xml should have a outdated tag")
	}
	err = ioutil.WriteFile(filepath.Join(wd, "testPaths", "Translation", "File03.xml"), []byte("<linha4></linha4>"), 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckTransDir(t *testing.T) {
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
		{
			d:       filepath.Join(wd, "English", "Dir1"),
			t:       filepath.Join(wd, "English", "Dir1"),
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
