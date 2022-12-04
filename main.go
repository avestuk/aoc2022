package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day4"
)

func main() {
	output, err := day4.Day4PartTwo("./day4/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
