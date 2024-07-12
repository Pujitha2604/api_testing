package analyze

import (
	"os"
	"path/filepath"
	"strings"
)

func getAllGoFiles(dir string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return goFiles, nil
}
