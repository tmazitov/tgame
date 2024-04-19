package gm_inventory

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type SlotOptions struct {
	ItemCollection string
}

type Slot struct {
	Item           *gm_item.Item
	ItemCollection string
	image          *gm_layer.Image
	x              float64
	y              float64
}

func NewSlot(x, y float64, image *gm_layer.Image, opt *SlotOptions) *Slot {

	var (
		collection string
	)

	if opt != nil {
		collection = opt.ItemCollection
	} else {
		collection = ""
	}

	return &Slot{
		Item:           nil,
		image:          image,
		x:              x,
		y:              y,
		ItemCollection: collection,
	}
}

func (s *Slot) IsFree() bool {
	return s.Item == nil
}

func (s *Slot) SetItem(item *gm_item.Item) bool {
	if item == nil {
		s.Item = nil
		return true
	}

	if s.ItemCollection != "" && item.Collection != s.ItemCollection {
		return false
	}
	s.Item = item
	return true
}

func (s *Slot) GetPosition() (float64, float64) {
	return s.x, s.y
}

func (s *Slot) Draw(font *gm_font.Font, screen *ebiten.Image) {

	var (
		itemX float64
		itemY float64
	)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.image.Inst, op)

	if s.Item != nil && !s.Item.GetIsMoving() {
		itemX, itemY = s.Item.GetPosition()
		if itemX != s.x && itemY != s.y {
			s.Item.SetPosition(s.x, s.y)
		}
		s.Item.Draw(screen, nil)
		if s.Item.GetAmount() != 1 {
			size := s.image.Inst.Bounds().Dx()
			x := int(s.x) + size - 12
			y := int(s.y) + size - 14
			font.Print(screen, fmt.Sprintf("%d", s.Item.GetAmount()), image.Pt(x, y), nil)
		}
	}
}
