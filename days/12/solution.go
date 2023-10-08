package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
)

var directions = []math.Vector2[int]{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

type Field struct {
	Importance int
	Level      int
	Position   math.Vector2[int]
}

type Map struct {
	Fields []*Field
	Start  *Field
	End    *Field
	Width  int
	Height int
}

func (m Map) FieldAt(pos math.Vector2[int]) **Field {
	return array.GetFirst(m.Fields, func(f *Field, i int) bool { return f.Position == pos })
}

func (m Map) LevelFields(level int) []*Field {
	return array.Filter(m.Fields, func(f *Field, i int) bool { return f.Level == level })
}

func (m *Map) mapImportance(origin *Field, f *Field) {
	for _, dir := range directions {
		pos := f.Position.Add(dir)
		if pos == origin.Position {
			continue
		}

		newFP := m.FieldAt(pos)
		if newFP == nil || (*newFP).Level < f.Level-1 {
			continue
		}
		newF := *newFP
		// newF is valid

		importance := f.Importance + 1
		if newF.Importance > importance || newF.Importance == -1 {
			newF.Importance = importance
			m.mapImportance(f, newF)
		}
	}
}

func (m *Map) MapImportance() {
	m.mapImportance(m.End, m.End)
	fmt.Println(m)
}

var selected *Field = nil

func (m Map) String() string {
	str := ""
	str += "\033[1;1H\033[2J"

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			cur := *m.FieldAt(math.Vector2[int]{X: x, Y: y})
			if cur == selected {
				str += fmt.Sprintf("\033[91m%03d\033[0m ", cur.Importance)
			} else {
				str += fmt.Sprintf("\033[38;2;%d;%d;%dm%03d\033[0m ", (cur.Level%2)*150+50, (cur.Level%3)*50+50, (cur.Level%4)*30+50, cur.Importance)
			}
			// str += fmt.Sprintf("%03d ", cur.Level)

		}
		str += "\n"
	}

	return str
}

func getInput() Map {
	return input.Process("./days/12/input.txt", func(str string) Map {
		lines := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n")

		m := Map{
			Fields: []*Field{},
			Height: len(lines),
			Width:  len(lines[0]),
		}

		for y, line := range lines {
			for x, char := range line {
				// map characters from a - z to 0 - 25
				level := int(char - 'a')

				f := &Field{
					Level: level,
					Position: math.Vector2[int]{
						X: x,
						Y: y,
					},
					Importance: -1,
				}

				// special case S
				if char == 'S' {
					f.Level = 0
					m.Start = f
				} else /* special case E */ if char == 'E' {
					f.Level = 25
					f.Importance = 0
					m.End = f
				}

				m.Fields = append(m.Fields, f)
			}
		}

		m.MapImportance()
		return m
	})
}

func part1(m Map) int {
	return m.Start.Importance
}

func part2(m Map) int {
	aFields := m.LevelFields(0)
	k := []*Field{}
	for _, f := range aFields {
		if f != nil && f.Importance != -1 {
			k = append(k, f)
		}
	}
	importances := array.Map(k, func(f *Field, i int) int { return f.Importance })

	return math.Min(importances)
}

func main() {
	m := getInput()
	fmt.Printf("PART 1: %d\n", part1(m))

	fmt.Printf("PART 2: %d\n", part2(m))
}
