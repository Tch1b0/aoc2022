package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	s "github.com/Tch1b0/polaris/strings"
)

type Command struct {
	Name string
	Arg  int
}

func getInput() []Command {
	return input.Process("days/10/input.txt", func(str string) []Command {
		lines := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n")
		cmds := []Command{}
		for _, line := range lines {
			splitted := strings.Split(line, " ")
			if splitted[0] == "noop" {
				cmds = append(cmds, Command{Name: "noop", Arg: 0})
			} else {
				n, err := s.Atoi(splitted[1])
				if err != nil {
					panic(err)
				}
				cmds = append(cmds, Command{Name: "addx", Arg: n})
			}
		}
		return cmds
	})
}

func part1() int {
	cycle := 1
	sum := 0
	x := 1

	appendToSum := func() {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			sum += x * cycle
		}
	}

	for _, cmd := range getInput() {
		cycle += 1
		appendToSum()

		if cmd.Name == "addx" {
			cycle += 1
			x += cmd.Arg
			appendToSum()
		}
	}

	return sum
}

func part2() string {
	cycle := 0
	x := 1

	image := ""

	appendToImage := func() {
		if len(image)%41 == 0 {
			image += "\n"
		}

		v := cycle % 40

		if math.Between(x, v-1, v+1) {
			image += "#"
		} else {
			image += " "
		}
	}

	for _, cmd := range getInput() {
		appendToImage()

		if cmd.Name == "addx" {
			cycle += 1
			appendToImage()
			x += cmd.Arg
		}
		cycle += 1
	}

	return image
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: \n%v\n", part2())
}
