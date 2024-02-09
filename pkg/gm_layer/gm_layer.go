package gm_layer

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Layer struct {
	Name     string
	raw      *Raw
	image    *Image
	TileSize int
}

type LayerBorder struct {
	X      float64
	Y      float64
	Width  int
	Height int
}

func NewLayer(name string, rawPath string, imagePath string, tileSize int) (*Layer, error) {

	var (
		image *Image
		raw   *Raw
		err   error
	)

	image, err = NewImageByPath(imagePath)
	if err != nil {
		return nil, err
	}

	raw, err = NewRaw(rawPath)
	if err != nil {
		return nil, err
	}

	return &Layer{
		raw:      raw,
		Name:     name,
		TileSize: tileSize,
		image:    image,
	}, err
}

func NewLayerByRaw(name string, raw *Raw, imagePath string, tileSize int) (*Layer, error) {

	var (
		image *Image
		err   error
	)

	image, err = NewImageByPath(imagePath)
	if err != nil {
		return nil, err
	}

	return &Layer{
		raw:      raw,
		Name:     name,
		image:    image,
		TileSize: tileSize,
	}, nil
}

func (l *Layer) GetSizes() (int, int) {
	return l.raw.height, l.raw.width
}

func (l *Layer) GetValue(b LayerBorder) []int {
	var (
		initX         float64 = b.X / float64(l.TileSize)
		initY         float64 = b.Y / float64(l.TileSize)
		initWidth     float64 = float64(b.Width / l.TileSize)
		initHeight    float64 = float64(b.Height / l.TileSize)
		layerHeight   float64 = float64(l.raw.height)
		layerWidth    float64 = float64(l.raw.width)
		limitValues   []int   = []int{}
		limitXCounter float64 = 0
		limitYCounter float64 = 0
	)

	if initHeight+initY > float64(l.raw.height) || initWidth+initX > float64(l.raw.width) {
		return l.raw.tiles
	}

	for _, tile := range l.raw.tiles {
		if limitXCounter >= initX && limitXCounter < initX+initWidth &&
			limitYCounter >= initY && limitYCounter < initY+initHeight {
			limitValues = append(limitValues, tile)
		}
		if limitXCounter == layerWidth {
			limitXCounter = 0
			limitYCounter += 1
			fmt.Printf("tile map size : %d\n", len(limitValues))
		} else {
			limitXCounter += 1
		}
		if limitYCounter == initHeight+initY || limitYCounter == layerHeight {
			break
		}
	}

	return limitValues
}

func GetCoordsByTile(image *Image, tile int) (int, int) {
	return (tile % image.TileXCount) * stgs.TileSize, (tile / image.TileXCount) * stgs.TileSize
}

func GetTranslateByTile(index int) (float64, float64) {
	return float64((index % stgs.TileXCount) * stgs.TileSize), float64((index / stgs.TileXCount) * stgs.TileSize)
}

func (l *Layer) Draw(screen *ebiten.Image, b LayerBorder) {
	for i, tile := range l.GetValue(b) {
		if tile == 0 {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(GetTranslateByTile(i))

		sx, sy := GetCoordsByTile(l.image, tile-1)
		// fmt.Printf("draw %s : %d %d : %d %d\n", l.name, sx, sy, i, tile)
		screen.DrawImage(l.image.Inst.SubImage(image.Rect(sx, sy, sx+stgs.TileSize, sy+stgs.TileSize)).(*ebiten.Image), op)
	}
}
