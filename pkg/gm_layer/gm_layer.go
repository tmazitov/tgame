package gm_layer

import (
	"math"
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

	if image, err = NewImageByPath(imagePath, tileSize); err != nil {
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

	if image, err = NewImageByPath(imagePath, tileSize); err != nil {
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

func (l *Layer) GetTileRows(b LayerBorder) [][]int {
	var (
		initX         int     = int(math.Round(b.X)) / l.TileSize
		initY         int     = int(math.Round(b.Y)) / l.TileSize
		initWidth     int     = b.Width / l.TileSize
		initHeight    int     = b.Height / l.TileSize
		layerWidth    int     = l.raw.width - 1
		limitValues   [][]int = [][]int{}
		limitRow      []int   = []int{}
		limitXCounter int     = 0
		limitYCounter int     = 0
	)

	if initHeight+initY > l.raw.height || initWidth+initX > l.raw.width {
		return [][]int{}
	}

	for _, tile := range l.raw.tiles {
		if limitXCounter >= initX && limitXCounter <= initX+initWidth &&
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
		if limitYCounter > initY+initHeight {
			break
		}
	}
	return limitValues
}
