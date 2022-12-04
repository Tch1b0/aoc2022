package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	"github.com/Tch1b0/polaris/strings"
)

func getInput() [][]math.Span[int] {
	return input.Process("./days/04/input.txt", func(str string) [][]math.Span[int] {
		groups := [][]math.Span[int]{}
		lines := stdstrings.Split(stdstrings.ReplaceAll(str, "\r", ""), "\n")

		for _, line := range lines {
			elves := make([]math.Span[int], 2)

			for i, elve := range stdstrings.Split(line, ",") {
				splitted := stdstrings.Split(elve, "-")

				a, err := strings.Atoi(splitted[0])
				if err != nil {
					panic(err)
				}

				b, err := strings.Atoi(splitted[1])
				if err != nil {
					panic(err)
				}

				elves[i] = math.Span[int]{From: a, To: b}
			}

			groups = append(groups, elves)
		}

		return groups
	})
}

func part1() int {
	sum := 0
	c := getInput()
	for _, group := range c {
		a, b := group[0], group[1]
		if a.ContainsSpan(b) || b.ContainsSpan(a) {
			sum += 1
		}
	}
	return sum
}

func part2() int {
	sum := 0
	c := getInput()
	for _, group := range c {
		a, b := group[0], group[1]
		if a.Overlaps(b) {
			sum += 1
		}
	}
	return sum
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
