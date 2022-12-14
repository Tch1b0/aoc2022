package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
)

type Signal struct {
	Left  []any
	Right []any
}

func isInt(v any) bool {
    _, ok := v.(int)
    return ok
}

func (s Signal) countCompare() bool {
    var sm []int
    sm, ok := smaller.([]int)
    if !ok {
        sm = []int{smaller.(int)}
    }
    
    var bm []int
    bm, ok = bigger.([]int)
    if !ok {
        bm = []int{bigger.(int)}
    }
    
    for i := 0; i < math.Min([]int{len(sm), len(bm)}) {

    }
}

func getInput() []Signal {
	return input.Process("./days/13/input.txt", func(str string) []Signal {
		blocks := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n\n")
		signals := []Signal{}

		for _, block := range blocks {
			s := Signal{}
			for i, line := range strings.Split(block, "\n") {
				var v []any
				if err := json.Unmarshal([]byte(line), &v); err != nil {
					panic(err)
				}

				if i == 0 {
					s.Left = v
				} else {
					s.Right = v
				}
			}
			signals = append(signals, s)
		}

		return signals
	})
}

func part1() int {
	fmt.Println(getInput())
	return -1
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
