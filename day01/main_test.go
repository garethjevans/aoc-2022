package day01_test

import (
	"github.com/garethjevans/aoc-2022/day01"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestCanFindTopElf(t *testing.T) {
	results, err := day01.GetElfs("testdata/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, 5, len(results))

	max := lo.MaxBy(results, func(item day01.Elf, max day01.Elf) bool {
		return item.Sum() > max.Sum()
	})

	assert.Equal(t, 24000, max.Sum())
}

func TestCanFindTop3Elf(t *testing.T) {
	results, err := day01.GetElfs("testdata/test.txt")
	assert.NoError(t, err)
	assert.Equal(t, 5, len(results))

	elfValues := lo.Map[day01.Elf, int](results, func(x day01.Elf, index int) int {
		return x.Sum()
	})

	sort.Ints(elfValues)

	lastThree := elfValues[len(elfValues)-3:]

	assert.Equal(t, 45000, lo.Sum(lastThree))
}

func TestCanFindTopElfOnInput(t *testing.T) {
	results, err := day01.GetElfs("testdata/input.txt")
	assert.NoError(t, err)

	max := lo.MaxBy(results, func(item day01.Elf, max day01.Elf) bool {
		return item.Sum() > max.Sum()
	})

	t.Logf("Part 1> %d", max.Sum())
}

func TestCanFindTop3ElfOnInput(t *testing.T) {
	results, err := day01.GetElfs("testdata/input.txt")
	assert.NoError(t, err)

	elfValues := lo.Map[day01.Elf, int](results, func(x day01.Elf, index int) int {
		return x.Sum()
	})

	sort.Ints(elfValues)

	lastThree := elfValues[len(elfValues)-3:]

	t.Logf("Part 2> %d", lo.Sum(lastThree))
}
