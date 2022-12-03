package day3

import (
	"fmt"

	"github.com/avestuk/aoc2022/pkg/parse"
	mapset "github.com/deckarep/golang-set/v2"
)

func Day3(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	totalPriority := 0
	for s.Scan() {
		if len(s.Text()) == 0 {
			// Does this happen?
			panic("got empty line")
		}

		item := findItem(s.Text())
		totalPriority += getPriority(item)
	}

	return totalPriority, nil
}

func Day3Part2(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	totalScore := 0
	rucksacks := []string{}
	for s.Scan() {
		if len(s.Text()) == 0 {
			panic("got empty line")
		}

		rucksacks = append(rucksacks, s.Text())
		if len(rucksacks) != 3 {
			continue
		}

		common := findIntersection(rucksacks)
		totalScore += getPriority(common)
		rucksacks = []string{}
	}

	return totalScore, nil
}

func findItem(rucksack string) rune {
	// Split the line
	left, right := splitLine(rucksack)

	// Loop through left and right
	for _, leftItem := range left {
		for _, rightItem := range right {
			if leftItem == rightItem {
				return leftItem
			}
		}
	}

	panic(fmt.Sprintf("couldn't find match for line: %s", rucksack))
}

func getPriority(r rune) int {
	value := int(r)

	switch {
	// is lowercase
	case value >= 97:
		return value - 96
	// is uppercase
	case value <= 96:
		return value - 38
	default:
		panic(fmt.Sprintf("got rune: %s", string(r)))
	}
}

func splitLine(s string) (string, string) {
	return s[:(len(s) / 2)], s[len(s)/2:]
}

func findIntersection(rucksacks []string) rune {
	common := map[rune]bool{}
	for _, item := range rucksacks[0] {
		common[item] = true
	}

	intersection := map[rune]bool{}
	for _, item := range rucksacks[1] {
		if common[item] {
			intersection[item] = true
		}
	}

	for _, item := range rucksacks[2] {
		if intersection[item] {
			// Item is in all three lists
			return item
		}
	}

	panic(fmt.Sprintf("no matches found for: %v", rucksacks))
}

func findIntersectionSet(rucksacks []string) rune {
	rucksack0 := mapset.NewSet[rune]()
	rucksack1 := mapset.NewSet[rune]()
	rucksack2 := mapset.NewSet[rune]()

	for _, item := range rucksacks[0] {
		rucksack0.Add(item)
	}

	for _, item := range rucksacks[1] {
		rucksack1.Add(item)
	}

	for _, item := range rucksacks[2] {
		rucksack2.Add(item)
	}

	common, _ := rucksack0.Intersect(rucksack1).Intersect(rucksack2).Pop()
	return common
}
