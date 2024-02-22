package gm_item

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type ItemDescription struct {
	image       *ebiten.Image
	height      int
	width       int
	textPadding int
	title       string
}

type ItemDescriptionOpt struct {
	Height      int
	Width       int
	TextPadding int
}

func NewItemDescription(title string, source *gm_layer.Image, opt ItemDescriptionOpt) (*ItemDescription, error) {

	if opt.Height == 0 || opt.Width == 0 {
		return nil, ErrZeroDescriptionSize
	}

	if source == nil || source.TileXCount != 3 || source.TileYCount != 3 {
		return nil, ErrInvalidSourceImage
	}

	var d *ItemDescription = &ItemDescription{
		height:      opt.Height,
		width:       opt.Width,
		textPadding: opt.TextPadding,
		title:       title,
		image:       ebiten.NewImage(opt.Width*source.TileSize, opt.Height*source.TileSize),
	}

	d.makeImage(source)

	return d, nil
}

func (i *ItemDescription) makeImage(source *gm_layer.Image) {

	var (
		image *ebiten.Image
		op    *ebiten.DrawImageOptions
	)

	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			if x == 0 && y == 0 {
				image = source.GetTile(0, 0)
			} else if x == i.width-1 && y == 0 {
				image = source.GetTile(2, 0)
			} else if x == 0 && y == i.height-1 {
				image = source.GetTile(0, 2)
			} else if x == i.width-1 && y == i.height-1 {
				image = source.GetTile(2, 2)
			} else if x == 0 {
				image = source.GetTile(0, 1)
			} else if x == i.width-1 {
				image = source.GetTile(2, 1)
			} else if y == 0 {
				image = source.GetTile(1, 0)
			} else if y == i.height-1 {
				image = source.GetTile(1, 2)
			} else {
				image = source.GetTile(1, 1)
			}
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*source.TileSize), float64(y*source.TileSize))
			i.image.DrawImage(image, op)
		}
	}

	var (
		titleX, titleY int = i.textPadding, i.textPadding
	)

	ebitenutil.DebugPrintAt(i.image, i.title, titleX, titleY)
}

func (i *ItemDescription) Draw(x, y float64, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(i.image, op)
}
