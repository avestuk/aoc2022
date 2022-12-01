package day1

import (
	"testing"
)

func TestMostCalories(t *testing.T) {
	elfIndex, mostCalories, err := Day1("./test.txt")
	if err != nil {
		t.Fatalf("got err: %s", err)
	}

	if elfIndex != 4 && mostCalories != 24000 {
		t.Errorf("got elfIndex: %d, want: %d,\n got mostCalories: %d, want: %d", elfIndex, 4, mostCalories, 24000)
	}
}

func TestTop3MostCalories(t *testing.T) {

	want := 45000
	elfCalories := map[int]int{0: 4000, 1: 6000, 2: 11000, 3: 24000, 4: 10000}

	mostCalories := topThree(elfCalories)

	got := 0
	for _, c := range mostCalories {
		got += c
	}

	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}

}
