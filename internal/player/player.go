package player

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type PlayerImagesPaths struct {
	Tiles  string
	Shadow string
}

type PlayerImages struct {
	Tiles  *gm_layer.Image
	Shadow *gm_layer.Image
}

type Player struct {
	X           float32
	Y           float32
	Speed       float32
	anime       *PlayerAnime
	images      *PlayerImages
	lastAction  PlayerAction
	actionState PlayerAction
	attack      *PlayerAttackSystem
}

func NewPlayer(x, y float32, imagesPaths PlayerImagesPaths) (*Player, error) {

	var (
		err         error
		playerAnime *PlayerAnime
		pl          *Player
	)

	pl = &Player{
		X:           x,
		Y:           y,
		Speed:       stgs.PlayerSpeed,
		images:      &PlayerImages{},
		anime:       playerAnime,
		actionState: Idle_PlayerAction,
		attack:      nil,
	}

	if pl.images.Tiles, err = gm_layer.NewImageByPath(imagesPaths.Tiles); err != nil {
		return nil, err
	}

	if pl.images.Shadow, err = gm_layer.NewImageByPath(imagesPaths.Shadow); err != nil {
		return nil, err
	}

	if pl.anime = NewPlayerAnime(pl.images.Tiles); err != nil {
		return nil, err
	}

	if pl.attack, err = NewPlayerAttackSystem(&pl.X, &pl.Y, &pl.lastAction); pl.attack == nil {
		return nil, err
	}

	if stgs.IsDebug {
		log.Println("Player create\t\tsuccess")
	}

	return pl, nil
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

func (p *Player) drawShadow(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.images.Shadow.Inst, op)
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	var tile *ebiten.Image = p.GetNextTile()
	if p.actionState == Left_PlayerAction || p.lastAction == Left_PlayerAction {
		tile = FlipVertical(tile)
	}

	screen.DrawImage(tile, op)
	p.drawShadow(screen)

	for _, fireball := range p.attack.GetFireballs() {
		fireball.Move()
		fireball.Draw(screen)
	}
}
