package photo

import "testing"

func TestGetOutputFilePath(t *testing.T) {
	tests := []struct {
		name               string
		outputPath         string
		photoFilePath      string
		wantOutputFilePath string
	}{
		{"Basic", "images/outputimages", "images/inputimages/kid.jpg", "images/outputimages/kid(tagged).jpg"},
		{"Basic (Png)", "images/outputimages", "images/inputimages/ladies.png", "images/outputimages/ladies(tagged).png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputFilePath := GetOutputFilePath(tt.outputPath, tt.photoFilePath); gotOutputFilePath != tt.wantOutputFilePath {
				t.Errorf("GetOutputFilePath() = %v, want %v", gotOutputFilePath, tt.wantOutputFilePath)
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
