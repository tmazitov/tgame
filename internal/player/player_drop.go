package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func (p *Player) DropItemHandler(touches []ebiten.TouchID) (*gm_item.Item, *gm_geometry.Point, *gm_geometry.Point) {

	var item *gm_item.Item = p.inventory.HandleDragAndDrop(touches)

	if item != nil {
		var (
			lastAction  PlayerAction          = p.lastAction
			shape       gm_geometry.IRect     = p.GetShape()
			shapePoints [4]*gm_geometry.Point = shape.Points()
			posX, posY  float64
			source      *gm_geometry.Point
			target      *gm_geometry.Point
		)

		if lastAction == Right_PlayerAction {
			posX = shapePoints[1].X
			posY = shapePoints[1].Y
		} else {
			posX = shapePoints[0].X
			posY = shapePoints[0].Y
		}

		source = gm_geometry.NewPoint(posX, posY)

		if lastAction == Right_PlayerAction {
			posX += stgs.ItemDropDistance
		} else if lastAction == Left_PlayerAction {
			posX -= stgs.ItemDropDistance
		} else if lastAction == Top_PlayerAction {
			posY -= stgs.ItemDropDistance
		} else if lastAction == Bot_PlayerAction {
			posY += stgs.ItemDropDistance
		}

		target = gm_geometry.NewPoint(posX, posY)

		return item, source, target
	}

	return nil, nil, nil
}
