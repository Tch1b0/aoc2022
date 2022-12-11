package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	s "github.com/Tch1b0/polaris/strings"
)

var vec0 = math.Vector2[int]{X: 0, Y: 0}

func ArrayContains(arr []math.Vector2[int], val math.Vector2[int]) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

type Knot struct {
	Child   *Knot
	Pos     math.Vector2[int]
	Visited []math.Vector2[int]
}

func (k *Knot) Move(vec math.Vector2[int]) {
	k.Pos = k.Pos.Add(vec)

	if !ArrayContains(k.Visited, k.Pos) {
		k.Visited = append(k.Visited, k.Pos)
	}

	if k.Child != nil {
		k.Child.Adjust(*k)
	}
}

func (k *Knot) Adjust(parent Knot) {
	if int(k.Pos.DistanceTo(parent.Pos)) <= 1 {
		return
	}

	dir := vec0

	if k.Pos.X < parent.Pos.X {
		dir.X += 1
	} else if k.Pos.X > parent.Pos.X {
		dir.X -= 1
	}

	if k.Pos.Y < parent.Pos.Y {
		dir.Y += 1
	} else if k.Pos.Y > parent.Pos.Y {
		dir.Y -= 1
	}

	k.Move(dir)
}

func (k Knot) Array() []Knot {
	knots := []Knot{}

	knot := &k
	for knot != nil {
		knots = append(knots, *knot)
		knot = knot.Child
	}

	return knots
}

func (k Knot) DrawField() {
	knots := k.Array()

	for y := 5; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			f := array.GetFirst(knots, func(v Knot, _ int) bool { return v.Pos == math.Vector2[int]{X: x, Y: y} })

			if f != nil {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

func NewKnot(pos math.Vector2[int], child *Knot) *Knot {
	return &Knot{
		Pos:     pos,
		Child:   child,
		Visited: []math.Vector2[int]{vec0},
	}
}

type Instruction struct {
	Direction rune
	Length    int
}

func (i Instruction) DirVec() math.Vector2[int] {
	switch i.Direction {
	case 'R':
		return math.Vector2[int]{X: 1, Y: 0}
	case 'L':
		return math.Vector2[int]{X: -1, Y: 0}
	case 'U':
		return math.Vector2[int]{X: 0, Y: 1}
	case 'D':
		return math.Vector2[int]{X: 0, Y: -1}
	}

	panic(fmt.Errorf("unrecognized direction: '%c'", i.Direction))
}

func (i Instruction) String() string {
	return fmt.Sprintf("%c %d", i.Direction, i.Length)
}

func getInput() []Instruction {
	return input.Process("days/09/input.txt", func(str string) []Instruction {
		lines := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n")
		ins := []Instruction{}

		for _, line := range lines {
			splitted := strings.Split(line, " ")
			dir := splitted[0]
			length, err := s.Atoi(splitted[1])
			if err != nil {
				panic(err)
			}

			ins = append(ins, Instruction{Direction: rune(dir[0]), Length: length})
		}

		return ins
	})
}

func part1() int {
	head := *NewKnot(vec0, NewKnot(vec0, nil))

	for _, in := range getInput() {
		dir := in.DirVec()
		for i := 0; i < in.Length; i++ {
			head.Move(dir)
		}
	}

	return len(head.Child.Visited)
}

func part2() int {
	var head *Knot = nil
	for i := 0; i < 10; i++ {
		head = NewKnot(vec0, head)
	}

	for _, in := range getInput() {
		dir := in.DirVec()
		for i := 0; i < in.Length; i++ {
			head.Move(dir)
		}
	}

	arr := head.Array()

	return len(arr[len(arr)-1].Visited)
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
