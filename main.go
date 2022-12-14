package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day12"
)

func main() {
	output, err := day12.Day12PartTwo("./day12/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
