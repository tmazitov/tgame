package gm_inventory

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ErrZeroSize           error = errors.New("inventory error : size is zero")
	ErrSlotImagePathEmpty error = errors.New("inventory error : slot image path is empty")
	ErrSlotImageSizeZero  error = errors.New("inventory error : slot image size is zero")
)

type IItem interface {
	Draw(screen *ebiten.Image)
	SetPosition(x, y float64)
	GetPosition() (float64, float64)
	GetIsMoving() bool
}
