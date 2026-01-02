package util

import (
	"os"
	"strings"
)

func FilenameToLines(filename string) ([]string, error) {
	lines, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(lines), "\n"), nil
}
