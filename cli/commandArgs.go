package cli

import (
	"flag"
)

func ParseCommandParameters() (photoFilePath string, tagFilePath string, outputFilePath string) {
	// get command line values
	flag.StringVar(&photoFilePath, "photoPath", "", "File path to the photo")
	flag.StringVar(&tagFilePath, "tagPath", "", "File path to the tag (PNG file)")
	flag.StringVar(&outputFilePath, "outputPhotoPath", "", "File path to the output photo")
	logLevelString := flag.String("logLevel", "Error", "Log Level (error, warning, info, debug, trace)")
	flag.Parse()

	// output values
	ConsoleLogger.Level = GetLogLevelFromString(*logLevelString)
	ConsoleLogger.Info("Command Argument photoPath: ", photoFilePath)
	ConsoleLogger.Info("Command Argument tagPath: ", tagFilePath)
	ConsoleLogger.Info("Command Argument outputFilePath: ", outputFilePath)
	ConsoleLogger.Trace("Done parsing command line arguments")

	return
}
