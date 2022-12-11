package day9

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestParseMove(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to parse file, got error: %s", err))
	}
	defer close()
	//..##..
	//...##.
	//.####.
	//....#.
	//s###..

	want := []point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{4, 2},
		{4, 3},
		{3, 4},
		{2, 4},
		{3, 3},
		{3, 2},
		{2, 2},
		{1, 2},
	}

	i := 0
	H := point{0, 0}
	T := point{0, 0}
	visitedPoints := []point{}

	for s.Scan() {
		H, T, visitedPoints = parseMove(s.Text(), H, T, visitedPoints)

		wantPoint := want[i]
		if visitedPoints[i] != wantPoint {
			t.Fatalf("got: %v from: %s (%d), want: %v", visitedPoints, s.Text(), i, want[i])
		}
		i++
	}

	fmt.Printf("\n visitedPoints: %v", visitedPoints)
}

func TestPointsTouching(t *testing.T) {

	got := pointsTouching(point{0, 0}, point{1, 1})
	want := true

	if got != want {
		t.Fatalf("case1: got: %t != want: %t", got, want)
	}

	got = pointsTouching(point{0, 0}, point{-1, 1})
	want = true

	if got != want {
		t.Fatalf("case2: got: %t != want: %t", got, want)
	}

	got = pointsTouching(point{0, 0}, point{1, -1})
	want = true

	if got != want {
		t.Fatalf("case3: got: %t != want: %t", got, want)
	}

	got = pointsTouching(point{0, 0}, point{-1, -1})
	want = true

	if got != want {
		t.Fatalf("case4: got: %t != want: %t", got, want)
	}

	got = pointsTouching(point{0, 0}, point{-2, 0})
	want = false

	if got != want {
		t.Fatalf("case5: got: %t != want: %t", got, want)
	}
}

func TestMoveT(t *testing.T) {
	got := moveT(point{0, 0}, point{1, 0})
	want := point{1, 0}

	if got != want {
		t.Fatalf("case1: got: %v != want: %v", got, want)
	}

	got = moveT(point{3, 1}, point{1, 1})
	want = point{2, 1}

	if got != want {
		t.Fatalf("case2: got: %v != want: %v", got, want)
	}

	got = moveT(point{0, -2}, point{0, 0})
	want = point{0, -1}

	if got != want {
		t.Fatalf("case3: got: %v != want: %v", got, want)
	}

	got = moveT(point{-1, -2}, point{0, 0})
	want = point{-1, -1}

	if got != want {
		t.Fatalf("case4: got: %v != want: %v", got, want)
	}

	got = moveT(point{1, -2}, point{0, 0})
	want = point{1, -1}

	if got != want {
		t.Fatalf("case5: got: %v != want: %v", got, want)
	}

	got = moveT(point{1, 2}, point{0, 0})
	want = point{1, 1}

	if got != want {
		t.Fatalf("case6: got: %v != want: %v", got, want)
	}

	got = moveT(point{3, 2}, point{2, 0})
	want = point{3, 1}

	if got != want {
		t.Fatalf("case7: got: %v != want: %v", got, want)
	}

}

func TestMoveRight(t *testing.T) {
	h1 := point{0, 0}
	t1 := point{0, 0}
	distace := 4
	wantH2 := point{4, 0}
	wantT2 := point{3, 0}
	wantVisitedPoints := []point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}

	gotH2, gotT2, gotVisitedPoints := move(h1, t1, distace, []point{{0, 0}}, moveRightf)

	if gotH2 != wantH2 {
		t.Fatalf("got: %v, want:%v", gotH2, wantH2)
	}
	if gotT2 != wantT2 {
		t.Fatalf("got: %v, want:%v", gotT2, wantT2)

	}
	if !reflect.DeepEqual(gotVisitedPoints, wantVisitedPoints) {
		t.Fatalf("got: %v, want:%v", gotVisitedPoints, wantVisitedPoints)

	}

}

func TestMoveLeft(t *testing.T) {
	h1 := point{0, 0}
	t1 := point{0, 0}
	distance := 4
	wantH2 := point{-4, 0}
	wantT2 := point{-3, 0}
	wantVisitedPoints := []point{
		{0, 0},
		{-1, 0},
		{-2, 0},
		{-3, 0},
	}

	gotH2, gotT2, gotVisitedPoints := move(h1, t1, distance, []point{{0, 0}}, moveLeftf)

	if gotH2 != wantH2 {
		t.Fatalf("got: %v, want:%v", gotH2, wantH2)
	}
	if gotT2 != wantT2 {
		t.Fatalf("got: %v, want:%v", gotT2, wantT2)

	}
	if !reflect.DeepEqual(gotVisitedPoints, wantVisitedPoints) {
		t.Fatalf("got: %v, want:%v", gotVisitedPoints, wantVisitedPoints)

	}

	h1 = point{-2, 2}
	t1 = point{-2, 1}
	h2, t2, _ := move(h1, t1, 1, []point{}, moveLeftf)

	wantH := point{-3, 2}
	if h2 != wantH {
		t.Fatalf("got: %v, want:%v", h2, point{-3, 2})
	}
	wantT := point{-2, 1}
	if t2 != wantT {
		t.Fatalf("got: %v, want:%v", t2, wantT)
	}

}

func TestMove(t *testing.T) {
	h1 := point{0, 0}
	t1 := point{0, 0}
	distance := 4
	wantH2 := point{-4, 0}
	wantT2 := point{-3, 0}
	wantVisitedPoints := []point{
		{0, 0},
		{-1, 0},
		{-2, 0},
		{-3, 0},
	}

	gotH2, gotT2, gotVisitedPoints := move(h1, t1, distance, []point{{0, 0}}, func(H point, i int) point {
		return point{H.x - i, H.y}
	})

	if gotH2 != wantH2 {
		t.Fatalf("got: %v, want:%v", gotH2, wantH2)
	}
	if gotT2 != wantT2 {
		t.Fatalf("got: %v, want:%v", gotT2, wantT2)

	}
	if !reflect.DeepEqual(gotVisitedPoints, wantVisitedPoints) {
		t.Fatalf("got: %v, want:%v", gotVisitedPoints, wantVisitedPoints)

	}

	h1 = point{-2, 2}
	t1 = point{-2, 1}
	h2, t2, _ := move(h1, t1, 1, []point{}, func(H point, i int) point {
		return point{H.x - i, H.y}
	})

	wantH := point{-3, 2}
	if h2 != wantH {
		t.Fatalf("got: %v, want:%v", h2, point{-3, 2})
	}
	wantT := point{-2, 1}
	if t2 != wantT {
		t.Fatalf("got: %v, want:%v", t2, wantT)
	}
}

func TestMoveUp(t *testing.T) {
	h1 := point{0, 0}
	t1 := point{0, 0}
	distance := 4
	wantH2 := point{0, 4}
	wantT2 := point{0, 3}
	wantVisitedPoints := []point{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}

	gotH2, gotT2, gotVisitedPoints := move(h1, t1, distance, []point{{0, 0}}, moveUpf)

	if gotH2 != wantH2 {
		t.Fatalf("got: %v, want:%v", gotH2, wantH2)
	}
	if gotT2 != wantT2 {
		t.Fatalf("got: %v, want:%v", gotT2, wantT2)

	}
	if !reflect.DeepEqual(gotVisitedPoints, wantVisitedPoints) {
		t.Fatalf("got: %v, want:%v", gotVisitedPoints, wantVisitedPoints)

	}

}

func TestMoveDown(t *testing.T) {
	h1 := point{0, 0}
	t1 := point{0, 0}
	distance := 4
	wantH2 := point{0, -4}
	wantT2 := point{0, -3}
	wantVisitedPoints := []point{
		{0, 0},
		{0, -1},
		{0, -2},
		{0, -3},
	}

	gotH2, gotT2, gotVisitedPoints := move(h1, t1, distance, []point{{0, 0}}, moveDownf)

	if gotH2 != wantH2 {
		t.Fatalf("got: %v, want:%v", gotH2, wantH2)
	}
	if gotT2 != wantT2 {
		t.Fatalf("got: %v, want:%v", gotT2, wantT2)

	}
	if !reflect.DeepEqual(gotVisitedPoints, wantVisitedPoints) {
		t.Fatalf("got: %v, want:%v", gotVisitedPoints, wantVisitedPoints)

	}

}
