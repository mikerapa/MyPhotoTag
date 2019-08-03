package main

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"github.com/mikerapa/MyPhotoTag/photo"
)

func main() {
	// set up the console logger with a default value
	cli.InitLogger("Error")

	photoFilePath, tagFilePath, outputFilePath := cli.ParseCommandLine()

	// Make sure there is a valid outputPhotoPath
	if outputFilePath == "" {
		outputFilePath = photo.DeriveOutputFilePath(photoFilePath)
		cli.ConsoleLogger.Infof("Output path was not specified. Using %s as output path for tagged photo.", outputFilePath)
	}

	photo.TagPhoto(photoFilePath, tagFilePath, outputFilePath)

}
