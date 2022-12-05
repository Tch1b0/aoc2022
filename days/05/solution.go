package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/strings"
)

func getCargos() [][]rune {
	return input.Process("days/05/input.txt", func(str string) [][]rune {
		staples := [][]rune{}
		fields := stdstrings.Split(str, "\n\n")
		lines := stdstrings.Split(fields[0], "\n")

		for _, field := range lines[:len(lines)-1] {
			i := 0
			for j, v := range field {
				if (j-2)%4 != 0 && v != ' ' {
					i += 1
					continue
				}

				if len(staples)-1 < j {
					staples = append(staples, []rune{})
				}

				staples[i] = append(staples[i], v)
				i += 1
			}
		}

		fmt.Println(staples)

		return array.Map(staples, func(staple []rune, _ int) []rune {
			return array.Reverse(staple)
		})
	})
}

func getInstructions() [][]int {
	return input.Process("days/05/input.txt", func(str string) [][]int {
		lines := stdstrings.Split(stdstrings.Split(str, "\n\n")[1], "\n")
		vals := [][]int{}
		for _, line := range lines {
			line = stdstrings.ReplaceAll(line, "move ", "")
			line = stdstrings.ReplaceAll(line, "from", "")
			line = stdstrings.ReplaceAll(line, "to", "")

			vals = append(vals, array.Map(stdstrings.Split(line, "  "), func(v string, _ int) int {
				i, err := strings.Atoi(stdstrings.ReplaceAll(v, " ", ""))
				if err != nil {
					panic(err)
				}

				return i
			}))
		}

		return vals
	})
}

func part1() []rune {
	cargo := getCargos()
	fmt.Println(cargo)
	ins := getInstructions()

	for _, in := range ins {
		n := in[0] - 1
		fromP, toP := &cargo[in[1]-1], &cargo[in[2]-1]
		from := cargo[in[1]-1]
		*toP = append(*toP, from[len(from)-n:]...)
		*fromP = (*fromP)[:n]
	}

	return array.Map(cargo, func(v []rune, _ int) rune {
		return v[len(v)-1]
	})
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %v\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
