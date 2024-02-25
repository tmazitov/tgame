package gm_font

import (
	"image"
	"image/color"
	"os"

	"github.com/Frabjous-Studios/ingenten"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Font struct {
	face *ingenten.PixelFont
}

type FontOptions struct {
	Size    float64      // Size is the font size in points
	DPI     float64      // DPI is the dots per inch resolution
	Hinting font.Hinting // Hinting selects how to quantize a vector font's glyph nodes
}

func NewFont(fontPath string) (*Font, error) {

	// var (
	// 	file         *os.File
	// 	fileBytes    []byte
	// 	openTypeFont *sfnt.Font
	// 	font         font.Face
	// 	err          error
	// )

	// if file, err = os.Open(fontPath); err != nil {
	// 	return nil, err
	// }

	// defer file.Close()

	// if fileBytes, err = ioutil.ReadAll(file); err != nil {
	// 	return nil, err
	// }

	// if openTypeFont, err = opentype.Parse(fileBytes); err != nil {
	// 	return nil, err
	// }

	// if font, err = opentype.NewFace(openTypeFont, &opentype.FaceOptions{
	// 	Size:    opt.Size,
	// 	DPI:     opt.DPI,
	// 	Hinting: opt.Hinting,
	// }); err != nil {
	// 	return nil, err
	// }

	font, err := ingenten.LoadPixelFont(fontPath, os.DirFS("."))
	if err != nil {
		return nil, err
	}

	return &Font{
		face: font,
	}, nil
}

type PrintOptions struct {
	X, Y  int
	Color color.Color
}

func (f *Font) Print(img *ebiten.Image, message string, pos image.Point, op *ebiten.DrawImageOptions) {

	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}

	shadowPos := image.Pt(pos.X+1, pos.Y)
	shadowOp := &ebiten.DrawImageOptions{}
	shadowOp.ColorScale.Scale(0, 0, 0, 0.5)
	f.face.PrintOpts(img, shadowPos, message, shadowOp)
	f.face.PrintOpts(img, pos, message, op)
}
