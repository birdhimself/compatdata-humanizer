package writer

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/birdhimself/compatdata-humanizer/internal/cli"
	"github.com/birdhimself/compatdata-humanizer/internal/config"
	"github.com/birdhimself/compatdata-humanizer/internal/steam"
)

var ErrAbortedByUser = errors.New("aborted by user")

// ResetOutputPath recursively deletes and recreates the output path.
func ResetOutputPath() error {
	if !config.Get().SkipConfirmation && !cli.Confirm(fmt.Sprintf("Overwrite output path \"%s\" and start", config.Get().OutputPath)) {
		return ErrAbortedByUser
	}

	if err := os.RemoveAll(config.Get().OutputPath); err != nil {
		return err
	}

	if err := os.MkdirAll(config.Get().OutputPath, 0755); err != nil {
		return err
	}

	return nil
}

// CreateLink creates a human-readable symlink in the output path to the
// compatdata path of the app given in appInfo.
func CreateLink(appInfo steam.AppInfo, libPath string) (string, error) {
	compatPath := steam.CompatDataPath(libPath, appInfo.AppId)

	info, err := os.Stat(compatPath)

	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return "", fmt.Errorf("\"%s\" is not a directory", compatPath)
	}

	linkPath := path.Join(config.Get().OutputPath, appInfo.InstallDir)

	for {
		_, err := os.Stat(linkPath)

		if err != nil {
			if os.IsNotExist(err) {
				break
			}
			return "", err
		}

		linkPath += "_"
	}

	if err := os.Symlink(compatPath, linkPath); err != nil {
		return "", err
	}

	return linkPath, nil
}
