package cli

import "testing"

func TestDeriveOutputFilePath(t *testing.T) {
	tests := []struct {
		name               string
		photoFilePath      string
		wantOutputFilePath string
	}{
		{"simple path", "photos/luca.jpg", "photos/luca(tagged).jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputFilePath := DeriveOutputFilePath(tt.photoFilePath); gotOutputFilePath != tt.wantOutputFilePath {
				t.Errorf("DeriveOutputFilePath() = %v, want %v", gotOutputFilePath, tt.wantOutputFilePath)
			}
		})
	}
}
