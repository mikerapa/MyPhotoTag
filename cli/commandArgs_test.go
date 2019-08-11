package cli

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	// TODO move this to a common location
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
		//{"all input", `-p images/kid.jpg -t images/tag.png -o images/taggedLuca.jpg -l Trace`, `images/kig.jpg`, `images/tag.png`, `images/taggedLuca.jpg`, false},
		{"all input", "-p images/kid.jpg -t images/tag.png -o images/taggedLuca.jpg -l Trace", "images/kid.jpg", "images/tag.png", "images/taggedLuca.jpg", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//_, _, _, err := ParseCommandLine(strings.Split(tt.argString, " "))
			gotPhotoFilePath, gotTagFilePath, gotOutputPhotoPath, err := ParseCommandLine(strings.Split(tt.argString, " "))
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
			if gotOutputPhotoPath != tt.wantOutputPhotoPath {
				t.Errorf("ParseCommandLine() got2 = %v, want %v", gotOutputPhotoPath, tt.wantOutputPhotoPath)
			}
		})
	}
}
