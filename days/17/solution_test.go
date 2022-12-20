package main_test

import (
    "testing"
    "strings"

    "github.com/Tch1b0/polaris/input"
    "github.com/stretchr/testify/require"
)

var (
    in = input.Read("./days/17/example_input.txt")
    output = input.Process("./days/17/example_output.txt", func(v string) []string {
        return strings.Split(v, "\n")
    })
)

func test_part1(t *testing.T) {
}

func test_part2(t *testing.T) {
}

