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
