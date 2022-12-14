package day12

import (
	"reflect"
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestFindMoves(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse file, got error: %s", err)
	}
	defer close()

	// Parse the input and find the starting point and endpoint
	grid, start, end := parseGrid(s)
	if !reflect.DeepEqual(start, point{0, 0}) {
		t.Fatalf("got: %v, want: %v", start, point{0, 0})
	}
	if !reflect.DeepEqual(end, point{5, 2}) {
		t.Fatalf("got: %v, want: %v", start, point{0, 0})
	}

	got := findMoves(grid, start)

	want := []point{{0, 1}, {1, 0}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}

	got = findMoves(grid, end)

	want = []point{{5, 1}, {5, 3}, {4, 2}, {6, 2}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}

}
