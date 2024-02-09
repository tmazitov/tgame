package gm_map

import "errors"

var (
	ErrMapWithoutBackground error = errors.New("map err: background is nil")
	ErrMapWithZeroTileSize  error = errors.New("map err: tile size is 0")
)
