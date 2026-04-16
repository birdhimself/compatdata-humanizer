package config

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/birdhimself/compatdata-humanizer/internal/cli"
)

var (
	once sync.Once

	config Config
)

type Config struct {
	OutputPath       string `toml:"output_path" comment:"Path where the symlinks will be created in"`
	SkipConfirmation bool   `toml:"skip_confirmation" comment:"Whether to skip the confirmation if the output path should be reset (deleted)"`
}

func Get() Config {
	once.Do(func() {
		cli.Title("Configuration:")

		if file, err := configFile(); err == nil {
			cli.Info("Using \"%s\"", file)
		}

		if err := processFile(); err != nil {
			cli.Error("Failed processing config file: %v", err)
		} else {
			cli.Success("Configuration file processed")
		}

		parseArgs()

		if err := defaults(&config); err != nil {
			fmt.Printf("Error applying default config: %v\n", err)
		}

		cli.Success("Loaded")
	})

	return config
}

func defaults(config *Config) error {
	if config.OutputPath == "" {
		home, err := os.UserHomeDir()

		if err != nil {
			return err
		}

		config.OutputPath = path.Join(home, "SteamCompats")
	}

	return nil
}
