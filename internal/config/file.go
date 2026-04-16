package config

import (
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

const configDirName = "compatdata-humanizer"
const configFileName = "config.toml"

func configFile() (string, error) {
	return xdg.ConfigFile(path.Join(configDirName, configFileName))
}

func processFile() error {
	p, err := configFile()

	if err != nil {
		return err
	}

	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		return err
	}

	defaultConfig := new(config)

	if err := defaults(defaultConfig); err != nil {
		return err
	}

	if *defaultConfig != config {
		if err := f.Truncate(0); err != nil {
			return err
		}

		if err := toml.NewEncoder(f).Encode(defaultConfig); err != nil {
			return err
		}
	}

	return nil
}
