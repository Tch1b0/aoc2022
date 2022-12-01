package main

import (
	"fmt"

	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	"github.com/Tch1b0/polaris/strings"
)

func getInput() [][]int {
	return input.Process("days/01/input.txt", func(str string) [][]int {
		elves := stdstrings.Split(str, "\n\n")
		total := [][]int{}
		for _, elve := range elves {
			e := stdstrings.Split(elve, "\n")
			t := []int{}
			for _, calorie := range e {
				i, err := strings.Atoi(calorie)
				if err != nil {
					panic(err)
				}
				t = append(t, i)
			}
			total = append(total, t)
		}
		return total
	})
}

func part1() int {
	c := getInput()
	sums := array.Map(c, func(elve []int, _ int) int {
		return math.Sum(elve)
	})
	return math.Max(sums)
}

func part2() int {
	c := getInput()
	tops := []int{}
	sums := array.Map(c, func(elve []int, _ int) int {
		return math.Sum(elve)
	})
	tops = math.Sort(sums)
	return math.Sum(tops[len(tops)-3:])
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
