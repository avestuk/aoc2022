package day2

import (
	"fmt"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

const (
	rock    = "A"
	paper   = "B"
	scissor = "C"

	win  = 6
	loss = 0
	draw = 3
)

// The winner is the player with the highest score
// Your total score is the sum of your scores for each round
// The score for a single round is the score for the shape you selected
// 1 for Rock, 2 for Paper, and 3 for Scissors
// Plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func Day2(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	totalScore := 0
	for s.Scan() {
		if len(s.Text()) == 0 {
			// Does this happen?
			panic("got empty line")
		}

		totalScore += parseRound(s.Text())
	}

	return totalScore, nil
}

func Day2Part2(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	totalScore := 0
	for s.Scan() {
		if len(s.Text()) == 0 {
			panic("got empty line")
		}

		totalScore += parseRoundTwo(s.Text())

	}

	return totalScore, nil
}

func parseRound(round string) int {
	opponent := strings.Split(round, " ")[0]
	me := strings.Split(round, " ")[1]
	normalizedMe := convertMove(me)

	score := 0
	score += roundOutcome(opponent, normalizedMe)
	score += moveScore(normalizedMe)

	return score
}

func moveScore(s string) int {
	switch s {
	case rock:
		return 1
	case paper:
		return 2
	case scissor:
		return 3
	default:
		return 0
	}
}

func roundOutcome(opponent, me string) int {
	switch {
	case me == opponent:
		return draw
	case me == winning(opponent):
		return win
	default:
		return loss
	}
}

func convertMove(s string) string {
	switch s {
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissor
	default:
		panic(fmt.Sprintf("couldn't convert move: %s", s))
	}
}

func winning(move string) string {
	switch move {
	case rock:
		return paper
	case scissor:
		return rock
	case paper:
		return scissor
	default:
		return ""
	}
}

func parseRoundTwo(round string) int {
	opponent := strings.Split(round, " ")[0]
	me := requiredOutcome(opponent, strings.Split(round, " ")[1])

	score := 0
	score += roundOutcome(opponent, me)
	score += moveScore(me)

	return score
}

func requiredOutcome(opponent, outcome string) string {
	switch outcome {
	case "X":
		return losing(opponent)
	case "Y":
		return opponent
	case "Z":
		return winning(opponent)
	default:
		panic("couldn't convert outcome")
	}
}

func losing(move string) string {
	switch move {
	case rock:
		return scissor
	case scissor:
		return paper
	case paper:
		return rock
	default:
		return ""
	}
}
