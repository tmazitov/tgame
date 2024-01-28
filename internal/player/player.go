package player

import (
	"fmt"
	"image"
	"log"
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
	lastAction  PlayerAction
	actionState PlayerAction
}

func NewPlayer(x, y int, tilesImagePath string) *Player {

	var (
		file        *os.File
		err         error
		img         image.Image
		playerImg   *gm_layer.Image
		playerAnime *PlayerAnime
	)

	if file, err = os.Open(tilesImagePath); err != nil {
		return nil
	}
	defer file.Close()

	if img, _, err = image.Decode(file); err != nil {
		return nil
	}

	if playerImg = gm_layer.NewImage(ebiten.NewImageFromImage(img)); err != nil {
		return nil
	}

	if playerAnime = NewPlayerAnime(playerImg); err != nil {
		return nil
	}

	if stgs.IsDebug {
		log.Println("Player create\t\tsuccess")
	}

	return &Player{
		X:           x,
		Y:           y,
		Speed:       stgs.PlayerSpeed,
		image:       playerImg,
		anime:       playerAnime,
		actionState: Idle_PlayerAction,
	}
}

func (p *Player) GetNextTile() *ebiten.Image {
	var anime *gm_anime.Anime = p.anime.GetCurrentAnime(p.actionState, p.lastAction)
	if anime == nil {
		return (p.anime.IdleBot.GetTile())
	}
	return anime.GetTile()
}

func FlipVertical(source *ebiten.Image) *ebiten.Image {
	result := ebiten.NewImage(source.Bounds().Dx(), source.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(result.Bounds().Dx()), 0)
	result.DrawImage(source, op)
	return result
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	var tile *ebiten.Image = p.GetNextTile()
	if p.actionState == Left_PlayerAction || p.lastAction == Left_PlayerAction {
		tile = FlipVertical(tile)
	}
	screen.DrawImage(tile, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player coords: %d %d", p.lastAction, p.actionState))
}
