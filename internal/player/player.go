package player

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Player struct {
	X           int
	Y           int
	Speed       int
	anime       *PlayerAnime
	image       *gm_layer.Image
	actionState PlayerAction
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
		X:           x,
		Y:           y,
		Speed:       stgs.PlayerSpeed,
		image:       gameImage,
		anime:       NewPlayerAnime(gameImage),
		actionState: Idle_PlayerAction,
	}
}

func (p *Player) GetNextTile() *ebiten.Image {
	var anime *gm_anime.Anime = p.anime.GetCurrentAnime(p.actionState)
	if anime == nil {
		return (p.anime.Idle.GetTile())
	}
	return anime.GetTile()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.GetNextTile(), op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player action: %d", p.actionState))
}
