package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day10"
)

func main() {
	output, err := day10.Day10PartTwo("./day10/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
