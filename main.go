package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day5"
)

func main() {
	output, err := day5.Day5PartTwo("./day5/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%s", output)
}
