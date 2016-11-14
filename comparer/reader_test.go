package comparer

import (
	"os"
	"reflect"
	"testing"
)

var dir, _ = os.Getwd()
var dirOriginal = dir + "/testPaths/Original/"
var dirTranslation = dir + "/testPaths/Translation/"

func Test_readFile(t *testing.T) {
	type args struct {
		file string
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test reading original file",
			args: args{
				file: "File01.xml",
				path: dirOriginal,
			},
			want: []string{"<linha 1>", "<linha 2>", "<linha 3>"},
		},
	}
	for _, tt := range tests {
		if got := readFile(tt.args.file, tt.args.path); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. readFile() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_readFiles(t *testing.T) {
	type args struct {
		files []string
		pathA string
		pathB string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test reading no files",
			args: args{
				files: []string{},
				pathA: dirOriginal,
				pathB: dirTranslation,
			},
			wantErr: false,
		},
		{
			name: "test reading multiple files",
			args: args{
				files: []string{"File01.xml", "File02.xml"},
				pathA: dirOriginal,
				pathB: dirTranslation,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := readFiles(tt.args.files, tt.args.pathA, tt.args.pathB); (err != nil) != tt.wantErr {
			t.Errorf("%q. readFiles() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
