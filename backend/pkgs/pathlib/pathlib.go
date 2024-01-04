// Package pathlib provides a way to safely create a file path without overwriting any existing files.
package pathlib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type dirReaderFunc func(name string) []string

var dirReader dirReaderFunc = func(directory string) []string {
	f, err := os.Open(directory)
	if err != nil {
		return nil
	}
	defer func() { _ = f.Close() }()

	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil
	}
	return names
}

func hasConflict(path string, neighbors []string) bool {
	filename := strings.ToLower(filepath.Base(path))

	for _, n := range neighbors {
		if strings.ToLower(n) == filename {
			return true
		}
	}
	return false
}

// Safe will take a destination path and return a validated path that is safe to use.
// without overwriting any existing files. If a conflict exists, it will append a number
// to the end of the file name. If the parent directory does not exist this function will
// return the original path.
func Safe(path string) string {
	parent := filepath.Dir(path)

	neighbors := dirReader(parent)
	if neighbors == nil {
		return path
	}

	if hasConflict(path, neighbors) {
		ext := filepath.Ext(path)

		name := strings.TrimSuffix(filepath.Base(path), ext)

		for i := 1; i < 1000; i++ {
			newName := fmt.Sprintf("%s (%d)%s", name, i, ext)
			newPath := filepath.Join(parent, newName)
			if !hasConflict(newPath, neighbors) {
				return newPath
			}
		}
	}

	return path
}
