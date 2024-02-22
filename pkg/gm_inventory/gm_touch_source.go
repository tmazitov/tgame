package gm_inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TouchSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

type MouseTouchSource struct {
	ID ebiten.MouseButton
}

func (m *MouseTouchSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseTouchSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(m.ID)
}

type OriginTouchSource struct {
	ID ebiten.TouchID
}

func (t *OriginTouchSource) Position() (int, int) {
	return ebiten.TouchPosition(t.ID)
}

func (t *OriginTouchSource) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.ID)
}
