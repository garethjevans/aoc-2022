package day01

import (
	"github.com/garethjevans/aoc-2022/pkg/file"
	"github.com/samber/lo"
	"strconv"
)

type Elf struct {
	calories []string
}

func (e Elf) Sum() int {
	return lo.SumBy(e.calories, func(item string) int {
		v, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		return v
	})
}

func GetElfs(filename string) ([]Elf, error) {
	lines, err := file.AsLines(filename)
	if err != nil {
		return nil, err
	}

	var results []Elf
	elf := Elf{}

	lo.ForEach[string](lines, func(x string, _ int) {
		if x == "" {
			results = append(results, elf)
			elf = Elf{}
		} else {
			elf.calories = append(elf.calories, x)
		}
	})

	// there is always room for one more
	results = append(results, elf)

	if err != nil {
		return nil, err
	}
	return results, nil
}
