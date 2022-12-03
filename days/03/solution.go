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

func part1() int {
	sum := 0
	for _, item := range getInput() {
		l, r := item[:len(item)/2], item[len(item)/2:]
		for _, c := range l {
			if stdstrings.Contains(r, string(c)) {
				if IsUpper(c) {
					sum += int(c) - 64 + 26
				} else {
					sum += int(c) - 96
				}
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
		if i%3 != 0 {
			continue
		}

		for _, c := range item {
			cs := string(c)
			if stdstrings.Contains(in[i+1], cs) && stdstrings.Contains(in[i+2], cs) {
				if IsUpper(c) {
					sum += int(c) - 64 + 26
				} else {
					sum += int(c) - 96
				}
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
