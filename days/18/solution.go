package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	s "github.com/Tch1b0/polaris/strings"
)

func getInput() []math.Vector3[int] {
	return input.Process("./days/18/input.txt", func(str string) []math.Vector3[int] {
		vecs := []math.Vector3[int]{}
		for _, line := range strings.Split(strings.ReplaceAll(str, "\r", ""), "\n") {
			vals := strings.Split(line, ",")

			x, err := s.Atoi(vals[0])
			if err != nil {
				panic(err)
			}

			y, err := s.Atoi(vals[1])
			if err != nil {
				panic(err)
			}

			z, err := s.Atoi(vals[2])
			if err != nil {
				panic(err)
			}

			vecs = append(vecs, math.Vector3[int]{X: x, Y: y, Z: z})
		}

		return vecs
	})
}

func part1() int {
	vals := getInput()
	surf := 0

	for _, val := range vals {
		for i, coord := range val.Array() {
			if array.None(vals, func(o math.Vector3[int], _ int) bool {
				return o.Array()[i] == coord+1
			}) {
				surf += 1
			}

			if array.None(vals, func(o math.Vector3[int], _ int) bool {
				return o.Array()[i] == coord-1
			}) {
				surf += 1
			}
		}
	}

	return surf
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
