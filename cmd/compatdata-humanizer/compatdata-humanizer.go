package main

import (
	"errors"
	"log"
	"os"
	"path"
	"slices"

	"github.com/birdhimself/compatdata-humanizer/internal/cli"
	"github.com/birdhimself/compatdata-humanizer/internal/config"
	"github.com/birdhimself/compatdata-humanizer/internal/steam"
	"github.com/birdhimself/compatdata-humanizer/internal/writer"
)

var version string

func main() {
	cli.Title("Compatdata Humanizer")
	cli.Info("Version %s", version)

	_ = config.Get()

	libFiles, err := steam.LibraryFiles()

	if err != nil {
		log.Fatal(err)
	}

	cli.Title("Detected library files:")
	cli.BulletList(libFiles)

	var libraries []string

	cli.Title("Detected library folders:")

	for _, libFile := range libFiles {
		lfs, err := steam.ParseLibraryFolders(libFile)

		if err != nil {
			cli.Warning("Could not parse \"%s\": %v", libFile, err)
			continue
		}

		libraries = slices.Concat(libraries, lfs)
	}

	libraries = slices.Compact(libraries)

	cli.BulletList(libraries)

	if err := writer.ResetOutputPath(); err != nil {
		if errors.Is(err, writer.ErrAbortedByUser) {
			cli.Warning("Aborted by user.")
			os.Exit(0)
			return
		}

		panic(err)
	}

	cli.Title("Processing:")

	for _, library := range libraries {
		els, err := os.ReadDir(steam.AppsPath(library))

		if err != nil {
			cli.Warning("Could not read library \"%s\": %v", library, err)
			continue
		}

		for _, el := range els {
			if path.Ext(el.Name()) != ".acf" {
				continue
			}

			appInfo, err := steam.AppInfoFromAcf(path.Join(steam.AppsPath(library), el.Name()))

			if err != nil {
				cli.Error("Parsing \"%s\" failed: %v", el.Name(), err)
				continue
			}

			linkPath, err := writer.CreateLink(appInfo, library)

			if err != nil {
				cli.Error("Could not create link for %s: %s", appInfo.Human(), err)
				continue
			}

			cli.Success("Created link for %s at \"%s\"", appInfo.Human(), linkPath)
		}
	}
}
