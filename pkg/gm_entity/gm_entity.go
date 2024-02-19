package gm_entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_inventory"
)

type GameEntity interface {
	Draw(screen *ebiten.Image)
}

type Player interface {
	GameEntity
	Draw(screen *ebiten.Image)
	StaffHandler(keys []ebiten.Key)
	AttackHandler(keys []ebiten.Key)
	MovementHandler(keys []ebiten.Key, stay bool)
	GetSpeed() *float64
	GetMoveVector(keys []ebiten.Key) (float64, float64)
	GetCollider() *gm_geometry.Collider
	GetPosition() (float64, float64)
	GetInventory() *gm_inventory.Inventory
	GetMoveSidePosition() (float64, float64)
}
