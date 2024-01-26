package gm_layer

import (
	"github.com/hajimehoshi/ebiten/v2"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Image struct {
	Inst       *ebiten.Image
	TileXCount int
	TileYCount int
}

func NewImage(inst *ebiten.Image) *Image {
	return &Image{
		Inst:       inst,
		TileXCount: inst.Bounds().Dx() / stgs.TileSize,
		TileYCount: inst.Bounds().Dy() / stgs.TileSize,
	}
}
