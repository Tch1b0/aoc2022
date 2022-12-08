package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/strings"
)

func getInput() [][]int {
	return input.Process("days/08/input.txt", func(str string) [][]int {
		splitted := stdstrings.Split(str, "\n")
		grid := [][]int{}
		for _, v := range splitted {
			x := []int{}
			for _, c := range v {
				m, err := strings.Atoi(string(c))
				if err != nil {
					panic(err)
				}
				x = append(x, m)
			}
			grid = append(grid, x)
		}

		return grid
	})
}

func swapRowCol(arr [][]int) [][]int {
	x := arr
	for i, row := range arr {
		for j, col := range row {
			x[j][i] = col
		}
	}

	return x
}

func part1() int {
	grid := getInput()
	xGrid := swapRowCol(grid)
	sum := len(grid)*2 + len(grid[0])*2 - 4
	for i, row := range grid {
		if i == 0 || i == len(grid)-1 {
			continue
		}

		for j, col := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}

			isLower := func(v, _ int) bool { return v < col }
			if array.All(grid[i][:j-1], isLower) || array.All(grid[i][j+1:], isLower) || array.All(xGrid[i][:j-1], isLower) || array.All(xGrid[i][j+1:], isLower) {
				sum += 1
			}
		}
	}
	return sum
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
