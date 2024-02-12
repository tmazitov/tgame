package gm_entity

import "github.com/hajimehoshi/ebiten/v2"

type GameEntity interface {
	Draw(screen *ebiten.Image)
}

type Player interface {
	GameEntity
	Draw(screen *ebiten.Image)
	AttackHandler(keys []ebiten.Key)
	MovementHandler(keys []ebiten.Key, stay bool)
	GetSpeed() *float64
	GetPosition() (float64, float64)
	GetMoveSidePosition() (float64, float64)
}
