package gm_layer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Layer struct {
	tiles []int
	name  string
	image *Image
}

func NewLayer(name string, tiles []int, image *Image) *Layer {
	return &Layer{
		tiles: tiles,
		name:  name,
		image: image,
	}
}

func (l *Layer) GetValue() []int {
	return l.tiles
}

func GetCoordsByTile(image *Image, tile int) (int, int) {
	return (tile % image.TileXCount) * stgs.TileSize, (tile / image.TileXCount) * stgs.TileSize
}

func GetTranslateByTile(index int) (float64, float64) {
	return float64((index % stgs.TileXCount) * stgs.TileSize), float64((index / stgs.TileXCount) * stgs.TileSize)
}

func (l *Layer) Draw(screen *ebiten.Image) {
	for i, tile := range l.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(GetTranslateByTile(i))

		sx, sy := GetCoordsByTile(l.image, tile)
		screen.DrawImage(l.image.Inst.SubImage(image.Rect(sx, sy, sx+stgs.TileSize, sy+stgs.TileSize)).(*ebiten.Image), op)
	}
}
