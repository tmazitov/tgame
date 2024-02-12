package gm_layer

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func CalcFramePosition(image *Image, tile int) (int, int) {
	return (tile % image.TileXCount) * stgs.TileSize, (tile / image.TileXCount) * stgs.TileSize
}

func CalcTilePosition(index int, rowIndex int) (float64, float64) {
	return float64(index * stgs.TileSize),
		float64(rowIndex * stgs.TileSize)
}

func (l *Layer) DrawRow(screen *ebiten.Image, row []int, rowIndex int, b LayerBorder) {

	var (
		dx        int = int(math.Round(b.X)) % l.TileSize
		dy        int = int(math.Round(b.Y)) % l.TileSize
		posX      float64
		posY      float64
		op        *ebiten.DrawImageOptions
		tileSize  int = stgs.TileSize
		tileFrame image.Rectangle
		frameX    int
		width     int = b.Width / l.TileSize
		frameY    int
	)

	for tileIndex, tile := range row {

		// If layer tile is equal to 0, transparent block
		if tile == 0 {
			continue
		}

		// Calculate tile position
		posX, posY = CalcTilePosition(tileIndex, rowIndex)
		if tileIndex != 0 {
			posX -= float64(dx)
		}
		if rowIndex != 0 {
			posY -= float64(dy)
		}
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(posX, posY)

		// Calculate tile frame position in the image
		frameX, frameY = CalcFramePosition(l.image, tile-1)

		// First row is very important thing to make smooth movement on Y axis
		if rowIndex == 0 {
			tileFrame = l.firstRowTile(dx, dy, tileIndex, width, frameX, frameY)
			screen.DrawImage(l.image.Inst.SubImage(tileFrame).(*ebiten.Image), op)
			continue
		}

		// Handle first and last row elements to make smooth movement on X axis
		if tileIndex == 0 && dx != 0 {
			tileFrame = image.Rect(frameX+dx, frameY, frameX+tileSize, frameY+tileSize)
		} else if tileIndex == len(row)-1 && dx != 0 {
			tileFrame = image.Rect(frameX, frameY, frameX+dx, frameY+tileSize)
		} else {
			tileFrame = image.Rect(frameX, frameY, frameX+tileSize, frameY+tileSize)
		}

		screen.DrawImage(l.image.Inst.SubImage(tileFrame).(*ebiten.Image), op)
	}
}

func (l *Layer) Draw(screen *ebiten.Image, b LayerBorder) {
	for index, tileRow := range l.GetTileRows(b) {
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
