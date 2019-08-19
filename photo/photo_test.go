package photo

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"image"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	// call InitLogger before running this test
	cli.InitLogger("Info")
	os.Exit(m.Run())
}

func Test_calculateTagCoordinate(t *testing.T) {
	tests := []struct {
		name                 string
		photoBounds          image.Rectangle
		tagBounds            image.Rectangle
		wantDestinationPoint image.Point
	}{
		{"Tag is smaller", image.Rect(0, 0, 500, 500), image.Rect(0, 0, 100, 100), image.Pt(400, 400)},
		{"Tag is larger by width", image.Rect(0, 0, 100, 500), image.Rect(0, 0, 101, 100), image.Pt(0, 0)},
		{"Tag is larger by height", image.Rect(0, 0, 500, 100), image.Rect(0, 0, 100, 101), image.Pt(0, 0)},
		{"Tag and photo have the same dimensions", image.Rect(0, 0, 100, 100), image.Rect(0, 0, 100, 100), image.Pt(0, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDestinationPoint := CalculateTagCoordinate(tt.photoBounds, tt.tagBounds); !reflect.DeepEqual(gotDestinationPoint, tt.wantDestinationPoint) {
				t.Errorf("calculateTagCoordinate() = %v, want %v", gotDestinationPoint, tt.wantDestinationPoint)
			}
		})
	}

}

func TestDeriveOutputFilePath(t *testing.T) {
	tests := []struct {
		name               string
		photoFilePath      string
		wantOutputFilePath string
	}{
		{"simple path", "photos/luca.jpg", "photos/luca(tagged).jpg"},
		{"Longer path", "/OtherFolder/photos/luca.jpg", "/OtherFolder/photos/luca(tagged).jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputFilePath := DeriveOutputFilePath(tt.photoFilePath); gotOutputFilePath != tt.wantOutputFilePath {
				t.Errorf("DeriveOutputFilePath() = %v, want %v", gotOutputFilePath, tt.wantOutputFilePath)
			}
		})
	}
}

func TestIsMediaFile(t *testing.T) {
	// Supported media ".gif", ".jpeg", ".png"
	tests := []struct {
		name string
		path string
		want bool
	}{
		{name: "jpg", path: "c:\\temp\\otherdir\\hang.jpg", want: true},
		{name: "JPG in caps", path: "c:\\temp\\thingy.JPG", want: true},
		{name: "png", path: "c:\\temp\\thingy.png", want: true},
		{name: "gif", path: "c:\\temp\\thingy.gif", want: true},
		{name: "xps", path: "c:\\temp\\thingy.xps", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMediaFile(tt.path); got != tt.want {
				t.Errorf("IsMediaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagPhoto(t *testing.T) {
	tests := []struct {
		name           string
		photoFilePath  string
		tagFilePath    string
		outputFilePath string
	}{
		{"jpg tag test", "../images/inputimages/kid.jpg", "../images/tag.png", "../images/outputimages/kid(tagged).jpg"},
		{"png tag test", "../images/inputimages/ladies.png", "../images/tag.png", "../images/outputimages/ladies(tagged).jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TagPhoto(tt.photoFilePath, tt.tagFilePath, tt.outputFilePath)
			if _, err := os.Stat(tt.outputFilePath); os.IsNotExist(err) {
				t.Errorf("Failed to create file: %s", tt.outputFilePath)
			} else {
				// remove the newly created photo
				removeError := os.Remove(tt.outputFilePath)
				if removeError != nil {
					t.Errorf("Could not clean up this test by deleting %s. This file may be locked", tt.outputFilePath)
				}
			}

		})
	}
}
