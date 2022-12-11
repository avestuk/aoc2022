package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day9"
)

func main() {
	output, err := day9.Day9PartTwo("./day9/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
