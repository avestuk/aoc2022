package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day4(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	totalFullContains := 0
	for s.Scan() {
		if len(s.Text()) == 0 {
			// Does this happen?
			panic("got empty line")
		}

		r1, r2 := findRanges(s.Text())
		totalFullContains += fullyContains(r1, r2)
	}

	return totalFullContains, nil
}

func Day4PartTwo(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	totalFullContains := 0
	for s.Scan() {
		if len(s.Text()) == 0 {
			// Does this happen?
			panic("got empty line")
		}

		r1, r2 := findRanges(s.Text())
		totalFullContains += overlaps(r1, r2)
	}

	return totalFullContains, nil

}

type assignment struct {
	lower, upper int
}

func findRanges(s string) (assignment, assignment) {
	a1, a2 := strings.Split(s, ",")[0], strings.Split(s, ",")[1]
	var r1, r2 assignment
	i, err := strconv.Atoi(strings.Split(a1, "-")[0])
	if err != nil {
		panic(fmt.Sprintf("could not convert: %s to int", (strings.Split(a1, "-")[0])))
	}
	r1.lower = i

	i, err = strconv.Atoi(strings.Split(a1, "-")[1])
	if err != nil {
		panic(fmt.Sprintf("could not convert: %s to int", (strings.Split(a1, "-")[0])))
	}
	r1.upper = i

	i, err = strconv.Atoi(strings.Split(a2, "-")[0])
	if err != nil {
		panic(fmt.Sprintf("could not convert: %s to int", (strings.Split(a2, "-")[0])))
	}
	r2.lower = i

	i, err = strconv.Atoi(strings.Split(a2, "-")[1])
	if err != nil {
		panic(fmt.Sprintf("could not convert: %s to int", (strings.Split(a2, "-")[0])))
	}
	r2.upper = i

	return r1, r2
}

func fullyContains(r1, r2 assignment) int {
	if r1.lower <= r2.lower {
		if r1.upper >= r2.upper {
			return 1
		}
	}

	if r2.lower <= r1.lower {
		if r2.upper >= r1.upper {
			return 1
		}
	}
	return 0
}

func overlaps(r1, r2 assignment) int {
	// Lower bound is less than r2 lower
	// Upper bound is greater than r2 lower
	// 1-3
	// 2-5
	// 100-300
	// 200-800
	if r1.lower <= r2.lower && r1.upper >= r2.lower {
		return 1
	}

	// Lower bound is bigger than r2 lower
	// Lower bound is less than r2 upper
	// 4-7
	// 3-4
	if r1.lower >= r2.lower && r1.lower <= r2.upper {
		return 1
	}
	return 0
}
