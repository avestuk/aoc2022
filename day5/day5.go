package day5

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day5(file string) (string, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return "", fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	stacks := getStacks(s)
	topContainers := moveStacks(s, stacks)

	return topContainers, nil
}

func Day5PartTwo(file string) (int, error) {
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
	}

	return totalFullContains, nil

}

func getStacks(s *bufio.Scanner) [][]string {
	stacks := make([][]string, 9)
	// positions := map[int]int{
	// 	2: 0,
	// 	6: 1,
	// 	10: 2,
	// 	14: 3,
	// 	18: 4,
	// 	22: 5,
	// 	26: 6,
	// 	30: 7,
	// 	34: 8,
	// 	38
	// }
	for s.Scan() {
		if len(s.Text()) == 0 {
			// This happens on the blank line between instructions and the layout
			return stacks
		}
		fmt.Printf("text: %s", s.Text())

		re, err := regexp.Compile("([A-Z])")
		if err != nil {
			panic(fmt.Sprintf("could not parse regex, got err: %s", err))
		}

		indexes := re.FindAllStringIndex(s.Text(), -1)
		positions := []int{}
		for _, index := range indexes {
			rawPosition := index[1]
			//if rawPosition == 2 {
			//	positions = append(positions, 0)
			//	continue
			//}

			i := ((rawPosition - 2) / 2) / 2
			positions = append(positions, i)
			if (rawPosition-2)%2 != 0 {
				panic(fmt.Sprintf("got weird index: %d", rawPosition))
			}
		}
		fmt.Printf("positions: %v\n", positions)
		containers := re.FindAllString(s.Text(), -1)
		for i, char := range containers {
			stacks[positions[i]] = append(stacks[positions[i]], string(char))
		}
	}
	return stacks
}

func moveStacks(s *bufio.Scanner, stacks [][]string) string {
	for s.Scan() {
		if len(s.Text()) == 0 {
			panic("text length was 0")
		}

		numberToMove, sourceIndex, destinationIndex := parseInstructions(s.Text())
		source := stacks[sourceIndex]
		destination := stacks[destinationIndex]

		var container string
		for i := 1; i <= numberToMove; i++ {
			container, source = source[0], source[1:]
			destination = append([]string{container}, destination...)

			stacks[sourceIndex] = source
			stacks[destinationIndex] = destination
		}
	}

	var topContainers string
	for _, stack := range stacks {
		topContainers += stack[0]

	}

	return topContainers
}

func moveStacksMultiple(s *bufio.Scanner, stacks [][]string) string {}

func parseInstructions(s string) (numberToMove, source, destination int) {
	matches, err := fmt.Sscanf(s, "move %d from %d to %d", &numberToMove, &source, &destination)
	if err != nil || matches != 3 {
		panic(fmt.Sprintf("got err: %s, got matches: %d, for: %s", err, matches, s))
	}

	return numberToMove, source - 1, destination - 1
}
