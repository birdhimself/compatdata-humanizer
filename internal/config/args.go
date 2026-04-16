package config

import "flag"

func parseArgs() {
	var (
		outputPath       string
		skipConfirmation bool
	)

	flag.StringVar(&outputPath, "o", "", "Path where the symlinks will be created in")
	flag.BoolVar(&skipConfirmation, "y", false, "Whether to skip the confirmation if the output path should be reset (deleted)")
	flag.Parse()

	if outputPath != "" {
		config.OutputPath = outputPath
	}

	if skipConfirmation {
		config.SkipConfirmation = true
	}
}
