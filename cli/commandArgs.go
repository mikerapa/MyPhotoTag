package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func ParseCommandLine(commandLineArgs []string) (string, string, string, error) {
	app := kingpin.New("MyPhotoTag", "My Photo Tag")
	var (
		photoFilePath  = app.Flag("photoPath", "File path of original photo or path to a folder of original photos.").Short('p').Required().String()
		tagFilePath    = app.Flag("tagPath", "File path to the tag png file").Short('t').Required().String()
		outputPath     = app.Flag("outputPath", "Folder path to place the tagged photo(s)").Short('o').String()
		logLevelString = app.Flag("logLevel", "Log Level (error, warning, info, debug, trace)").Short('l').String()
	)

	_, err := app.Parse(commandLineArgs)

	if err != nil {
		ConsoleLogger.Error("Error while parsing command line:", err)
	}

	ConsoleLogger.Level = GetLogLevelFromString(*logLevelString)

	ConsoleLogger.Info("Command line args:", commandLineArgs)
	ConsoleLogger.Info("Command Argument photoPath: ", *photoFilePath)
	ConsoleLogger.Info("Command Argument tagPath: ", *tagFilePath)
	ConsoleLogger.Info("Command Argument outputFilePath: ", *outputPath)
	ConsoleLogger.Trace("Done parsing command line arguments")

	return *photoFilePath, *tagFilePath, *outputPath, err
}
