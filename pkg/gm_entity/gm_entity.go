package gm_entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
)

type GameEntity interface {
	Draw(screen *ebiten.Image)
}

type Player interface {
	GameEntity
	Draw(screen *ebiten.Image)
	AttackHandler(keys []ebiten.Key)
	MovementHandler(keys []ebiten.Key, stay bool)
	GetSpeed() *float64
	GetMoveVector(keys []ebiten.Key) (float64, float64)
	GetCollider() *gm_geometry.Collider
	GetPosition() (float64, float64)
	GetMoveSidePosition() (float64, float64)
}
