package file

import (
	"os"
	"strings"
)

func AsLines(filename string) ([]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	if err != nil {
		return nil, err
	}

	return lines, nil
}
