package day2

import (
	"strings"
	"testing"
)

func TestRoundOutcome(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	wantOutcome := []int{
		6,
		0,
		3,
	}

	wantScore := []int{
		8,
		1,
		6,
	}

	for i, round := range input {
		opponent := strings.Split(round, " ")[0]
		me := convertMove(strings.Split(round, " ")[1])
		gotOutcome := roundOutcome(opponent, me)

		if gotOutcome != wantOutcome[i] {
			t.Fatalf("got: %d, want: %d", gotOutcome, wantOutcome[i])
		}

		gotScore := parseRound(round)
		if gotScore != wantScore[i] {
			t.Fatalf("got: %d, want: %d", gotScore, wantScore[i])
		}
	}
}

func TestMoveScore(t *testing.T) {
	input := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	for move, want := range input {
		got := moveScore(move)
		if got != want {
			t.Fatalf("got: %d, want: %d for move: %s", got, want, move)
		}
	}
}

func TestRequiredOutcome(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	want := []string{
		rock,
		rock,
		rock,
	}

	for i, round := range input {
		opponent := strings.Split(round, " ")[0]
		outcome := strings.Split(round, " ")[1]
		got := requiredOutcome(opponent, outcome)

		if got != want[i] {
			t.Fatalf("got: %s, want: %s", got, want)
		}
	}
}
