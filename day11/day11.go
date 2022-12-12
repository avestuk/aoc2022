package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day11(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	monkeys := parseInput(s)

	//Monkey 0:
	//  Monkey inspects an item with a worry level of 79.
	//    Worry level is multiplied by 19 to 1501.
	//    Monkey gets bored with item. Worry level is divided by 3 to 500.
	//    Current worry level is not divisible by 23.
	//    Item with worry level 500 is thrown to monkey 3.
	//  Monkey inspects an item with a worry level of 98.
	//    Worry level is multiplied by 19 to 1862.
	//    Monkey gets bored with item. Worry level is divided by 3 to 620.
	//    Current worry level is not divisible by 23.
	//    Item with worry level 620 is thrown to monkey 3.
	//
	// In this example, the two most active monkeys inspected items 101 and 105
	// times. The level of monkey business in this situation can be found by
	// multiplying these together: 10605.

	// Figure out which monkeys to chase by counting how many items they inspect
	// over 20 rounds. What is the level of monkey business after 20 rounds of
	// stuff-slinging simian shenanigans?
	playRounds(monkeys, 20)

	_ = monkeyBusiness(monkeys)
	//return monkeyBusiness, nil
	return 0, nil
}

func Day11PartTwo(file string) (int64, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	monkeys := parseInput(s)

	commonDivisor := int64(1)
	for _, monkey := range monkeys {
		commonDivisor *= int64(monkey.test)
	}

	playRoundsPartTwo(monkeys, commonDivisor, 10000)

	monkeyBusiness := monkeyBusiness(monkeys)
	return monkeyBusiness, nil
}

func monkeyBusiness(monkeys []*monkey) int64 {
	var mostInspections, secondMostInspections int64
	for _, monkey := range monkeys {
		if monkey.inspections >= mostInspections {
			secondMostInspections = mostInspections
			mostInspections = monkey.inspections
		} else if monkey.inspections >= secondMostInspections {
			secondMostInspections = monkey.inspections
		}
	}
	return mostInspections * secondMostInspections
}

func playRounds(monkeys []*monkey, rounds int) {
	for i := 1; i <= rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := monkey.operation(item)
				worryLevel = worryLevel / 3

				catcher := monkey.ifFalse
				if catch := worryLevel % monkey.test; catch == 0 {
					catcher = monkey.ifTrue
				}
				monkeys[catcher].items = append(monkeys[catcher].items, worryLevel)
				monkey.inspections++
			}
			monkey.items = []int64{}
		}
	}
}

func playRoundsPartTwo(monkeys []*monkey, commonDivisor int64, rounds int) {
	for i := 1; i <= rounds; i++ {
		switch i {
		case 20, 1000, 2000, 3000, 4000, 5000, 6000:
			fmt.Printf("Round %d\n", i)
			fmt.Printf("===========\n")
			for i, monkey := range monkeys {
				fmt.Printf("monkey %d: %d\n", i, monkey.inspections)
			}
		}
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := monkey.operation(item) % commonDivisor

				catcher := monkey.ifFalse
				if catch := worryLevel % monkey.test; catch == 0 {
					catcher = monkey.ifTrue
				}
				monkeys[catcher].items = append(monkeys[catcher].items, worryLevel)
				monkey.inspections++
			}
			monkey.items = []int64{}
		}
	}
}

type monkey struct {
	items       []int64
	operation   func(old int64) int64
	test        int64
	ifTrue      int64
	ifFalse     int64
	inspections int64
}

func parseInput(s *bufio.Scanner) []*monkey {
	var (
		monkeyX   int64
		operator  string
		i         string
		operation func(old int64) int64
		test      int64
		ifTrue    int64
		ifFalse   int64
	)
	monkeyItems := []int64{}

	monkeys := []*monkey{}

	for s.Scan() {
		text := strings.TrimSpace(s.Text())
		switch {
		case strings.HasPrefix(text, "Monkey"):
			fmt.Sscanf(text, "Monkey %d:", &monkeyX)
		case strings.HasPrefix(text, "Starting items:"):
			items := strings.Split(s.Text(), ":")[1]
			for _, s := range strings.Split(items, ",") {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					panic(fmt.Sprintf("failed to parse: %s", s))
				}
				monkeyItems = append(monkeyItems, int64(i))
			}
		case strings.HasPrefix(text, "Operation:"):
			fmt.Sscanf(text, "Operation: new = old %s %s", &operator, &i)
			var (
				constant int64
			)
			if i != "old" {
				c, err := strconv.Atoi(i)
				if err != nil {
					panic(fmt.Sprintf("failed to parse: %s", i))
				}
				constant = int64(c)
			}
			switch operator {
			case "+":
				operation = func(old int64) int64 { return old + constant }
			case "-":
				operation = func(old int64) int64 { return old - constant }
			case "*":
				if i == "old" {
					operation = func(old int64) int64 { return old * old }
				} else {
					operation = func(old int64) int64 { return old * constant }
				}
			case "/":
				if i == "old" {
					operation = func(old int64) int64 { return 1 }
				} else {
					operation = func(old int64) int64 { return old / constant }
				}
			}
		case strings.HasPrefix(text, "Test:"):
			fmt.Sscanf(text, "Test: divisible by %d", &test)
		case strings.HasPrefix(text, "If true:"):
			fmt.Sscanf(text, "If true: throw to monkey %d", &ifTrue)
		case strings.HasPrefix(text, "If false:"):
			fmt.Sscanf(text, "If false: throw to monkey %d", &ifFalse)
			monkeys = append(monkeys, &monkey{
				items:       monkeyItems,
				operation:   operation,
				test:        test,
				ifTrue:      ifTrue,
				ifFalse:     ifFalse,
				inspections: 0,
			})

			monkeyItems = []int64{}
		}
	}
	return monkeys
}
