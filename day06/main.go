package day06

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

func FindMarker(input string) int {
	c := strings.Split(input, "")

	length := 4

	for x := 0; x <= len(c)-length; x++ {
		sub := c[x : x+length]
		fmt.Printf("looking at %+v\n", sub)
		if len(lo.Uniq(sub)) == length {
			fmt.Printf("unique %+v [%d->%d]\n", sub, x, x+length)
			return x + length
		}
	}
	return -1
}

func FindMessage(input string) int {
	c := strings.Split(input, "")

	length := 14

	for x := 0; x <= len(c)-length; x++ {
		sub := c[x : x+length]
		fmt.Printf("looking at %+v\n", sub)
		if len(lo.Uniq(sub)) == length {
			fmt.Printf("unique %+v [%d->%d]\n", sub, x, x+length)
			return x + length
		}
	}
	return -1
}
