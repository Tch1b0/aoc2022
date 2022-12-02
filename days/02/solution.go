package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
)

func getInput() [][]rune {
	return input.Process("days/02/input.txt", func(str string) [][]rune {
		lines := stdstrings.Split(str, "\n")
		arr := [][]rune{}
		for _, line := range lines {
			arr = append(arr, []rune{rune(line[0]), rune(line[2])})
		}
		return arr
	})
}

const (
	LOST = iota
	DRAW
	WON
)

var (
	loseToChar = map[rune]rune{
		'A': 'C',
		'B': 'A',
		'C': 'B',
	}
	beatChar = map[rune]rune{
		'A': 'B',
		'B': 'C',
		'C': 'A',
	}
)

func part1() int {
	c := getInput()
	sum := 0
	for _, game := range c {
		var state int
		game[1] -= 23
		if game[0] == game[1] {
			fmt.Println(string(game[1]))
			state = DRAW
		} else if game[0] == 'A' && game[1] == 'C' || game[0] == 'C' && game[1] == 'B' || game[0] == 'B' && game[1] == 'A' {
			state = LOST
		} else {
			state = WON
		}
		if state == DRAW {
			sum += 3
		} else if state == WON {
			sum += 6
		}
		sum += int(game[1]) - 64
	}
	return sum
}

func part2() int {
	c := getInput()
	sum := 0
	for _, game := range c {
		var chr rune
		enemy := game[0]
		outc := game[1]
		if outc == 'X' {
			chr = loseToChar[enemy]
		} else if outc == 'Y' {
			chr = enemy
			sum += 3
		} else {
			chr = beatChar[enemy]
			sum += 6
		}
		sum += int(chr) - 64
	}
	return sum
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
