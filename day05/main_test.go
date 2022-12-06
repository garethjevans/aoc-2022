package day05

import (
	"fmt"
	"github.com/garethjevans/aoc-2022/pkg/structure"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseTestInput(t *testing.T) {
	stacks := GetTestStack()

	commands, err := GetCommands("testdata/test.txt")
	assert.NoError(t, err)

	max := 3

	DumpStacks(t, max, stacks)

	lo.ForEach[Command](commands, func(item Command, index int) {
		from := stacks[item.From]
		to := stacks[item.To]

		for x := 1; x <= item.Move; x++ {
			v, _ := from.Pop()
			if v == "" {
				panic("empty stack")
			}
			to.Push(v)
		}
	})

	DumpStacks(t, max, stacks)
	message := GetMessage(max, stacks)
	assert.Equal(t, "CMZ", message)
}

func TestParseTestInputPart2(t *testing.T) {
	stacks := GetTestStack()

	commands, err := GetCommands("testdata/test.txt")
	assert.NoError(t, err)

	max := 3

	DumpStacks(t, max, stacks)

	lo.ForEach[Command](commands, func(item Command, index int) {
		from := stacks[item.From]
		to := stacks[item.To]

		var temp []string
		for x := 1; x <= item.Move; x++ {
			v, _ := from.Pop()
			if v == "" {
				panic("empty stack")
			}
			temp = append(temp, v)

		}

		temp = lo.Reverse(temp)

		for _, v := range temp {
			to.Push(v)
		}
	})

	DumpStacks(t, max, stacks)
	message := GetMessage(max, stacks)
	assert.Equal(t, "MCD", message)
}

func TestParseRealInputFromPart1(t *testing.T) {
	stacks := GetInputStack()

	commands, err := GetCommands("testdata/input.txt")
	assert.NoError(t, err)

	max := 9

	DumpStacks(t, max, stacks)

	lo.ForEach[Command](commands, func(item Command, index int) {
		from := stacks[item.From]
		to := stacks[item.To]

		for x := 1; x <= item.Move; x++ {
			v, _ := from.Pop()
			if v == "" {
				panic("empty stack")
			}
			to.Push(v)
		}
	})

	DumpStacks(t, max, stacks)
	message := GetMessage(max, stacks)
	t.Logf("%s", message)
}

func TestParseRealInputFromPart2(t *testing.T) {
	stacks := GetInputStack()

	commands, err := GetCommands("testdata/input.txt")
	assert.NoError(t, err)

	max := 9

	DumpStacks(t, max, stacks)

	lo.ForEach[Command](commands, func(item Command, index int) {
		from := stacks[item.From]
		to := stacks[item.To]

		var temp []string
		for x := 1; x <= item.Move; x++ {
			v, _ := from.Pop()
			if v == "" {
				panic("empty stack")
			}
			temp = append(temp, v)
		}

		temp = lo.Reverse(temp)

		for _, v := range temp {
			to.Push(v)
		}
	})

	DumpStacks(t, max, stacks)
	message := GetMessage(max, stacks)
	t.Logf("%s", message)
}

func GetTestStack() map[int]*structure.Stack {
	stacks := make(map[int]*structure.Stack, 3)

	stack1 := structure.Stack{}
	stack1.Push("Z")
	stack1.Push("N")
	stacks[1] = &stack1

	stack2 := structure.Stack{}
	stack2.Push("M")
	stack2.Push("C")
	stack2.Push("D")
	stacks[2] = &stack2

	stack3 := structure.Stack{}
	stack3.Push("P")
	stacks[3] = &stack3
	return stacks
}

func GetInputStack() map[int]*structure.Stack {
	stacks := make(map[int]*structure.Stack, 9)

	// [N] [G]                     [Q]
	// [H] [B]         [B] [R]     [H]
	// [S] [N]     [Q] [M] [T]     [Z]
	// [J] [T]     [R] [V] [H]     [R] [S]
	// [F] [Q]     [W] [T] [V] [J] [V] [M]
	// [W] [P] [V] [S] [F] [B] [Q] [J] [H]
	// [T] [R] [Q] [B] [D] [D] [B] [N] [N]
	// [D] [H] [L] [N] [N] [M] [D] [D] [B]
	//  1   2   3   4   5   6   7   8   9

	stack1 := structure.Stack([]string{"D", "T", "W", "F", "J", "S", "H", "N"})
	stacks[1] = &stack1
	stack2 := structure.Stack([]string{"H", "R", "P", "Q", "T", "N", "B", "G"})
	stacks[2] = &stack2
	stack3 := structure.Stack([]string{"L", "Q", "V"})
	stacks[3] = &stack3
	stack4 := structure.Stack([]string{"N", "B", "S", "W", "R", "Q"})
	stacks[4] = &stack4
	stack5 := structure.Stack([]string{"N", "D", "F", "T", "V", "M", "B"})
	stacks[5] = &stack5
	stack6 := structure.Stack([]string{"M", "D", "B", "V", "H", "T", "R"})
	stacks[6] = &stack6
	stack7 := structure.Stack([]string{"D", "B", "Q", "J"})
	stacks[7] = &stack7
	stack8 := structure.Stack([]string{"D", "N", "J", "V", "R", "Z", "H", "Q"})
	stacks[8] = &stack8
	stack9 := structure.Stack([]string{"B", "N", "H", "M", "S"})
	stacks[9] = &stack9
	return stacks
}

func GetMessage(max int, stacks map[int]*structure.Stack) string {
	var chars []string
	for x := 1; x <= max; x++ {
		chars = append(chars, stacks[x].Peek())
	}
	return strings.Join(chars, "")
}

func DumpStacks(t *testing.T, max int, stacks map[int]*structure.Stack) {
	t.Logf("--------------------------------------")
	for x := 1; x <= max; x++ {
		fmt.Printf("%d > ", x)
		fmt.Printf("%s\n", strings.Join(*stacks[x], " "))
	}
	t.Logf("--------------------------------------")
}
