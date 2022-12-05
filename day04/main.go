package day04

import (
	"github.com/garethjevans/aoc-2022/pkg/file"
	"github.com/garethjevans/aoc-2022/pkg/iterator"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type IdSet struct {
	Start0 int
	End0   int
	Start1 int
	End1   int
}

func (i IdSet) Overlap() bool {
	return (i.Start1 >= i.Start0 && i.End1 <= i.End0) ||
		(i.Start0 >= i.Start1 && i.End0 <= i.End1)
}

func (i IdSet) PartialOverlap() bool {
	set0 := iterator.Fill(i.Start0, i.End0)
	set1 := iterator.Fill(i.Start1, i.End1)

	matches := lo.Intersect[int](set0, set1)
	return len(matches) > 0
}

func ParseFile(filename string) ([]IdSet, error) {
	lines, err := file.AsLines(filename)
	if err != nil {
		return nil, err
	}

	return lo.Map[string, IdSet](lines, func(i string, _ int) IdSet {
		parts := strings.Split(i, ",")
		s0 := strings.Split(parts[0], "-")
		s1 := strings.Split(parts[1], "-")

		return IdSet{Start0: Must(s0[0]), End0: Must(s0[1]), Start1: Must(s1[0]), End1: Must(s1[1])}
	}), nil
}

func Must(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
