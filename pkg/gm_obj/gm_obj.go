package gm_obj

import (
	"github.com/hajimehoshi/ebiten/v2"
	gm_layer "github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type GameObjOptions struct {
	X         int
	Y         int
	Width     int
	Height    int
	Raw       []int
	ImagePath string
}

type GameObj struct {
	Name   string
	X      int
	Y      int
	Width  int
	Height int
	raw    []int
	layer  *gm_layer.Layer
}

func NewGameObj(name string, opt GameObjOptions) *GameObj {
	var obj *GameObj = &GameObj{
		X:      opt.X,
		Y:      opt.Y,
		Width:  opt.Width,
		Height: opt.Height,
		Name:   name,
		raw:    opt.Raw,
		layer:  nil,
	}

	obj.makeLayer(opt.ImagePath)

	return obj
}

func (g *GameObj) Draw(screen *ebiten.Image) {
	g.layer.Draw(screen)
}

func (g *GameObj) makeLayer(imagePath string) {
	var (
		layer []int = []int{}
	)

	// Fill top border
	for h := 0; h < g.Y; h++ {
		for w := 0; w < stgs.TileXCount; w++ {
			layer = append(layer, 0)
		}
	}

	var tileCounter int = 0

	for h := 0; h < g.Height; h++ {
		// Fill left border
		for w := 0; w < g.X; w++ {
			layer = append(layer, 0)
		}

		// Fill content
		for w := g.X; w <= g.Width; w++ {
			layer = append(layer, g.raw[tileCounter])
			tileCounter++
		}

		// Fill right border
		for w := g.X + g.Width; w < stgs.TileXCount; w++ {
			layer = append(layer, 0)
		}
	}

	g.layer = gm_layer.NewLayer(g.Name, layer, imagePath)
}
