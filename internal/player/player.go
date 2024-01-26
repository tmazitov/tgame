package player

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Player struct {
	X     int
	Y     int
	anime *PlayerAnime
	image *gm_layer.Image
}

func NewPlayer(x, y int, tilesImagePath string) *Player {

	file, err := os.Open(tilesImagePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil
	}

	gameImage := gm_layer.NewImage(ebiten.NewImageFromImage(img))

	return &Player{
		X:     x,
		Y:     y,
		image: gameImage,
		anime: NewPlayerAnime(gameImage),
	}
}

func (p *Player) GetNextTile() *ebiten.Image {
	return p.anime.Idle.GetTile()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.GetNextTile(), op)
}
