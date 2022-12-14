package day12

import (
	"bufio"
	"fmt"
	"math"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day12(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	// Parse the input and find the starting point and endpoint
	grid, start, end := parseGrid(s)

	moveCount := bfs(grid, start, end)
	//moveCount := bfsv2(grid, start, end)

	return moveCount, nil
}

func Day12PartTwo(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	// Parse the input and find the starting point and endpoint
	grid, start, end, lowest := parseGridPartTwo(s)

	moveCount := bfs(grid, start, end)

	for _, p := range lowest {
		if moveCountA := bfs(grid, p, end); moveCountA < moveCount {
			moveCount = moveCountA
		}
	}

	return moveCount, nil
}

func parseGrid(s *bufio.Scanner) ([][]rune, point, point) {
	grid := [][]rune{}
	line := 0
	var start, end point
	for s.Scan() {
		grid = append(grid, []rune{})
		for i, r := range s.Text() {
			grid[line] = append(grid[line], r)

			switch r {
			// S
			case 83:
				start = point{i, line}
			// E
			case 69:
				end = point{i, line}
			}
		}
		line++
	}

	return grid, start, end
}

func parseGridPartTwo(s *bufio.Scanner) ([][]rune, point, point, []point) {
	grid := [][]rune{}
	line := 0
	var start, end point
	lowestPoints := []point{}
	for s.Scan() {
		grid = append(grid, []rune{})
		for i, r := range s.Text() {
			grid[line] = append(grid[line], r)

			switch r {
			// S
			case 83:
				start = point{i, line}
			// E
			case 69:
				end = point{i, line}
			case 'a':
				lowestPoints = append(lowestPoints, point{i, line})
			}
		}
		line++
	}

	return grid, start, end, lowestPoints
}

func bfs(grid [][]rune, start, end point) int {
	// Create a map to keep track of cords we have seen
	seen := map[point]bool{}
	seen[start] = true
	queue := []point{start}
	moveCount := 0

	for len(queue) > 0 {
		if moveCount == 24 {
			fmt.Print()
		}
		k := len(queue)
		for i := 0; i < k; i++ {
			current := queue[0]
			queue = queue[1:]

			// Check if we are at the desired point
			if current == end {
				return moveCount
			}

			// Find all the allowed moves for the current point. If we have seen
			// them before then skip them
			newMoves := findMoves(grid, current)
			for _, move := range newMoves {
				// If we haven't visited the point then
				if ok := seen[move]; !ok {
					queue = append(queue, move)
					seen[move] = true
				}
			}
		}
		moveCount++
	}

	return math.MaxInt
}

func (p *point) getPoint(grid [][]rune) rune {
	return grid[p.y][p.x]
}

// findMoves for p
func findMoves(grid [][]rune, p point) []point {
	potentialMoves := []point{}
	// Up
	if p.y-1 >= 0 {
		potentialMoves = append(potentialMoves, point{p.x, p.y - 1})
	}
	// Down
	if p.y+1 < len(grid) {
		potentialMoves = append(potentialMoves, point{p.x, p.y + 1})
	}
	// Left
	if p.x-1 >= 0 {
		potentialMoves = append(potentialMoves, point{p.x - 1, p.y})
	}
	// Right
	if p.x+1 < len(grid[p.y]) {
		potentialMoves = append(potentialMoves, point{p.x + 1, p.y})
	}

	moves := []point{}
	for _, m := range potentialMoves {
		if moveAllowed(grid[p.y][p.x], grid[m.y][m.x]) {
			moves = append(moves, m)
		}
	}

	return moves
}

func moveAllowed(source, destination rune) bool {
	if source == 'S' && destination == 'a' {
		return true
	}
	if destination == 'E' && source == 'z' {
		return true
	}
	// Account for comparison with 'E'
	// destination: b
	// source: a
	// delta = 1
	delta := destination - source
	max := 'a' - 'z'
	return delta <= 1 && delta >= max
}

type point struct {
	x, y int
}
