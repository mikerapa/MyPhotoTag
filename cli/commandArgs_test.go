package cli

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// call InitLogger before running this test
	InitLogger("Error")
	os.Exit(m.Run())
}

func TestParseCommandLine(t *testing.T) {
	tests := []struct {
		name                string
		argString           string
		wantPhotoFilePath   string
		wantTagFilePath     string
		wantOutputPhotoPath string
		wantErr             bool
	}{
		{"no input", "", "", "", "", true},
		{"all input", "-p images/kid.jpg -t images/tag.png -o images/taggedLuca.jpg -l Trace", "images/kid.jpg", "images/tag.png", "images/taggedLuca.jpg", false},
		{"no output file", "-p images/kid.jpg -t images/tag.png -l Trace", "images/kid.jpg", "images/tag.png", "", false},
		{"all input with long names", "--photoPath images/kid.jpg --tagPath images/tag.png -o images/taggedLuca.jpg -l Trace", "images/kid.jpg", "images/tag.png", "images/taggedLuca.jpg", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotPhotoFilePath, gotTagFilePath, gotOutputPath, err := ParseCommandLine(strings.Split(tt.argString, " "))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPhotoFilePath != tt.wantPhotoFilePath {
				t.Errorf("ParseCommandLine() got = %v, want %v", gotPhotoFilePath, tt.wantPhotoFilePath)
			}
			if gotTagFilePath != tt.wantTagFilePath {
				t.Errorf("ParseCommandLine() got1 = %v, want %v", gotTagFilePath, tt.wantTagFilePath)
			}
			if gotOutputPath != tt.wantOutputPhotoPath {
				t.Errorf("ParseCommandLine() got2 = %v, want %v", gotOutputPath, tt.wantOutputPhotoPath)
			}
		})
	}
}
