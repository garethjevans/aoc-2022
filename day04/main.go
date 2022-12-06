package day04

import (
	"github.com/garethjevans/aoc-2022/pkg/file"
	"github.com/garethjevans/aoc-2022/pkg/iterator"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type IdSet struct {
	Set0 []int
	Set1 []int
}

func (i IdSet) Overlap() bool {
	return lo.Every[int](i.Set0, i.Set1) || lo.Every[int](i.Set1, i.Set0)
}

func (i IdSet) PartialOverlap() bool {
	matches := lo.Intersect[int](i.Set0, i.Set1)
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

		set0 := iterator.Fill(Must(s0[0]), Must(s0[1]))
		set1 := iterator.Fill(Must(s1[0]), Must(s1[1]))

		return IdSet{Set0: set0, Set1: set1}
	}), nil
}

func Must(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
