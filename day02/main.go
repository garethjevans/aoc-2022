package day02

import (
	"github.com/garethjevans/aoc-2022/pkg/file"
)

type Round struct {
	Game string
}

func ParseInput(filename string) ([]Round, error) {
	lines, err := file.AsLines(filename)
	if err != nil {
		return nil, err
	}

	var rounds []Round
	for _, line := range lines {
		rounds = append(rounds, Round{Game: line})
	}

	return rounds, nil
}
