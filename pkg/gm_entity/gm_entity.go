package gm_entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_inventory"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

type GameEntity interface {
	Draw(screen *ebiten.Image, camera *gm_camera.Camera)
}

type Player interface {
	GameEntity
	StaffHandler(keys []ebiten.Key)
	AttackHandler(keys []ebiten.Key)
	MouseHandler(touches []ebiten.TouchID)
	MovementHandler(keys []ebiten.Key)
	DropItemHandler(touches []ebiten.TouchID) (*gm_item.Item, *gm_geometry.Point, *gm_geometry.Point)
	CollectItemsHandler(items []*gm_item.Item, camera *gm_camera.Camera) []*gm_item.Item
	GetSpeed() *float64
	GetMoveVector(keys []ebiten.Key) (float64, float64)
	GetCollider() *gm_geometry.Collider
	GetPosition() (float64, float64)
	GetInventory() *gm_inventory.Inventory
	GetMoveSidePosition() (float64, float64)
	GetShape() gm_geometry.IRect
}
