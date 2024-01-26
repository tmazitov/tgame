package gm_entity

import "github.com/hajimehoshi/ebiten/v2"

type GameEntity interface {
	Draw(screen *ebiten.Image)
}
