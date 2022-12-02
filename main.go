package main

import (
	"fmt"

	"github.com/avestuk/aoc2022/day2"
)

func main() {
	mostCalories, err := day2.Day2Part2("./day2/day2_input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("mostCalories: %d", mostCalories)
}
