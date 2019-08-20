package photo

import (
	"bufio"
	"github.com/mikerapa/MyPhotoTag/cli"
	"github.com/sirupsen/logrus"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const taggedString = "(tagged)"

func DecodeImageFile(filePath string) (image.Image, string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		cli.ConsoleLogger.Errorf("Unable to open file %s", filePath)
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(bufio.NewReader(f))
}

func DeriveOutputFilePath(photoFilePath string) (outputFilePath string) {
	dir, fileName := filepath.Split(photoFilePath)
	outputFilePath = filepath.Join(dir, strings.TrimSuffix(fileName, path.Ext(fileName))+taggedString+path.Ext(fileName))
	return
}

var ConsoleLogger logrus.Logger

// IsMediaFile Check to see if this file path is for a supported media file
func IsMediaFile(filePath string) bool {
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".jpg", ".gif", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

//FindFilesAsync finds media files and returns them to a chan
func FindFilesAsync(searchPath string, foundFilesChan chan string) {
	var fileList []string
	filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if IsMediaFile(path) {
			ConsoleLogger.Info("Found file:", path)
			fileList = append(fileList, path)
			foundFilesChan <- path
		}

		return nil
	})
	close(foundFilesChan)
}
