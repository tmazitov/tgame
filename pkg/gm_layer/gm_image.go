package gm_layer

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Image struct {
	Inst       *ebiten.Image
	TileXCount int
	TileYCount int
}

func NewImage(inst *ebiten.Image) *Image {
	if inst == nil {
		return nil
	}
	return &Image{
		Inst:       inst,
		TileXCount: inst.Bounds().Dx() / stgs.TileSize,
		TileYCount: inst.Bounds().Dy() / stgs.TileSize,
	}
}

func NewImageByPath(path string) (*Image, error) {
	var (
		img  image.Image
		file *os.File
		err  error
	)

	if file, err = os.Open(path); err != nil {
		return nil, err
	}
	defer file.Close()

	if img, _, err = image.Decode(file); err != nil {
		return nil, err
	}

	return NewImage(ebiten.NewImageFromImage(img)), nil
}
