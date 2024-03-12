package gm_item

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type ItemDescription struct {
	image       *ebiten.Image
	height      int
	width       int
	textPadding int
	title       string
	font        *gm_font.Font
}

type ItemDescriptionOpt struct {
	Height      int
	Width       int
	TextPadding int
	Font        *gm_font.Font
}

func NewItemDescription(title string, source *gm_layer.Image, opt ItemDescriptionOpt) (*ItemDescription, error) {

	if opt.Height == 0 || opt.Width == 0 {
		return nil, ErrZeroDescriptionSize
	}

	if source == nil || source.TileXCount != 5 || source.TileYCount != 5 {
		return nil, ErrInvalidSourceImage
	}

	var d *ItemDescription = &ItemDescription{
		height:      opt.Height,
		width:       opt.Width,
		textPadding: opt.TextPadding,
		title:       title,
		image:       ebiten.NewImage(opt.Width*source.TileSize, opt.Height*source.TileSize),
		font:        opt.Font,
	}

	d.makeImage(source)

	return d, nil
}

func (i *ItemDescription) makeImage(source *gm_layer.Image) {

	var (
		img         *ebiten.Image
		tileCountX  int = i.width / source.TileSize
		tileCountY  int = i.height / source.TileSize
		imageHeight int = i.height
		imageWidth  int = i.width
		paddingX    int
		paddingY    int
		op          *ebiten.DrawImageOptions
	)

	if i.width%source.TileSize != 0 {
		tileCountX++
	}

	if i.height%source.TileSize != 0 {
		tileCountY++
	}

	for y := 0; y < tileCountY; y++ {
		for x := 0; x < tileCountX; x++ {

			if imageHeight < source.TileSize {
				paddingY = source.TileSize - int(math.Max(float64(imageHeight), 2))
			} else {
				paddingY = 0
			}

			if imageWidth < source.TileSize {
				paddingX = source.TileSize - int(math.Max(float64(imageWidth), 2))
			} else {
				paddingX = 0
			}

			if x == 0 && y == 0 {
				img = source.GetTilePadding(0, 0, paddingX, paddingY)
			} else if x == tileCountX-1 && y == 0 {
				img = source.GetTilePadding(2, 0, paddingX, paddingY)
			} else if x == 0 && y == tileCountY-1 {
				img = source.GetTilePadding(0, 2, paddingX, paddingY)
			} else if x == tileCountX-1 && y == tileCountY-1 {
				img = source.GetTilePadding(2, 2, paddingX, paddingY)
			} else if x == 0 {
				img = source.GetTilePadding(0, 1, paddingX, paddingY)
			} else if x == tileCountX-1 {
				img = source.GetTilePadding(2, 1, paddingX, paddingY)
			} else if y == 0 {
				img = source.GetTilePadding(1, 0, paddingX, paddingY)
			} else if y == tileCountY-1 {
				img = source.GetTilePadding(1, 2, paddingX, paddingY)
			} else {
				img = source.GetTilePadding(1, 1, paddingX, paddingY)
			}
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*source.TileSize), float64(y*source.TileSize))
			i.image.DrawImage(img, op)
			imageWidth -= source.TileSize
		}
		imageWidth = i.width
		imageHeight -= source.TileSize
	}

	var (
		titleX, titleY int = i.textPadding, i.textPadding
	)

	// Draw title
	i.font.Print(i.image, i.title, image.Pt(titleX, titleY), nil)
}

func (i *ItemDescription) Draw(x, y float64, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(i.image, op)
}
