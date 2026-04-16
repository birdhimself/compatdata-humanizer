package steam

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"

	"github.com/birdhimself/compatdata-humanizer/internal/cli"
)

var homeLibraryPaths = []string{
	".steam/steam/steamapps/libraryfolders.vdf",
	".local/share/Steam/steamapps/libraryfolders.vdf",
	".var/app/com.valvesoftware.Steam/data/Steam/steamapps/libraryfolders.vdf",
}

const libraryFoldersKey = "libraryfolders"
const appStateKey = "AppState"

func ensureAndResolvePath(p string, ensureDir bool) (string, error) {
	var err error
	var i os.FileInfo

	if i, err = os.Stat(p); err != nil {
		return p, err
	} else if ensureDir && !i.IsDir() {
		return p, fmt.Errorf("%s is not a directory", p)
	}

	if p, err = filepath.EvalSymlinks(p); err != nil {
		return p, err
	}

	if p, err = filepath.Abs(p); err != nil {
		return p, err
	}

	return p, nil
}

// LibraryFiles checks common locations for existing libraryfolders.vdf files
// and returns a slice of unique, fully resolved paths to them.
func LibraryFiles() ([]string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	var existing []string

	for _, p := range homeLibraryPaths {
		raw := path.Join(home, p)

		if raw, err = ensureAndResolvePath(raw, false); err != nil {
			continue
		}

		if slices.Contains(existing, raw) {
			continue
		}

		existing = append(existing, raw)
	}

	return existing, nil
}

// ParseLibraryFolders takes the vdfPath to a libraryfolders.vdf file, parses
// it, and returns the fully resolved path to all libraries contained within
// that actually exist.
func ParseLibraryFolders(vdfPath string) ([]string, error) {
	data, err := openAndParseVdf(vdfPath)

	if err != nil {
		return nil, err
	}

	libFolders, err := getVdfObject(data, libraryFoldersKey)

	if err != nil {
		return nil, err
	}

	var paths []string

	for k := range libFolders {
		libFolder, err := getVdfObject(libFolders, k)

		if err != nil {
			cli.Warning("Library %s is not an object\n", k)
			continue
		}

		p, err := getVdfString(libFolder, "path")

		if err != nil {
			cli.Warning("Library %s does not have a path\n", k)
			continue
		}

		if p, err = ensureAndResolvePath(p, true); err != nil {
			cli.Warning("Can't resolve path \"%s\" of library folder %s: %v", p, k, err)
			continue
		}

		if slices.Contains(paths, p) {
			continue
		}

		paths = append(paths, p)
	}

	return paths, nil
}

// AppsPath returns the "steamapps" path of a given libPath.
func AppsPath(libPath string) string {
	return path.Join(libPath, "steamapps")
}

// CompatDataPath returns the "steamapps/compatdata/{appId}" path of a given
// libPath.
func CompatDataPath(libPath string, appId string) string {
	return path.Join(AppsPath(libPath), "compatdata", appId)
}

// AppInfo contains information about a Steam app.
type AppInfo struct {
	AppId      string
	Name       string
	InstallDir string
}

// Human returns a simplified string representation of the struct for
// human-readable logging.
func (a *AppInfo) Human() string {
	return fmt.Sprintf("\"%s\" (%s)", a.Name, a.AppId)
}

// AppInfoFromAcf extracts AppInfo from acfPath.
func AppInfoFromAcf(acfPath string) (AppInfo, error) {
	acf, err := openAndParseVdf(acfPath)

	if err != nil {
		return AppInfo{}, err
	}

	appState, err := getVdfObject(acf, appStateKey)

	if err != nil {
		return AppInfo{}, err
	}

	appId, err := getVdfString(appState, "appid")

	if err != nil {
		return AppInfo{}, fmt.Errorf("could not find appid for \"%s\"", path.Base(acfPath))
	}

	name, err := getVdfString(appState, "name")

	if err != nil {
		return AppInfo{}, fmt.Errorf("could not find name for \"%s\"", path.Base(acfPath))
	}

	installDir, err := getVdfString(appState, "installdir")

	if err != nil {
		return AppInfo{}, fmt.Errorf("could not find installdir for \"%s\"", path.Base(acfPath))
	}

	return AppInfo{
		AppId:      appId,
		Name:       name,
		InstallDir: installDir,
	}, nil
}
