package gm_layer

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image struct {
	Inst       *ebiten.Image
	TileSize   int
	TileXCount int
	TileYCount int
}

func NewImage(inst *ebiten.Image, tileSize int) (*Image, error) {

	if tileSize == 0 {
		return nil, ErrImageTileSizeZero
	}

	if inst == nil {
		return nil, ErrImageInstNil
	}
	return &Image{
		Inst:       inst,
		TileSize:   tileSize,
		TileXCount: inst.Bounds().Dx() / tileSize,
		TileYCount: inst.Bounds().Dy() / tileSize,
	}, nil
}

func NewImageByPath(path string, tileSize int) (*Image, error) {
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

	return NewImage(ebiten.NewImageFromImage(img), tileSize)
}

func (i *Image) Height() int {
	return i.Inst.Bounds().Dy()
}

func (i *Image) Width() int {
	return i.Inst.Bounds().Dx()
}

func (i *Image) GetTile(x, y int) *ebiten.Image {

	if x < 0 || x >= i.TileXCount || y < 0 || y >= i.TileYCount {
		return nil
	}

	return i.Rect(x*i.TileSize, y*i.TileSize, i.TileSize, i.TileSize)
}

func (i *Image) Rect(x, y, height, width int) *ebiten.Image {
	return i.Inst.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
}
