package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day8"
)

func main() {
	output, err := day8.Day8("./day8/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%d", output)
}
