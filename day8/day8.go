package day8

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day8(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	trees := parseLines(s)
	count := countVisible(trees)

	return count, nil
}

func Day8PartTwo(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	trees := parseLines(s)
	mostScenic := mostScenic(trees)

	return mostScenic, nil
}

func parseLines(s *bufio.Scanner) [][]int {
	trees := make([][]int, 0)
	lineCount := 0
	for s.Scan() {
		if len(s.Text()) == 0 {
			panic("got 0 length line")
		}

		line := []int{}
		for _, r := range s.Text() {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				panic(fmt.Sprintf("failed to convert: %s to int, got err: %s", string(r), err))
			}

			line = append(line, i)
		}
		trees = append(trees, line)
		lineCount++
	}

	return trees
}

func countVisible(trees [][]int) int {
	visibleCount := 0
	visibleCount += len(trees) * 2
	visibleCount += ((len(trees[0]) - 2) * 2)

	for i, row := range trees {
		if i == 0 || i == len(trees)-1 {
			continue
		}
	Row:
		for j, tree := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}
			var (
				visibleLeft, visibleRight = true, true
				visibleTop, visibleBottom = true, true
			)
			for k, horizontalComparison := range row {
				// If we have reached the tree itself check if it's visible
				if k == j {
					if visibleLeft {
						visibleCount += 1
						continue Row
					}
					continue
				}

				if k == len(row)-1 {
					if visibleRight && tree > horizontalComparison {
						visibleCount += 1
						continue Row
					}
				}

				// If the tree is larger than the one to its left continue checking
				if k < j {
					if tree > horizontalComparison {
						continue
					}
					visibleLeft = false
				}
				// If the tree is larger than the one to its right continue checking
				if k > j {
					if tree > horizontalComparison {
						continue
					}
					visibleRight = false
				}
			}

			for ii := 0; ii < len(trees); ii++ {
				// Get tree in the same colum
				verticalComparison := trees[ii][j]
				if ii == i {
					if visibleTop {
						visibleCount += 1
						continue Row
					}
					continue
				}

				if ii == len(trees)-1 {
					if visibleBottom && tree > verticalComparison {
						visibleCount += 1
						continue Row
					}
					continue
				}

				if ii < i {
					if tree > verticalComparison {
						continue
					}
					visibleTop = false
				}

				if ii > i {
					if tree > verticalComparison {
						continue
					}
					visibleBottom = false
				}
			}
		}
	}

	return visibleCount
}

func mostScenic(trees [][]int) int {
	mostScenic := 0

	for i, row := range trees {
		if i == 0 || i == len(trees)-1 {
			continue
		}
		for j, tree := range row {
			// Stop if you reach an edge or a tree that is the same height or taller
			// Include that three
			var (
				countLeft, countRight, countTop, countBottom int
			)

			// check Left
			// TODO Account for edge
			for k := j - 1; k >= 0; k-- {
				horizontalComparison := row[k]
				if tree <= horizontalComparison {
					countLeft += 1
					break
				}
				countLeft += 1
			}

			// Check Right
			for k := j + 1; k <= len(row)-1; k++ {
				horizontalComparison := row[k]
				if tree <= horizontalComparison {
					countRight += 1
					break
				}
				countRight += 1
			}

			// Check top
			for k := i - 1; k >= 0; k-- {
				verticalComparison := trees[k][j]
				if tree <= verticalComparison {
					countTop += 1
					break
				}
				countTop += 1
			}

			// Check bottom
			for k := i + 1; k <= len(trees)-1; k++ {
				verticalComparison := trees[k][j]
				if tree <= verticalComparison {
					countBottom += 1
					break
				}
				countBottom += 1
			}

			total := countBottom * countTop * countRight * countLeft
			if total > mostScenic {
				mostScenic = total
			}

		}

	}
	return mostScenic
}
