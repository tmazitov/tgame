package gm_map

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

func (m *Map) Update(touches []ebiten.TouchID, keys []ebiten.Key) error {

	var err error

	if err = m.handlePlayerMove(keys); err != nil {
		return err
	}

	m.handleCollectItem()
	m.handleDropItem(touches)
	m.player.MouseHandler(touches)
	m.player.StaffHandler(keys)
	m.player.AttackHandler(keys)

	for _, item := range m.droppedItems {
		item.Update()
	}

	return nil
}

func (m *Map) playerMayMove(keys []ebiten.Key) bool {
	var (
		playerMoveVectorX float64
		playerMoveVectorY float64
		playerCollider    *gm_geometry.Collider = m.player.GetCollider()
	)

	playerMoveVectorX, playerMoveVectorY = m.player.GetMoveVector(keys)
	for _, obj := range m.objs {
		if obj.IntersectVector(playerCollider, playerMoveVectorX, playerMoveVectorY) {
			return false
		}
	}
	for _, entity := range m.entities {
		if entity == m.player {
			continue
		}

		if entity.IntersectVector(playerCollider, playerMoveVectorX, playerMoveVectorY) {
			return false
		}
	}
	return true
}

// func (m *Map) entityMayMove(entity gm_entity.GameEntity, vectorX, vectorY float64) bool {
// 	for _, obj := range m.objs {
// 		if obj.IntersectVector(entity.GetCollider(), vectorX, vectorY) {
// 			return false
// 		}
// 	}
// 	for _, entity := range m.entities {
// 		if entity == m.player {
// 			continue
// 		}

// 		if entity.IntersectVector(entity.GetCollider(), vectorX, vectorY) {
// 			return false
// 		}
// 	}
// 	return true
// }

func (m *Map) handlePlayerMove(keys []ebiten.Key) error {

	var (
		area gm_camera.CameraArea
		err  error
	)

	if !m.playerMayMove(keys) {
		return nil
	}

	area = m.camera.GetPointArea(m.player.GetPosition())
	_, err = m.camera.MovementHandler(keys, area)
	if err != nil {
		return err
	}
	m.player.MovementHandler(keys)
	return nil
}

func (m *Map) handleDropItem(touches []ebiten.TouchID) {

	var (
		dropItem   *gm_item.Item
		dropTarget *gm_geometry.Point
		dropSource *gm_geometry.Point
	)

	dropItem, dropSource, dropTarget = m.player.DropItemHandler(touches)
	if dropItem == nil || dropTarget == nil || dropSource == nil {
		return
	}

	if dropSource.Y == dropTarget.Y {
		dropSource.X -= float64(m.tileSize / 2)
		dropSource.Y -= float64(m.tileSize)
		dropTarget.X -= float64(m.tileSize / 2)
		dropTarget.Y -= float64(m.tileSize)
	} else if dropSource.Y > dropTarget.Y {
		dropSource.Y -= float64(m.tileSize / 2)
	}

	dropItem.SetPosition(dropSource.X, dropSource.Y)
	dropItem.Drop(dropSource, dropTarget)

	m.AddDropItem(dropItem)
}

func (m *Map) handleCollectItem() {

	var collectedItems []*gm_item.Item = m.player.CollectItemsHandler(m.droppedItems, m.camera)

	for _, item := range collectedItems {
		m.DelDropItem(item)
	}
}
