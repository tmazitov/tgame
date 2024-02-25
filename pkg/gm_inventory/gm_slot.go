package gm_inventory

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Slot struct {
	Item  *gm_item.Item
	image *gm_layer.Image
}

func NewSlot(image *gm_layer.Image) *Slot {
	return &Slot{
		Item:  nil,
		image: image,
	}
}

func (s *Slot) IsFree() bool {
	return s.Item == nil
}

func (s *Slot) SetItem(item *gm_item.Item) {
	s.Item = item
}

func (s *Slot) Draw(x, y float64, font *gm_font.Font, screen *ebiten.Image) {

	var (
		itemX float64
		itemY float64
	)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(s.image.Inst, op)

	if s.Item != nil && !s.Item.GetIsMoving() {
		itemX, itemY = s.Item.GetPosition()
		if itemX != x && itemY != y {
			s.Item.SetPosition(x, y)
		}
		s.Item.Draw(screen, nil)
		if s.Item.GetAmount() != 1 {
			size := s.image.Inst.Bounds().Dx()
			x := int(x) + size - 12
			y := int(y) + size - 14
			font.Print(screen, fmt.Sprintf("%d", s.Item.GetAmount()), image.Pt(x, y), nil)
		}
	}
}
