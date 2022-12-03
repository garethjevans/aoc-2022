package day03_test

import (
	"github.com/garethjevans/aoc-2022/day03"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCanParsePart1(t *testing.T) {
	rucksacks, err := day03.ParseInput("testdata/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, 6, len(rucksacks))

	v := lo.Map[day03.Rucksack, int](rucksacks, func(r day03.Rucksack, _ int) int {
		return lo.Sum(lo.Map[string, int](r.InBoth(), func(r string, _ int) int {
			return day03.Rules[r]
		}))
	})

	assert.Equal(t, 157, lo.Sum(v))
}

func TestCanParsePart1WithInput(t *testing.T) {
	rucksacks, err := day03.ParseInput("testdata/input.txt")
	assert.NoError(t, err)
	assert.Equal(t, 300, len(rucksacks))

	v := lo.Map[day03.Rucksack, int](rucksacks, func(r day03.Rucksack, _ int) int {
		return lo.Sum(lo.Map[string, int](r.InBoth(), func(r string, _ int) int {
			return day03.Rules[r]
		}))
	})

	t.Logf("Part 1> %d", lo.Sum(v))
}

func TestCanParsePart2(t *testing.T) {
	rucksacks, err := day03.ParseInput("testdata/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, 6, len(rucksacks))

	groups := lo.Chunk(rucksacks, 3)

	v := lo.Map[[]day03.Rucksack, int](groups, func(group []day03.Rucksack, _ int) int {
		g0 := strings.Split(group[0].Items, "")
		g1 := strings.Split(group[1].Items, "")
		g2 := strings.Split(group[2].Items, "")
		one := lo.Intersect(g0, g1)
		two := lo.Intersect(one, g2)

		u := lo.Uniq(two)
		if len(u) != 1 {
			panic("invalid length")
		}

		return day03.Rules[u[0]]
	})

	assert.Equal(t, 70, lo.Sum(v))
}

func TestCanParseInputPart2(t *testing.T) {
	rucksacks, err := day03.ParseInput("testdata/input.txt")
	assert.NoError(t, err)
	assert.Equal(t, 300, len(rucksacks))

	groups := lo.Chunk(rucksacks, 3)

	v := lo.Map[[]day03.Rucksack, int](groups, func(group []day03.Rucksack, _ int) int {
		g0 := strings.Split(group[0].Items, "")
		g1 := strings.Split(group[1].Items, "")
		g2 := strings.Split(group[2].Items, "")
		one := lo.Intersect(g0, g1)
		two := lo.Intersect(one, g2)

		u := lo.Uniq(two)
		if len(u) != 1 {
			panic("invalid length")
		}

		return day03.Rules[u[0]]
	})

	t.Logf("Part 2> %d", lo.Sum(v))
}
