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
	return nil
}

func (m *Map) playerMayMove(keys []ebiten.Key) bool {
	var (
		playerMoveVectorX float64
		playerMoveVectorY float64
	)

	playerMoveVectorX, playerMoveVectorY = m.player.GetMoveVector(keys)
	for _, obj := range m.objs {
		if obj.IntersectVector(m.player.GetCollider(), playerMoveVectorX, playerMoveVectorY) {
			return false
		}
	}
	return true
}

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
		dropItem  *gm_item.Item
		dropPoint *gm_geometry.Point
	)

	dropItem, dropPoint = m.player.DropItemHandler(touches)
	if dropItem == nil || dropPoint == nil {
		return
	}
	dropItem.Drop(dropPoint.X, dropPoint.Y)

	m.AddDropItem(dropItem)

}

func (m *Map) handleCollectItem() {

	var collectedItems []*gm_item.Item = m.player.CollectItemsHandler(m.droppedItems, m.camera)

	for _, item := range collectedItems {
		m.DelDropItem(item)
	}
}
