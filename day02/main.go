package day02

import (
	"os"
	"strings"
)

type Round struct {
	Game string
}

func ParseInput(filename string) ([]Round, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	if err != nil {
		return nil, err
	}

	var rounds []Round
	for _, line := range lines {
		rounds = append(rounds, Round{Game: line})
	}

	return rounds, nil
}
