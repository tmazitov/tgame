package player

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Player struct {
	X           float32
	Y           float32
	Speed       float32
	anime       *PlayerAnime
	image       *gm_layer.Image
	lastAction  PlayerAction
	actionState PlayerAction
	attack      *PlayerAttackSystem
}

func NewPlayer(x, y float32, tilesImagePath string) *Player {

	var (
		file        *os.File
		err         error
		img         image.Image
		playerImg   *gm_layer.Image
		playerAnime *PlayerAnime
		pl          *Player
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

	pl = &Player{
		X:           x,
		Y:           y,
		Speed:       stgs.PlayerSpeed,
		image:       playerImg,
		anime:       playerAnime,
		actionState: Idle_PlayerAction,
		attack:      nil,
	}

	pl.attack = NewPlayerAttackSystem(&pl.X, &pl.Y, &pl.lastAction)
	if pl.attack == nil {
		return nil
	}
	if stgs.IsDebug {
		log.Println("Player create\t\tsuccess")
	}

	return pl
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

	for _, fireball := range p.attack.GetFireballs() {
		fireball.Move()
		fireball.Draw(screen)
	}
}
