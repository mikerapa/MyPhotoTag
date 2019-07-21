package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"path"
	"path/filepath"
	"strings"
)

const taggedString = "(tagged)"

func ParseCommandLine() (string, string, string) {

	var (
		photoFilePath   = kingpin.Flag("photoPath", "File path to the photo original").Short('p').Required().String()
		tagFilePath     = kingpin.Flag("tagPath", "File path to the tag png file").Short('t').Required().String()
		outputPhotoPath = kingpin.Flag("outputPhotoPath", "location to place the tagged photo").Short('o').String()
		logLevelString  = kingpin.Flag("logLevel", "Log Level (error, warning, info, debug, trace)").Short('l').String()
	)

	kingpin.Parse()

	ConsoleLogger.Level = GetLogLevelFromString(*logLevelString)

	// Make sure there is a valid outputPhotoPath
	if *outputPhotoPath == "" {
		*outputPhotoPath = DeriveOutputFilePath(*photoFilePath)
		ConsoleLogger.Infof("Output path was not specified. Using %s as output path for tagged photo.", *outputPhotoPath)
	}

	ConsoleLogger.Info("Command Argument photoPath: ", *photoFilePath)
	ConsoleLogger.Info("Command Argument tagPath: ", *tagFilePath)
	ConsoleLogger.Info("Command Argument outputFilePath: ", *outputPhotoPath)
	ConsoleLogger.Trace("Done parsing command line arguments")

	return *photoFilePath, *tagFilePath, *outputPhotoPath
}

func DeriveOutputFilePath(photoFilePath string) (outputFilePath string) {
	dir, fileName := filepath.Split(photoFilePath)
	outputFilePath = filepath.Join(dir, strings.TrimSuffix(fileName, path.Ext(fileName))+taggedString+path.Ext(fileName))
	return
}
