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
		file      *os.File
		scanner   *bufio.Scanner
		digit     int
		line      string
		err       error
		lineWidth int
		width     int = 0
		height    int = 0
		raw       []int
	)

	file, err = os.Open(tilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		lineWidth = 0
		if len(line) == 0 {
			return nil, ErrMapParsingEmptyLine
		}
		for _, tile := range strings.Split(line, " ") {
			digit, err = strconv.Atoi(tile)
			if err != nil {
				return nil, err
			}
			lineWidth++
			raw = append(raw, digit)
		}
		if width != 0 && width != lineWidth {
			return nil, ErrMapParsingDifferentLine
		}
		width = lineWidth
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
