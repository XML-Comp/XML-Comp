package comparer

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_readFile(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirOriginal := dir + "/testPaths/Original/"
	// dirTranslation := dir + "/testPaths/Translation/"
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
