package main

import (
    "fmt"

    "github.com/Tch1b0/polaris/input"
)


func getInput() string {
    return input.Process("./days/14/input.txt", func(str string) string {
        return str
    })
}

func part1() int {
    return -1
}

func part2() int {
    return -1
}

func main() {
    fmt.Printf("PART 1: %d\n", part1())

    
    fmt.Printf("PART 2: %d\n", part2())
}

