package day7

import (
	"fmt"
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestParseLine(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse input, got err: %s", err)
	}
	defer close()

	root := parseLines(s)

	for _, subDir := range root.subDirs {
		fmt.Printf("subDir: %s, %#v\n", subDir.name, subDir)
	}

	total := root.filterSizes(100000)

	if total != 95437 {
		t.Fatalf("got: %d, want: 95437", total)
	}
}
