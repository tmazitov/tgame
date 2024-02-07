package gm_layer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Layer struct {
	name  string
	raw   *Raw
	image *Image
}

func NewLayer(name string, rawPath string, imagePath string) (*Layer, error) {

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
		raw:   raw,
		name:  name,
		image: image,
	}, err
}

func NewLayerByRaw(name string, raw *Raw, imagePath string) (*Layer, error) {

	var (
		image *Image
		err   error
	)

	image, err = NewImageByPath(imagePath)
	if err != nil {
		return nil, err
	}

	return &Layer{
		raw:   raw,
		name:  name,
		image: image,
	}, nil
}

func (l *Layer) GetValue() []int {
	return l.raw.tiles
}

func GetCoordsByTile(image *Image, tile int) (int, int) {
	return (tile % image.TileXCount) * stgs.TileSize, (tile / image.TileXCount) * stgs.TileSize
}

func GetTranslateByTile(index int) (float64, float64) {
	return float64((index % stgs.TileXCount) * stgs.TileSize), float64((index / stgs.TileXCount) * stgs.TileSize)
}

func (l *Layer) Draw(screen *ebiten.Image) {
	for i, tile := range l.GetValue() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(GetTranslateByTile(i))

		sx, sy := GetCoordsByTile(l.image, tile)
		// fmt.Printf("draw %s : %d %d : %d %d\n", l.name, sx, sy, i, tile)
		screen.DrawImage(l.image.Inst.SubImage(image.Rect(sx, sy, sx+stgs.TileSize, sy+stgs.TileSize)).(*ebiten.Image), op)
	}
}
