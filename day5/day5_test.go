package day5

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestGetStacks(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to parse file, got error: %s", err))
	}
	defer close()

	// Original ordering
	// want := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	want := [][]string{
		{"N", "Z"},
		{"D", "C", "M"},
		{"P"},
	}

	got := getStacks(s)

	for i, stack := range got {
		if len(stack) != 0 {
			if !reflect.DeepEqual(stack, want[i]) {
				t.Fatalf("got: %v, want: %v", stack, want[i])
			}
		}
	}
}

func TestMoveStacks(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to parse file, got error: %s", err))
	}
	defer close()

	stacks := getStacks(s)

	got := moveStacks(s, stacks)
	want := "CMZ"

	if got != want {
		t.Fatalf("got: %s, want: %s", got, want)
	}

}

func TestMoveStacksMultiple(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to parse file, got error: %s", err))
	}
	defer close()

	stacks := getStacks(s)

	got := moveStacksMultiple(s, stacks)
	want := "MCD"

	if got != want {
		t.Fatalf("got: %s, want: %s", got, want)
	}

}

func TestParseInstructions(t *testing.T) {
	instructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	want := [][]int{
		{1, 1, 0},
		{3, 0, 2},
		{2, 1, 0},
		{1, 0, 1},
	}

	for i, s := range instructions {
		gotNumber, gotSource, gotDestination := parseInstructions(s)
		got := []int{gotNumber, gotSource, gotDestination}

		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf("got: %v, want: %v", got, want[i])
		}

	}
}
