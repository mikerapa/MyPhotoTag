package photo

import (
	"image"
	"reflect"
	"testing"
)

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
