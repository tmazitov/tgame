package ground

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Ground struct {
	layers []*gm_layer.Layer
}

func NewGround(layers []*gm_layer.Layer) *Ground {

	for _, layer := range layers {
		if layer == nil {
			return nil
		}
	}

	return &Ground{
		layers: layers,
	}
}

func (g *Ground) Draw(screen *ebiten.Image) {
	for _, layer := range g.layers {
		layer.Draw(screen)
	}
}

func GroundRaw(tileNumber int) []int {
	var ground []int = []int{}

	for h := 0; h < stgs.ScreenHeight; h += stgs.TileSize {
		for w := 0; w < stgs.ScreenWidth; w += stgs.TileSize {
			ground = append(ground, 0)
		}
	}
	return ground
}
