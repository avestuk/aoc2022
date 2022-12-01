package parse

import (
	"bufio"
	"os"
)

func ParseInput(f string) (*bufio.Scanner, func() error, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, file.Close, nil
}
