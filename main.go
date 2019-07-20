package main

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"github.com/mikerapa/MyPhotoTag/photo"
)

func main() {
	// set up the console logger with a default value
	cli.InitLogger("Error")

	photoFilePath, tagFilePath, outputFilePath := cli.ParseCommandParameters()
	photo.TagPhoto(photoFilePath, tagFilePath, outputFilePath)

}
