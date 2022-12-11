package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
	s "github.com/Tch1b0/polaris/strings"
)

const (
	MULTIPLY = iota
	ADD
)

type MathSign int

type MathOperation struct {
	Sign   MathSign
	Number *int
}

func (mo MathOperation) Execute(n int) int {
	var num int
	if mo.Number == nil {
		num = n
	} else {
		num = *mo.Number
	}

	if mo.Sign == MULTIPLY {
		return n * num
	} else {
		return n + num
	}
}

func (mo MathOperation) String() string {
	var sign string
	if mo.Sign == MULTIPLY {
		sign = "MULTIPLY"
	} else {
		sign = "ADD"
	}

	var num string
	if mo.Number == nil {
		num = "old"
	} else {
		num = s.Itoa(*mo.Number)
	}

	return fmt.Sprintf("MathOperation{Sign: %s, Number: %s}", sign, num)
}

type Monkey struct {
	Items     []int
	Operation MathOperation
	Test      MonkeyTest
}

func (m Monkey) String() string {
	return fmt.Sprintf("Monkey{Items: %v, Operation: %s, Test: %s}", m.Items, m.Operation.String(), m.Test.String())
}

type MonkeyTest struct {
	Divisor     int
	TrueMonkey  int
	FalseMonkey int
}

// returns the monkey the item is going to be thrown to
func (mt MonkeyTest) Execute(num int) int {
	if num%mt.Divisor == 0 {
		return mt.TrueMonkey
	} else {
		return mt.FalseMonkey
	}
}

func (mt MonkeyTest) String() string {
	return fmt.Sprintf("MonkeyTest{Divisor: %d, TrueMonkey: %d, FalseMonkey: %d}", mt.Divisor, mt.TrueMonkey, mt.FalseMonkey)
}

func getInput() []*Monkey {
	return input.Process("./days/11/input.txt", func(str string) []*Monkey {
		lines := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n\n")
		monkeys := []*Monkey{}

		for _, line := range lines {
			m := &Monkey{}
			rows := strings.Split(line, "\n")

			// Monkey ##:

			// Starting items: ...
			strItems := strings.Split(strings.ReplaceAll(rows[1], "  Starting items: ", ""), ", ")
			m.Items = array.Map(strItems, func(v string, _ int) int {
				n, err := s.Atoi(v)
				if err != nil {
					panic(err)
				}

				return n
			})

			// Operation: new = old ...
			strOp := strings.Split(strings.ReplaceAll(rows[2], "  Operation: new = old ", ""), " ")
			var opN *int
			if strOp[1] == "old" {
				opN = nil
			} else {
				tmp, err := s.Atoi(strOp[1])
				if err != nil {
					panic(err)
				}
				opN = &tmp
			}
			var opSign MathSign
			if strOp[0] == "*" {
				opSign = MULTIPLY
			} else {
				opSign = ADD
			}
			m.Operation = MathOperation{Sign: MathSign(opSign), Number: opN}

			// Test: divisible by ...
			strTestDiv := strings.ReplaceAll(rows[3], "  Test: divisible by ", "")
			testDiv, err := s.Atoi(strTestDiv)
			if err != nil {
				panic(err)
			}
			strTrueMonkey := strings.ReplaceAll(rows[4], "    If true: throw to monkey ", "")
			trueMonkey, err := s.Atoi(strTrueMonkey)
			if err != nil {
				panic(err)
			}
			strFalseMonkey := strings.ReplaceAll(rows[5], "    If false: throw to monkey ", "")
			falseMonkey, err := s.Atoi(strFalseMonkey)
			if err != nil {
				panic(err)
			}
			m.Test = MonkeyTest{Divisor: testDiv, TrueMonkey: trueMonkey, FalseMonkey: falseMonkey}

			// END
			monkeys = append(monkeys, m)
		}

		return monkeys
	})
}

func part1() int {
	monkeys := getInput()
	inspections := make([]int, len(monkeys))

	for i := 0; i < 20; i++ {
		for j, m := range monkeys {
			for _, item := range m.Items {
				inspections[j] += 1
				item = m.Operation.Execute(item)
				item /= 3
				idx := m.Test.Execute(item)
				monkeys[idx].Items = append(monkeys[idx].Items, item)
			}
			m.Items = []int{}
		}
	}

	sortedInspections := array.Reverse(math.Sort(inspections))
	return sortedInspections[0] * sortedInspections[1]
}

func part2() int {
	monkeys := getInput()
	inspections := make([]int, len(monkeys))
	prod := 1

	for _, m := range monkeys {
		prod *= m.Test.Divisor
	}

	for i := 1; i <= 10_000; i++ {
		for j, m := range monkeys {
			inspections[j] += len(m.Items)
			for _, item := range m.Items {
				item = m.Operation.Execute(item)
				item = item % prod
				idx := m.Test.Execute(item)
				monkeys[idx].Items = append(monkeys[idx].Items, item)
			}
			m.Items = []int{}
		}
	}

	sortedInspections := array.Reverse(math.Sort(inspections))
	return sortedInspections[0] * sortedInspections[1]
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
