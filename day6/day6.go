package day6

import (
	"fmt"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day6(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	for s.Scan() {
		return ParseMarker(s.Text(), 4), nil

	}

	return 0, nil
}

func Day6PartTwo(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	for s.Scan() {
		return ParseMarker(s.Text(), 14), nil

	}

	return 0, nil
}

func ParseMarker(s string, distinctCount int) int {
	var parsedChars string
nextChar:
	for i, r := range s {
		parsedChars += string(r)
		if i >= distinctCount-1 {
			for _, r := range parsedChars {
				count := strings.Count(parsedChars, string(r))
				if count > 1 {
					parsedChars = parsedChars[1:]
					continue nextChar
				}
			}
			return i + 1
		}
	}

	panic("didn't find a start of packet marker")
}
