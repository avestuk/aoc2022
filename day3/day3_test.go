package day3

import (
	"testing"
)

func TestSplitLine(t *testing.T) {
	t.Run("first line", func(t *testing.T) {
		line := "vJrwpWtwJgWrhcsFMMfFFhFp"
		got1, got2 := splitLine(line)
		if len(got1) != len(got2) && len(got1)+len(got2) != len(line) {
			t.Fatalf("line lengths not equal: %d, %d, %d", len(got1), len(got2), len(line))
		}
		want1, want2 := "vJrwpWtwJgWr", "hcsFMMfFFhFp"

		if got1 != want1 && got2 != want2 {
			t.Fatalf("got1: %s, want1: %s\n got2: %s, want2: %s", got1, want1, got2, want2)
		}

	})
}

func TestFindItem(t *testing.T) {
	got := findItem("vJrwpWtwJgWrhcsFMMfFFhFp")
	want := 'p'

	if got != want {
		t.Fatalf("got: %s, want: %s", string(got), string(want))
	}
}

func TestGetPriority(t *testing.T) {
	got := getPriority('A')
	want := 27

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = getPriority('Z')
	want = 52

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = getPriority('a')
	want = 1

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}

	got = getPriority('z')
	want = 26

	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestFindIntersection(t *testing.T) {
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}
	got := findIntersection(rucksacks)
	want := 'r'

	if got != want {
		t.Fatalf("got: %s, want: %s", string(got), string(want))
	}

	rucksacks = []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	got = findIntersection(rucksacks)
	want = 'Z'

	if got != want {
		t.Fatalf("got: %s, want: %s", string(got), string(want))
	}

}
