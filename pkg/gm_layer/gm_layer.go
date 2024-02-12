package gm_layer

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func (l *Layer) GetValue(b LayerBorder) [][]int {
	var (
		initX         float64 = b.X / float64(l.TileSize)
		initY         float64 = b.Y / float64(l.TileSize)
		initWidth     float64 = float64(b.Width/l.TileSize) + 1
		initHeight    float64 = float64(b.Height / l.TileSize)
		layerHeight   float64 = float64(l.raw.height)
		layerWidth    float64 = float64(l.raw.width)
		limitValues   [][]int = [][]int{}
		limitRow      []int   = []int{}
		limitXCounter float64 = 0
		limitYCounter float64 = 0
	)

	if initHeight+initY > float64(l.raw.height) || initWidth+initX > float64(l.raw.width) {
		return [][]int{}
	}

	if int(b.Y)%l.TileSize != 0 {
		initHeight += 1
	}

	for _, tile := range l.raw.tiles {
		if int(b.X)%l.TileSize != 0 &&
			limitXCounter >= initX && limitXCounter < initX+initWidth &&
			limitYCounter >= initY && limitYCounter <= initY+initHeight {
			limitRow = append(limitRow, tile)
		}
		if int(b.X)%l.TileSize == 0 &&
			limitXCounter >= initX && limitXCounter <= initX+initWidth &&
			limitYCounter >= initY && limitYCounter <= initY+initHeight {
			limitRow = append(limitRow, tile)
		}
		if limitXCounter == layerWidth {
			if len(limitRow) != 0 {
				limitValues = append(limitValues, limitRow)
				limitRow = []int{}
			}
			limitXCounter = 0
			limitYCounter += 1
		} else {
			limitXCounter += 1
		}
		if (limitYCounter > initHeight+initY) || limitYCounter == layerHeight {
			break
		}
	}

	return limitValues
}

func CalcFramePosition(image *Image, tile int) (int, int) {
	return (tile % image.TileXCount) * stgs.TileSize, (tile / image.TileXCount) * stgs.TileSize
}

func CalcTilePosition(index int, rowIndex int) (float64, float64) {
	return float64(index * stgs.TileSize),
		float64(rowIndex * stgs.TileSize)
}

func (l *Layer) DrawRow(screen *ebiten.Image, row []int, rowIndex int, b LayerBorder) {

	var (
		dx        int = int(b.X) % l.TileSize
		dy        int = int(b.Y) % l.TileSize
		posX      float64
		posY      float64
		op        *ebiten.DrawImageOptions
		tileSize  int = stgs.TileSize
		tileFrame image.Rectangle
		frameX    int
		width     int = b.Width / l.TileSize
		height    int = b.Height / l.TileSize
		frameY    int
	)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d %d\n", dx, dy), 5, 20)

	for tileIndex, tile := range row {
		if tile == 0 {
			continue
		}
		posX, posY = CalcTilePosition(tileIndex, rowIndex)

		if tileIndex != 0 {
			posX -= float64(dx)
		}
		if rowIndex != 0 {
			posY -= float64(dy)
		}

		if dx == 0 {
			posX -= float64(tileSize)
		}
		if dy == 0 && rowIndex != 0 {
			posY -= float64(tileSize)
		}

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(posX, posY)

		// Smooth Y axis movement

		frameX, frameY = CalcFramePosition(l.image, tile-1)
		if rowIndex == 0 {
			if dy == 0 {
				return
			}
			tileFrame = l.firstRowTile(dx, dy, tileIndex, width, frameX, frameY)
			screen.DrawImage(l.image.Inst.SubImage(tileFrame).(*ebiten.Image), op)
			continue
		}
		// Smooth X axis movement

		if rowIndex != 0 && rowIndex != height-1 {
			if tileIndex == 0 {
				if dx != 0 {
					tileFrame = image.Rect(frameX+dx, frameY, frameX+tileSize, frameY+tileSize)
				} else {
					tileFrame = image.Rect(frameX, frameY, frameX+tileSize, frameY+tileSize)
				}
			} else if tileIndex == len(row)-1 {
				if dx != 0 {
					tileFrame = image.Rect(frameX, frameY, frameX+dx, frameY+tileSize)
				} else {
					tileFrame = image.Rect(frameX, frameY, frameX+tileSize, frameY+tileSize)
				}
			} else {
				tileFrame = image.Rect(frameX, frameY, frameX+tileSize, frameY+tileSize)
			}
		} else {
			tileFrame = image.Rect(frameX, frameY, frameX+tileSize, frameY+tileSize)
		}

		screen.DrawImage(l.image.Inst.SubImage(tileFrame).(*ebiten.Image), op)
	}
}

func (l *Layer) Draw(screen *ebiten.Image, b LayerBorder) {
	for index, tileRow := range l.GetValue(b) {
		l.DrawRow(screen, tileRow, index, b)
	}
}

func (l *Layer) firstRowTile(dx, dy int, index int, width int, frameX, frameY int) image.Rectangle {

	var (
		tileFrame image.Rectangle
	)

	// First row element

	if index == 0 {
		if dx != 0 && dy != 0 {
			tileFrame = image.Rect(frameX+dx, frameY+dy, frameX+l.TileSize, frameY+l.TileSize)
		} else if dx != 0 {
			tileFrame = image.Rect(frameX+dx, frameY, frameX+l.TileSize, frameY+l.TileSize)
		} else if dy != 0 {
			tileFrame = image.Rect(frameX, frameY+dy, frameX+l.TileSize, frameY+l.TileSize)
		} else {
			tileFrame = image.Rect(frameX, frameY, frameX+l.TileSize, frameY+l.TileSize)
		}
		return tileFrame
	}

	// Last row element

	if index == width {
		if dx != 0 && dy != 0 {
			tileFrame = image.Rect(frameX, frameY+dy, frameX+dx, frameY+l.TileSize)
		} else if dx != 0 {
			tileFrame = image.Rect(frameX, frameY, frameX+dx, frameY+l.TileSize)
		} else if dy != 0 {
			tileFrame = image.Rect(frameX, frameY+dy, frameX+l.TileSize, frameY+l.TileSize)
		} else {
			tileFrame = image.Rect(frameX, frameY, frameX+l.TileSize, frameY+l.TileSize)
		}
		return tileFrame
	}

	if dy != 0 {
		tileFrame = image.Rect(frameX, frameY+dy, frameX+l.TileSize, frameY+l.TileSize)
	} else {
		tileFrame = image.Rect(frameX, frameY, frameX+l.TileSize, frameY+l.TileSize)
	}

	return tileFrame
}
