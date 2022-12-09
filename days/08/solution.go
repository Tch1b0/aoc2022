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
	x := make([][]int, len(arr[0]))
	for _, row := range arr {
		for j, col := range row {
			x[j] = append(x[j], col)
		}
	}

	return x
}

func viewScore(grid [][]int, y, x int) int {
	prod := 1
	val := grid[y][x]

	v := 0
	for k := x + 1; k < len(grid[0]); k++ {
		v += 1
		if grid[y][k] >= val {
			break
		}
	}
	prod *= v

	v = 0
	for k := x - 1; k >= 0; k-- {
		v += 1
		if grid[y][k] >= val {
			break
		}
	}
	prod *= v

	v = 0
	for k := y + 1; k < len(grid); k++ {
		v += 1
		if grid[k][x] >= val {
			break
		}
	}
	prod *= v

	v = 0
	for k := y - 1; k >= 0; k-- {
		v += 1
		if grid[k][x] >= val {
			break
		}
	}
	prod *= v

	return prod
}

func part1() int {
	grid := getInput()
	xGrid := swapRowCol(grid)
	sum := len(grid)*2 + (len(xGrid)-2)*2
	for i, row := range grid {
		if i == 0 || i == len(grid)-1 {
			continue
		}

		for j, col := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}

			isLower := func(v, _ int) bool { return v < col }
			if array.All(grid[i][:j], isLower) || array.All(grid[i][j+1:], isLower) || array.All(xGrid[j][:i], isLower) || array.All(xGrid[j][i+1:], isLower) {
				sum += 1
			}
		}
	}
	return sum
}

func part2() int {
	scores := []int{}

	grid := getInput()

	for i, row := range grid {
		if i == 0 || i == len(grid)-1 {
			continue
		}

		for j, _ := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}

			scores = append(scores, viewScore(grid, i, j))
		}
	}

	return math.Max(scores)
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
