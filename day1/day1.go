package day1

import (
	"fmt"
	"strconv"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day1(file string) (int, int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	elfCalories := map[int]int{}
	elfCount := 1
	for s.Scan() {
		if len(s.Text()) == 0 {
			// New Elf
			elfCount += 1
			continue
		}

		calories, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse calories: %s, got err: %s", s.Text(), err)
		}

		elfCalories[elfCount] += calories
	}

	mostCalories := 0
	elfIndex := 0
	for elf, calories := range elfCalories {
		if calories > mostCalories {
			mostCalories = calories
			elfIndex = elf
		}
	}

	return elfIndex, mostCalories, nil
}

func Day1Part2(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	elfCalories := map[int]int{}
	elfCount := 1
	for s.Scan() {
		if len(s.Text()) == 0 {
			// New Elf
			elfCount += 1
			continue
		}

		calories, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, fmt.Errorf("failed to parse calories: %s, got err: %s", s.Text(), err)
		}

		elfCalories[elfCount] += calories
	}

	calorieCount := 0
	for _, calories := range topThree(elfCalories) {
		calorieCount += calories
	}

	return calorieCount, nil
}

func topThree(elfCalories map[int]int) []int {
	mostCalories := []int{0, 0, 0}
elfCalories:
	for _, calories := range elfCalories {
		for i, mc := range mostCalories {
			if calories >= mc {
				mostCalories = append(mostCalories[:i+1], mostCalories[i:2]...)
				mostCalories[i] = calories
				continue elfCalories
			}
		}
	}

	return mostCalories
}
