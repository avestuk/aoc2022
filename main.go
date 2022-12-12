package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day11"
)

func main() {
	output, err := day11.Day11PartTwo("./day11/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
