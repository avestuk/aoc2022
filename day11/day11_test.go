package day11

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func TestParse(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse file, got error: %s", err)
	}
	defer close()

	want := &monkey{
		items: []int{79, 98},
		operation: func(old int) int {
			return old * 19
		},
		test:        23,
		inspections: 0,
	}

	got := parseInput(s)

	if reflect.DeepEqual(got[0], want) {
		t.Fatalf("got: %v\nwant: %v", *got[0], *want)
	}
}

func TestPlayRounds(t *testing.T) {
	s, close, err := parse.ParseInput("./test.txt")
	if err != nil {
		t.Fatalf("failed to parse file, got error: %s", err)
	}
	defer close()

	monkeys := parseInput(s)

	playRounds(monkeys, 1)

	if err := itemsEqual(monkeys[0].items, []int{20, 23, 27, 26}); err != nil {
		t.Fatal("monkey 0 items not right")
	}

	if err := itemsEqual(monkeys[1].items, []int{2080, 25, 167, 207, 401, 1046}); err != nil {
		t.Fatal("monkey 0 items not right")
	}

	if len(monkeys[2].items) != 0 || len(monkeys[3].items) != 0 {
		t.Fatal("monkey 2 or 3 not right")
	}

	playRounds(monkeys, 9)

	if err := itemsEqual(monkeys[0].items, []int{91, 16, 20, 98}); err != nil {
		t.Fatal("monkey 0 items not right")
	}
	if err := itemsEqual(monkeys[1].items, []int{481, 245, 22, 26, 1092, 30}); err != nil {
		t.Fatal("monkey 0 items not right")
	}

	if len(monkeys[2].items) != 0 || len(monkeys[3].items) != 0 {
		t.Fatal("monkey 2 or 3 not right")
	}

	playRounds(monkeys, 5)

	if err := itemsEqual(monkeys[0].items, []int{83, 44, 8, 184, 9, 20, 26, 102}); err != nil {
		t.Fatal("monkey 0 items not right")
	}
	if err := itemsEqual(monkeys[1].items, []int{110, 36}); err != nil {
		t.Fatal("monkey 0 items not right")
	}

	if len(monkeys[2].items) != 0 || len(monkeys[3].items) != 0 {
		t.Fatal("monkey 2 or 3 not right")
	}

	playRounds(monkeys, 5)

	if err := itemsEqual(monkeys[0].items, []int{10, 12, 14, 26, 34}); err != nil {
		t.Fatal("monkey 0 items not right")
	}
	if err := itemsEqual(monkeys[1].items, []int{245, 93, 53, 199, 115}); err != nil {
		t.Fatal("monkey 0 items not right")
	}

	if len(monkeys[2].items) != 0 || len(monkeys[3].items) != 0 {
		t.Fatal("monkey 2 or 3 not right")
	}

	gotMB := monkeyBusiness(monkeys)
	wantMB := 10605

	if gotMB != wantMB {
		t.Fatalf("got MB: %d want: %d", gotMB, wantMB)
	}

}

func itemsEqual(a, b []int) error {
	for i, item := range a {
		if b[i] != item {
			return fmt.Errorf("item: %d at index: %d do not match", item, i)
		}
	}

	return nil
}
