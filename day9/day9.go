package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day9(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	H := point{0, 0}
	T := point{0, 0}
	visitedPoints := []point{}
	for s.Scan() {
		H, T, visitedPoints = parseMove(s.Text(), H, T, visitedPoints)
	}

	return len(visitedPoints), nil
}

type point struct {
	x, y int
}

func parseMove(s string, H, T point, visitedPoints []point) (point, point, []point) {
	if len(visitedPoints) == 0 {
		visitedPoints = append(visitedPoints, point{0, 0})
	}

	move := strings.Split(s, " ")
	direction, lengthStr := move[0], move[1]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		panic(fmt.Sprintf("could not convert: %s to int, got err: %s", lengthStr, err))
	}

	switch direction {
	case "R":
		return moveRight(H, T, length, visitedPoints)
	case "L":
		return moveLeft(H, T, length, visitedPoints)
	case "D":
		return moveDown(H, T, length, visitedPoints)
	case "U":
		return moveUp(H, T, length, visitedPoints)
	default:
		panic(fmt.Sprintf("could not match direction: %s", direction))
	}

}

func pointsTouching(p1, p2 point) bool {
	if math.Abs(float64(p1.x-p2.x)) <= 1 && math.Abs(float64(p1.y-p2.y)) <= 1 {
		return true

	}
	return false
}

func moveRight(H1, T point, v int, visitedPoints []point) (point, point, []point) {
	var h2, t2 point
	for i := 1; i <= v; i++ {
		h2 = point{H1.x + i, H1.y}

		if !pointsTouching(h2, T) {
			t2 = moveT(h2, T)
			var t2Visited bool
			for _, p := range visitedPoints {
				if p == t2 {
					t2Visited = true
					break
				}
			}

			if !t2Visited {
				visitedPoints = append(visitedPoints, t2)
			}
			T = t2
		}
	}

	return h2, t2, visitedPoints

}

func moveLeft(H1, T point, v int, visitedPoints []point) (point, point, []point) {
	var h2, t2 point
	for i := 1; i <= v; i++ {
		h2 = point{H1.x - i, H1.y}

		if !pointsTouching(h2, T) {
			t2 = moveT(h2, T)
			var t2Visited bool
			for _, p := range visitedPoints {
				if p == t2 {
					t2Visited = true
					break
				}
			}

			if !t2Visited {
				visitedPoints = append(visitedPoints, t2)
			}
			T = t2
		}
	}

	return h2, t2, visitedPoints

}

func moveUp(H1, T point, v int, visitedPoints []point) (point, point, []point) {
	var h2, t2 point
	for i := 1; i <= v; i++ {
		h2 = point{H1.x, H1.y + i}

		if !pointsTouching(h2, T) {
			t2 = moveT(h2, T)
			var t2Visited bool
			for _, p := range visitedPoints {
				if p == t2 {
					t2Visited = true
					break
				}
			}

			if !t2Visited {
				visitedPoints = append(visitedPoints, t2)
			}
			T = t2
		}
	}

	return h2, t2, visitedPoints

}
func moveDown(H1, T point, v int, visitedPoints []point) (point, point, []point) {
	var h2, t2 point
	for i := 1; i <= v; i++ {
		h2 = point{H1.x, H1.y - i}

		if !pointsTouching(h2, T) {
			t2 = moveT(h2, T)
			var t2Visited bool
			for _, p := range visitedPoints {
				if p == t2 {
					t2Visited = true
					break
				}
			}

			if !t2Visited {
				visitedPoints = append(visitedPoints, t2)
			}
			T = t2
		}
	}

	return h2, t2, visitedPoints

}
func moveT(H, T point) point {
	dX := H.x - T.x
	dY := H.y - T.y
	dT := point{dX, dY}

	// Move is legal
	if math.Abs(float64(dT.x)) <= 1 && math.Abs(float64(dT.y)) <= 1 {
		return T
	}

	if dX > 1 {
		dX--
	} else if dX < -1 {
		dX++
	}

	if dY > 1 {
		dY--
	} else if dY < -1 {
		dY++
	}

	t := point{dX + T.x, dY + T.y}
	if pointsTouching(H, t) {
		return t
	}
	panic(fmt.Sprintf("failed on H: %v, T: %v, t: %v", H, T, t))
}
