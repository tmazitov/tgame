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

func NewGameObj(name string, opt GameObjOptions) (*GameObj, error) {
	var obj *GameObj = &GameObj{
		X:      opt.X,
		Y:      opt.Y,
		Width:  opt.Width,
		Height: opt.Height,
		Name:   name,
		raw:    opt.Raw,
		layer:  nil,
	}

	if err := obj.makeLayer(opt.ImagePath); err != nil {
		return nil, err
	}
	return obj, nil
}

func (g *GameObj) Draw(screen *ebiten.Image) {
	g.layer.Draw(screen)
}

func (g *GameObj) makeLayer(imagePath string) error {
	var (
		tiles []int = []int{}
		err   error
	)

	// Fill top border
	for h := 0; h < g.Y; h++ {
		for w := 0; w < stgs.TileXCount; w++ {
			tiles = append(tiles, 0)
		}
	}

	var tileCounter int = 0

	for h := 0; h < g.Height; h++ {
		// Fill left border
		for w := 0; w < g.X; w++ {
			tiles = append(tiles, 0)
		}

		// Fill content
		for w := g.X; w <= g.Width; w++ {
			tiles = append(tiles, g.raw[tileCounter])
			tileCounter++
		}

		// Fill right border
		for w := g.X + g.Width; w < stgs.TileXCount; w++ {
			tiles = append(tiles, 0)
		}
	}

	g.layer, err = gm_layer.NewLayerByRaw(
		g.Name,
		gm_layer.NewRawByTiles(tiles, g.Width, g.Height),
		imagePath,
	)
	return err
}
