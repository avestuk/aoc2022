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

	i := 0
	H := point{0, 0}
	T := point{0, 0}
	visitedPoints := []point{}
	for s.Scan() {
		i++
		H, T, visitedPoints = parseMove(s.Text(), H, T, visitedPoints)
		if !pointsTouching(H, T) {
			panic(fmt.Sprintf("H: %v, T: %v not touching, s: %s, i: %d", H, T, s.Text(), i))
		}
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

	input := strings.Split(s, " ")
	direction, lengthStr := input[0], input[1]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		panic(fmt.Sprintf("could not convert: %s to int, got err: %s", lengthStr, err))
	}

	switch direction {
	case "R":
		return move(H, T, length, visitedPoints, moveRightf)
	case "L":
		return move(H, T, length, visitedPoints, moveLeftf)
	case "D":
		return move(H, T, length, visitedPoints, moveDownf)
	case "U":
		return move(H, T, length, visitedPoints, moveUpf)
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

var (
	moveLeftf = func(H point, i int) point {
		return point{H.x - i, H.y}
	}
	moveRightf = func(H point, i int) point {
		return point{H.x + i, H.y}
	}
	moveUpf = func(H point, i int) point {
		return point{H.x, H.y + i}
	}
	moveDownf = func(H point, i int) point {
		return point{H.x, H.y - i}
	}
)

func move(H1, T point, v int, visitedPoints []point, moveH func(H point, i int) point) (point, point, []point) {
	var h2, t2 point
	for i := 1; i <= v; i++ {
		h2 = moveH(H1, i)

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

	return h2, T, visitedPoints

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
