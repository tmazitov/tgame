package gm_item

import "errors"

var (
	ErrInvalidParams       error = errors.New("item error : invalid params")
	ErrInvalidSourceImage  error = errors.New("item error : invalid source image")
	ErrZeroDescriptionSize error = errors.New("item error : zero description size")
)

var (
	ErrEmptyDescriptionSourcePath error = errors.New("item collection error: empty description source path")
	ErrZeroItemSize               error = errors.New("item collection error: zero item size")
	ErrEmptyJsonPath              error = errors.New("item collection error: empty json path")
)

var (
	ErrEmptyConfigPath error = errors.New("item collection storage error: empty config path")
)
