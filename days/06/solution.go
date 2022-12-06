package main

import (
	"fmt"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
)

func getInput() string {
	return input.Read("days/06/input.txt")
}

func part1() int {
	c := getInput()

	for i := 0; i < len(c)-3; i++ {
		x := []rune(c[i : i+3])
		fmt.Println(x)
		dup := false
		for i, v1 := range x {
			if array.None(x, func(v2 rune, j int) bool {
				return v1 == v2 && i != j
			}) {
				dup = true
			}
		}
		if !dup {
			return i + 4
		}
	}
	return -1
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
