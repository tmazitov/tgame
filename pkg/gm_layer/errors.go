package gm_layer

import "errors"

var (
	ErrMapParsingEmptyLine     error = errors.New("raw parsing err: empty string")
	ErrMapParsingDifferentLine error = errors.New("raw parsing err: different line length")
	ErrImageTileSizeZero       error = errors.New("image error: tile size is zero")
	ErrImageInstNil            error = errors.New("image error: image instance is nil")
)
