package gm_layer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Raw struct {
	tiles  []int
	height int
	width  int
}

func NewRaw(tilePath string) (*Raw, error) {

	var (
		file    *os.File
		scanner *bufio.Scanner
		digit   int
		line    string
		err     error
		width   int = 0
		height  int = 0
		raw     []int
	)

	file, err = os.Open(tilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			return nil, ErrMapParsingEmptyLine
		}
		if width != 0 && len(line) != width {
			return nil, ErrMapParsingDifferentLine
		}
		for _, char := range strings.Split(line, " ") {
			digit, err = strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			raw = append(raw, digit)
		}
		width = len(line)
		height++
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return &Raw{
		tiles:  raw,
		height: height,
		width:  width,
	}, nil
}

func NewRawByTiles(tiles []int, height, width int) *Raw {
	return &Raw{
		tiles:  tiles,
		height: height,
		width:  width,
	}
}
