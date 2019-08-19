package photo

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const taggedString = "(tagged)"

func openImageFile(imageFilePath string) (openedFile *os.File) {
	openedFile, err := os.Open(imageFilePath)
	if err != nil {
		// If the file open failed, try to log the application path as well.
		applicationPath, _ := filepath.Abs(os.Args[0])
		absoluteFilePath, _ := filepath.Abs(imageFilePath)
		cli.ConsoleLogger.Infof("Current application path: %s", applicationPath)
		cli.ConsoleLogger.Fatalf("failed to open: %s. Absolute path: %s.", err, absoluteFilePath)
	}
	return
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
