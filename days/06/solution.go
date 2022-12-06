package main

import (
	"fmt"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
)

func getInput() string {
	return input.Read("days/06/input.txt")
}

func unique[T comparable](arr []T) []T {
	n := []T{}
	for _, v := range arr {
		if !array.Some(n, func(x T, _ int) bool { return v == x }) {
			n = append(n, v)
		}
	}

	return n
}

func part1(n int) int {
	c := getInput()

	for i := 0; i < len(c)-(n-1); i++ {
		x := []rune(c[i : i+n])
		if len(unique(x)) == len(x) {
			return i + n
		}
	}
	return -1
}

func part2() int {
	return part1(14)
}

func main() {
	fmt.Printf("PART 1: %d\n", part1(4))

	fmt.Printf("PART 2: %d\n", part2())
}
