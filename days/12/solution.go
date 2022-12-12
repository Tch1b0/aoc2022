package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
)

type Hill struct {
	Points [][]rune
}

func getInput() Hill {
	return input.Process("./days/12/input.txt", func(str string) Hill {
		h := Hill{}

		for _, v := range strings.Split(str, "\n") {
			row := []rune{}
			for _, c := range v {
				row = append(row, c)
			}
			h.Points = append(h.Points, row)
		}

		return h
	})
}

var directions = []math.Vector2[int]{
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

func getShortestPath(pos math.Vector2[int], h Hill, x int) int {
	cur := h.Points[pos.Y][pos.X]
	for _, dir := range directions {
		npos := pos.Add(dir)
		n := h.Points[npos.Y][npos.X]
		if n == 'E' {
			return x
		} else if n <= cur {
			return getShortestPath(npos, h, x+1)
		} else {
			break
		}
	}

	return -1
}

func part1() int {
	fmt.Println(getShortestPath(math.Vector2[int]{0, 0}, getInput(), 0))
	return -1
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
