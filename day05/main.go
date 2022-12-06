package day05

import (
	"github.com/garethjevans/aoc-2022/day04"
	"github.com/garethjevans/aoc-2022/pkg/file"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func GetCommands(filename string) ([]Command, error) {
	lines, err := file.AsLines(filename)
	if err != nil {
		return nil, err
	}

	filteredLines := lo.Filter[string](lines, func(x string, index int) bool {
		return strings.HasPrefix(x, "move")
	})

	commands := lo.Map[string, Command](filteredLines, func(i string, index int) Command {
		re := regexp.MustCompile(`\d+`)
		v := re.FindAllString(i, -1)
		return Command{
			Move: day04.Must(v[0]),
			From: day04.Must(v[1]),
			To:   day04.Must(v[2]),
		}
	})
	return commands, nil
}

type Command struct {
	Move int
	From int
	To   int
}
