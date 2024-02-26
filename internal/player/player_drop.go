package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func (p *Player) DropItemHandler(touches []ebiten.TouchID) (*gm_item.Item, *gm_geometry.Point) {

	var item *gm_item.Item = p.inventory.HandleDragAndDrop(touches)

	if item != nil {
		var (
			lastAction PlayerAction = p.lastAction
			posX, posY float64      = p.X, p.Y
		)

		if lastAction == Right_PlayerAction {
			posX += stgs.ItemDropDistance
		} else if lastAction == Left_PlayerAction {
			posX -= stgs.ItemDropDistance
		} else if lastAction == Top_PlayerAction {
			posY -= stgs.ItemDropDistance
		} else if lastAction == Bot_PlayerAction {
			posY += stgs.ItemDropDistance
		}

		posX += float64(stgs.TileSize / 2)
		posY += float64(stgs.TileSize / 2)

		// item.Drop(posX, posY)
		return item, gm_geometry.NewPoint(posX, posY)
	}

	return nil, nil
}
