package gm_layer

import "errors"

var (
	ErrMapParsingEmptyLine     error = errors.New("raw parsing err: empty string")
	ErrMapParsingDifferentLine error = errors.New("raw parsing err: empty string")
)
