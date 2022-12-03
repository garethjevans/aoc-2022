package day02_test

import (
	"github.com/garethjevans/aoc-2022/day02"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var part1Rules = map[string]int{
	"A Y": 8, "B Z": 9, "C X": 7, //win
	"A Z": 3, "B X": 1, "C Y": 2, //loss
	"A X": 4, "B Y": 5, "C Z": 6, //draw
}

var part2Rules = map[string]int{
	"A Z": 8, "B Z": 9, "C Z": 7, //win
	"A X": 3, "B X": 1, "C X": 2, //loss
	"A Y": 4, "B Y": 5, "C Y": 6, //draw
}

func TestCanParsePart1(t *testing.T) {
	rounds, err := day02.ParseInput("testdata/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(rounds))

	values := lo.Map[day02.Round, int](rounds, func(x day02.Round, index int) int {
		return part1Rules[x.Game]
	})

	assert.Equal(t, 15, lo.Sum(values))
}

func TestCanParsePart2(t *testing.T) {
	rounds, err := day02.ParseInput("testdata/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(rounds))

	values := lo.Map[day02.Round, int](rounds, func(x day02.Round, index int) int {
		return part2Rules[x.Game]
	})

	assert.Equal(t, 12, lo.Sum(values))
}

func TestCanParseInputPart1(t *testing.T) {
	rounds, err := day02.ParseInput("testdata/input.txt")
	assert.NoError(t, err)
	assert.Equal(t, 2500, len(rounds))

	values := lo.Map[day02.Round, int](rounds, func(x day02.Round, index int) int {
		return part1Rules[x.Game]
	})

	t.Logf("Part 1> %d", lo.Sum(values))
}

func TestCanParseInputPart2(t *testing.T) {
	rounds, err := day02.ParseInput("testdata/input.txt")
	assert.NoError(t, err)
	assert.Equal(t, 2500, len(rounds))

	values := lo.Map[day02.Round, int](rounds, func(x day02.Round, index int) int {
		return part2Rules[x.Game]
	})

	t.Logf("Part 2> %d", lo.Sum(values))
}
