package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day10(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	cycle := 1
	x := 1
	strength := 0
	for s.Scan() {

		opCompletion, operation := parseInput(s.Text())
		for i := 1; i <= opCompletion; i++ {
			switch cycle {
			case 20, 60, 100, 140, 180, 220:
				strength += x * cycle
			}

			cycle++
		}
		x = x + operation
	}
	return strength, nil
}

func parseInput(s string) (int, int) {
	switch {
	case strings.HasPrefix(s, "noop"):
		return 1, 0
	case strings.HasPrefix(s, "addx"):
		inc := strings.Split(s, " ")[1]
		i, err := strconv.Atoi(inc)
		if err != nil {
			panic(fmt.Sprintf("failed to convert: %s to int", inc))
		}
		return 2, i
	default:
		panic(fmt.Sprintf("could not parse: %s", s))
	}
}
