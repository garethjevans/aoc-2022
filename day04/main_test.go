package day04

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	s, err := ParseFile("testdata/test.txt")
	assert.NoError(t, err)

	assert.Equal(t, 6, len(s))
	filtered := lo.Filter[IdSet](s, func(i IdSet, _ int) bool {
		return i.Overlap()
	})

	assert.Equal(t, 2, len(filtered))
}

func TestParseFileFromInput(t *testing.T) {
	s, err := ParseFile("testdata/input.txt")
	assert.NoError(t, err)

	assert.Equal(t, 1000, len(s))
	filtered := lo.Filter[IdSet](s, func(i IdSet, _ int) bool {
		return i.Overlap()
	})

	t.Logf("Part 1> %d", len(filtered))
}

func TestParseFilePart2(t *testing.T) {
	s, err := ParseFile("testdata/test.txt")
	assert.NoError(t, err)

	assert.Equal(t, 6, len(s))
	filtered := lo.Filter[IdSet](s, func(i IdSet, _ int) bool {
		return i.PartialOverlap()
	})

	assert.Equal(t, 4, len(filtered))
}

func TestParseFileFromInput2(t *testing.T) {
	s, err := ParseFile("testdata/input.txt")
	assert.NoError(t, err)

	assert.Equal(t, 1000, len(s))
	filtered := lo.Filter[IdSet](s, func(i IdSet, _ int) bool {
		return i.PartialOverlap()
	})

	t.Logf("Part 2> %d", len(filtered))
}
