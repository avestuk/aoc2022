package day8

import (
	"reflect"
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestParseLine(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse input, got err: %s", err)
	}
	defer close()

	trees := parseLines(s)

	want := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	for i, got := range trees {
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf("got: %v, want: %v", got, want[i])
		}

	}
}

func TestCountVisible(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse input, got err: %s", err)
	}
	defer close()

	trees := parseLines(s)

	got := countVisible(trees)
	want := 21

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestMostScenic(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse input, got err: %s", err)
	}
	defer close()

	trees := parseLines(s)

	got := mostScenic(trees)
	want := 8

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}
