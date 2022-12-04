package day4

import (
	"testing"
)

func TestFindRanges(t *testing.T) {
	inputRanges := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		//"2-8,3-7",
		//"6-6,4-6",
		//"2-6,4-8",
	}
	wantRanges := [][]assignment{
		[]assignment{
			assignment{
				lower: 2,
				upper: 4,
			},
			assignment{
				lower: 6,
				upper: 8,
			},
		},
		[]assignment{
			assignment{
				lower: 2,
				upper: 3,
			},
			assignment{
				lower: 4,
				upper: 5,
			},
		},
		[]assignment{
			assignment{
				lower: 5,
				upper: 7,
			},
			assignment{
				lower: 7,
				upper: 9,
			},
		},
	}

	for i, inputRange := range inputRanges {
		g1, g2 := findRanges(inputRange)
		w1, w2 := wantRanges[i][0], wantRanges[i][1]

		if g1.lower != w1.lower && g1.upper != w1.upper {
			t.Fatalf("got: %v, want: %v", g1, w1)
		}
		if g2.lower != w2.lower && g2.upper != w2.upper {
			t.Fatalf("got: %v, want: %v", g2, w2)
		}
	}
}

func TestFullyContains(t *testing.T) {
	got := fullyContains(
		assignment{
			lower: 1,
			upper: 4,
		},
		assignment{
			lower: 2,
			upper: 4,
		},
	)
	want := 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = fullyContains(
		assignment{
			lower: 30,
			upper: 40,
		},
		assignment{
			lower: 20,
			upper: 50,
		},
	)
	want = 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = fullyContains(
		assignment{
			lower: 30,
			upper: 40,
		},
		assignment{
			lower: 50,
			upper: 55,
		},
	)
	want = 0

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestOverlaps(t *testing.T) {
	got := overlaps(
		assignment{
			lower: 5,
			upper: 7,
		},
		assignment{
			lower: 7,
			upper: 9,
		},
	)
	want := 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = overlaps(
		assignment{
			lower: 2,
			upper: 8,
		},
		assignment{
			lower: 3,
			upper: 7,
		},
	)
	want = 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = overlaps(
		assignment{
			lower: 6,
			upper: 6,
		},
		assignment{
			lower: 4,
			upper: 6,
		},
	)
	want = 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
	got = overlaps(
		assignment{
			lower: 6,
			upper: 6,
		},
		assignment{
			lower: 4,
			upper: 8,
		},
	)
	want = 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
	got = overlaps(
		assignment{
			lower: 6,
			upper: 90,
		},
		assignment{
			lower: 4,
			upper: 8,
		},
	)
	want = 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}
