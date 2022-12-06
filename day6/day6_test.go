package day6

import (
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestParseLine(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse file, got error: %s", err)
	}
	defer close()

	want := []int{
		5,
		6,
		10,
		11,
	}

	lineCount := 0
	for s.Scan() {
		got := ParseMarker(s.Text(), 4)
		want := want[lineCount]

		if got != want {
			t.Fatalf("got: %d, want: %d", got, want)
		}
		lineCount++
	}

}
