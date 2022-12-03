package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
)

func getInput() []string {
	return input.Process("./days/03/input.txt", func(str string) []string {
		return stdstrings.Split(str, "\n")
	})
}

func IsLower(c rune) bool {
	return c >= 97 && c <= 122
}

func IsUpper(c rune) bool {
	return c >= 65 && c <= 90
}

func calcScore(c rune) int {
	if IsUpper(c) {
		return int(c) - 64 + 26
	} else if IsLower(c) {
		return int(c) - 96
	}

	return 0
}

func part1() int {
	sum := 0

	for _, item := range getInput() {
		half := len(item) / 2
		l, r := item[:half], item[half:]

		for _, c := range l {
			if stdstrings.Contains(r, string(c)) {
				sum += calcScore(c)
				break
			}
		}
	}

	return sum
}

func part2() int {
	sum := 0
	in := getInput()
	for i, item := range in {
		// only let through the first of a group of 3
		if i%3 != 0 {
			continue
		}

		for _, c := range item {
			cs := string(c)

			// check if the following two strings contain the character
			// assumes that len(in) >= i + 3
			if stdstrings.Contains(in[i+1], cs) && stdstrings.Contains(in[i+2], cs) {
				sum += calcScore(c)
				break
			}
		}
	}

	return sum
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
