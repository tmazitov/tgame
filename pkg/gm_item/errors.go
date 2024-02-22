package gm_item

import "errors"

var (
	ErrInvalidParams       error = errors.New("item error : invalid params")
	ErrInvalidSourceImage  error = errors.New("item error : invalid source image")
	ErrZeroDescriptionSize error = errors.New("item error : zero description size")
)
