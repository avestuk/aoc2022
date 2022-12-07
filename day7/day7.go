package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/avestuk/aoc2022/pkg/parse"
)

func Day7(file string) (int, error) {
	s, close, err := parse.ParseInput(file)
	if err != nil {
		return 0, fmt.Errorf("failed to parse file, got error: %s", err)
	}
	defer close()

	for s.Scan() {
		return 0, nil

	}

	return 0, nil
}

func parseLines(s *bufio.Scanner) *directory {
	root := newDirectory("/", nil)
	var (
		cwd        *directory
		lsContents bool
	)

	for s.Scan() {
		// This is a command

		text := s.Text()
		switch {
		case strings.HasPrefix(text, "$ cd"):
			cwdString := strings.Split(text, " ")[2]

			if cwdString == "/" {
				cwd = root
				continue
			}

			// Check this works
			if cwdString == ".." {
				cwd = cwd.parent
				continue
			}

			subDirFound := false
			for _, subDir := range cwd.subDirs {
				subDir := subDir
				if subDir.name == cwdString {
					cwd = subDir
					subDirFound = true
				}
			}

			if !subDirFound {
				panic(fmt.Sprintf("subdir: %s not found in parent: %s", cwdString, cwd.name))
			}

		// This is a directory
		case strings.HasPrefix(text, "$ ls"):
			lsContents = true

		case strings.HasPrefix(text, "dir"):
			if !lsContents {
				panic("got dir but lsContents != true")
			}

			subDir := strings.Split(text, " ")[1]
			cwd.subDirs = append(cwd.subDirs, newDirectory(subDir, cwd))
		default:
			// This is a file listing
			lsOutput := strings.Split(text, " ")
			fileSizeStr, _ := lsOutput[0], lsOutput[1]
			fileSize, err := strconv.Atoi(fileSizeStr)
			if err != nil {
				panic(fmt.Sprintf("failed to convert fileSize: %s", fileSizeStr))
			}
			cwd.updateSize(fileSize)
		}

	}

	return root
}

type directory struct {
	name    string
	size    int
	parent  *directory
	subDirs []*directory
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:    name,
		parent:  parent,
		subDirs: []*directory{},
	}
}

func (d *directory) updateSize(size int) {
	d.size += size

	if d.parent != nil {
		d.parent.updateSize(size)
	}
}

func (d *directory) filterSizes(sizeLimit int) int {
	total := 0
	if len(d.subDirs) != 0 {
		for _, subDir := range d.subDirs {
			total += subDir.filterSizes(sizeLimit)
		}
	}

	if d.size <= sizeLimit {
		return total + d.size
	}
	return total
}
