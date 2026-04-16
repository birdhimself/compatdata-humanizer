package steam

import (
	"fmt"
	"os"

	"github.com/andygrunwald/vdf"
)

func openAndParseVdf(vdfPath string) (map[string]interface{}, error) {
	f, err := os.Open(vdfPath)

	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("error closing file \"%s\": %s\n", vdfPath, err)
		}
	}(f)

	return vdf.NewParser(f).Parse()
}

func getVdfObject(data map[string]interface{}, key string) (map[string]interface{}, error) {
	if data[key] == nil {
		return nil, fmt.Errorf("vdf root key %s not found", key)
	}

	obj, ok := data[key].(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("%s is not an object", key)
	}

	return obj, nil
}

func getVdfString(data map[string]interface{}, key string) (string, error) {
	x, ok := data[key]

	if !ok {
		return "", fmt.Errorf("%s not found", key)
	}

	p, ok := x.(string)

	if !ok {
		return "", fmt.Errorf("%s is not a string", key)
	}

	return p, nil
}
