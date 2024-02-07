package gm_map

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Ground struct {
	background *gm_layer.Layer
	additional []*gm_layer.Layer
}

type GroundLayerImages struct {
	Background string
}

func GroundRaw(height, width int) []int {
	var ground []int = []int{}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			ground = append(ground, 0)
		}
	}
	return ground
}

func NewGround(background *gm_layer.Layer) *Ground {

	return &Ground{
		additional: nil,
		background: background,
	}
}

func (g *Ground) AddLayer(layer *gm_layer.Layer) {
	if layer == nil {
		return
	}
	if g.additional == nil {
		g.additional = []*gm_layer.Layer{}
	}
	g.additional = append(g.additional, layer)
}

func (g *Ground) Draw(screen *ebiten.Image) {
	if g.background != nil {
		g.background.Draw(screen)
	}

	if g.additional != nil {
		for _, layer := range g.additional {
			layer.Draw(screen)
		}
	}
}
