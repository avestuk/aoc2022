package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day6"
)

func main() {
	output, err := day6.Day6PartTwo("./day6/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
