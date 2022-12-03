package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day3"
)

func main() {
	output, err := day3.Day3Part2("./day3/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
