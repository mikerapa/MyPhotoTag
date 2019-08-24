package main

import (
	"github.com/mikerapa/MyPhotoTag/cli"
	"github.com/mikerapa/MyPhotoTag/photo"
	"os"
)

func main() {
	// set up the console logger with a default value
	cli.InitLogger("Error")

	photoFilePath, tagFilePath, outputPath, err := cli.ParseCommandLine(os.Args[1:])

	if err != nil {
		cli.ConsoleLogger.Info("Command line %a", os.Args[1:])
		cli.ConsoleLogger.Panic("Cannot process command line input")
	}
	if outputPath == "" {
		// no output path is specified
		outputPath = photoFilePath
		cli.ConsoleLogger.Infof("Output path was not specified. Using %s as output path for tagged photo.", outputPath)
	}

	if stat, err := os.Stat(photoFilePath); err == nil && stat.IsDir() {
		// path is a directory
		// set up the channels
		foundFiles := make(chan string)

		go photo.FindFilesAsync(photoFilePath, foundFiles)
		// media.GetMediaMetaData(<-foundFiles, log)

		for f := range foundFiles {
			outputFilePath := photo.GetOutputFilePath(outputPath, f)
			cli.ConsoleLogger.Trace("Ready to tag file ", f, "and output to ", outputFilePath)
			photo.TagPhoto(f, tagFilePath, outputFilePath)
		}
	} else {
		// single file
		outputFilePath := photo.GetOutputFilePath(outputPath, photoFilePath)
		photo.TagPhoto(photoFilePath, tagFilePath, outputFilePath)
	}

}
