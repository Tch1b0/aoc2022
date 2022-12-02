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
        // split elve blocks
		elves := stdstrings.Split(str, "\n\n")
		total := [][]int{}
		for _, block := range elves {
            // split numbers of single elve
			elve := stdstrings.Split(block, "\n")
			values := []int{}
			for _, calorie := range elve {
				value, err := strings.Atoi(calorie)
				if err != nil {
					panic(err)
				}
				values = append(values, value)
			}
			total = append(total, values)
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
	sums := array.Map(c, func(elve []int, _ int) int {
		return math.Sum(elve)
	})

    // Get top 3 by sorting the array
	tops := math.Sort(sums)
	return math.Sum(tops[len(tops)-3:])
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
