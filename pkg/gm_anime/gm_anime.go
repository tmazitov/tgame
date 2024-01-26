package gm_anime

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type AnimeOptions struct {
	TileImage    *gm_layer.Image
	TileLifeTime int
	TileCount    int
	TileSize     int
	TileCoords   [][2]int
}

type Anime struct {
	tileImage    *gm_layer.Image
	tileCounter  int
	tileLifeTime int
	tileCount    int
	tileSize     int
	tileCoords   [][2]int
}

func NewAnime(opt AnimeOptions) *Anime {
	return &Anime{
		tileImage:    opt.TileImage,
		tileCounter:  0,
		tileLifeTime: opt.TileLifeTime,
		tileCoords:   opt.TileCoords,
		tileCount:    opt.TileCount,
		tileSize:     opt.TileSize,
	}
}

func (a *Anime) increment() {
	a.tileCounter++
}

func (a *Anime) GetTile() *ebiten.Image {
	var (
		tileX     int
		tileY     int
		tileIndex int
		tileRect  image.Rectangle
	)

	if a.tileCounter == a.tileLifeTime*a.tileCount {
		a.tileCounter = 0
	}

	defer a.increment()
	tileIndex = a.tileCounter / a.tileLifeTime

	tileX = a.tileCoords[tileIndex][0]
	tileY = a.tileCoords[tileIndex][1]
	tileRect = image.Rect(tileX, tileY, tileX+a.tileSize, tileY+a.tileSize)
	return a.tileImage.Inst.SubImage(tileRect).(*ebiten.Image)
}
