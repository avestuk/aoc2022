package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day7"
)

func main() {
	output, err := day7.Day7PartTwo("./day7/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
