package gm_inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Slot struct {
	Value IItem
	image *gm_layer.Image
}

func NewSlot(image *gm_layer.Image) *Slot {
	return &Slot{
		Value: nil,
		image: image,
	}
}

func (s *Slot) IsFree() bool {
	return s.Value == nil
}

func (s *Slot) SetItem(item IItem) {
	s.Value = item
}

func (s *Slot) Draw(x, y float64, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(s.image.Inst, op)
}
