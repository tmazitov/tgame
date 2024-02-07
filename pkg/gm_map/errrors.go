package gm_map

import "errors"

var (
	ErrMapWithoutBackground error = errors.New("map err: background is nil")
)
