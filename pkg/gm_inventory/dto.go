package gm_inventory

import (
	"errors"
)

var (
	ErrZeroSize           error = errors.New("inventory error : size is zero")
	ErrNilFont            error = errors.New("inventory error : font is nil")
	ErrSlotImagePathEmpty error = errors.New("inventory error : slot image path is empty")
	ErrSlotImageSizeZero  error = errors.New("inventory error : slot image size is zero")
)
