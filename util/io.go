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

func FilenameToMap(filename string) ([][]rune, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	results := [][]rune{}

	for i, line := range lines {
		results = append(results, []rune{})
		for _, char := range line {
			results[i] = append(results[i], char)
		}
	}

	return results, nil
}
