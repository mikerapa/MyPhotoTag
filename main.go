package main

import (
	"./cli"
	"./photo"
)

func main() {
	// set up the console logger with a default value
	cli.InitLogger("Error")

	photoFilePath, tagFilePath, outputFilePath := cli.ParseCommandParameters()
	photo.TagPhoto(photoFilePath, tagFilePath, outputFilePath)

}
