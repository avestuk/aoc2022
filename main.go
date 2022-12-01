package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day1"
)

func main() {
	mostCalories, err := day1.Day1Part2("./day1/day1_input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("mostCalories: %d", mostCalories)
}
