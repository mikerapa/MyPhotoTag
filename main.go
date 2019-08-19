package main

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"github.com/mikerapa/MyPhotoTag/photo"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// set up the console logger with a default value
	cli.InitLogger("Error")

	photoFilePath, tagFilePath, outputFilePath, err := cli.ParseCommandLine(os.Args[1:])

	if err != nil {
		logrus.Panic("Cannot process command line input")
	}

	if stat, err := os.Stat(photoFilePath); err == nil && stat.IsDir() {
		// path is a directory

	} else {
		// single file
		// Make sure there is a valid outputPhotoPath
		if outputFilePath == "" {
			outputFilePath = photo.DeriveOutputFilePath(photoFilePath)
			cli.ConsoleLogger.Infof("Output path was not specified. Using %s as output path for tagged photo.", outputFilePath)
		}
		photo.TagPhoto(photoFilePath, tagFilePath, outputFilePath)
	}

	// set up the channels
	foundFiles := make(chan string)

	go photo.FindFilesAsync(photoFilePath, foundFiles)
	// media.GetMediaMetaData(<-foundFiles, log)

	for f := range foundFiles {
		outputFilePath = photo.DeriveOutputFilePath(f)
		cli.ConsoleLogger.Trace("Ready to tag file ", f, "and output to ", outputFilePath)
		photo.TagPhoto(f, tagFilePath, outputFilePath)
	}

}
