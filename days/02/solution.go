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

// enum "GameState"
const (
	LOST = iota
	DRAW
	WON
)

var (
    // which char is needed to loose?
	loseToChar = map[rune]rune{
		'A': 'C',
		'B': 'A',
		'C': 'B',
	}

    // which char is needed to win?
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
        // translate X/Y/Z to A/B/C
		game[1] -= 23
		enemy, me := game[0], game[1]

		var state int
		if enemy == me {
			state = DRAW
		} else if enemy == 'A' && me == 'C' || enemy == 'C' && me == 'B' || enemy == 'B' && me == 'A' {
			state = LOST
		} else {
			state = WON
		}

        // add draw/win score
		if state == DRAW {
			sum += 3
		} else if state == WON {
			sum += 6
		}

        // A/B/C - 64 = 1/2/3
		sum += int(me) - 64
	}

	return sum
}

func part2() int {
	c := getInput()
	sum := 0

	for _, game := range c {
		enemy := game[0]
		outc := game[1]

        var chr rune
		if outc == 'X' {
            // get char that is required to lose
			chr = loseToChar[enemy]
		} else if outc == 'Y' {
			chr = enemy
			sum += 3
		} else {
            // get char that is required to win
			chr = beatChar[enemy]
			sum += 6
		}

        // A/B/C - 64 = 1/2/3
		sum += int(chr) - 64
	}
    
	return sum
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
