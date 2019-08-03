package photo

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const taggedString = "(tagged)"

func openImageFile(imageFilePath string) (openedFile *os.File) {
	openedFile, err := os.Open(imageFilePath)
	if err != nil {
		cli.ConsoleLogger.Fatalf("failed to open: %s", err)
	}
	return
}

func DeriveOutputFilePath(photoFilePath string) (outputFilePath string) {
	dir, fileName := filepath.Split(photoFilePath)
	outputFilePath = filepath.Join(dir, strings.TrimSuffix(fileName, path.Ext(fileName))+taggedString+path.Ext(fileName))
	return
}
