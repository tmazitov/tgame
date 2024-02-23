package player

import (
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func (p *Player) Collect(item *gm_item.Item) bool {
	return p.inventory.PutItemToFreeSlot(item)
}

func (p *Player) CollectItemsHandler(items []*gm_item.Item) []*gm_item.Item {

	var (
		playerPos gm_geometry.Point = gm_geometry.Point{
			X: p.X,
			Y: p.Y,
		}
		itemPos        gm_geometry.Point
		collectedItems []*gm_item.Item
	)

	for _, item := range items {
		itemPos = gm_geometry.Point{
			X: item.X - float64(item.Size()/2),
			Y: item.Y - float64(item.Size()/2),
		}
		if gm_geometry.Length(itemPos, playerPos) <= stgs.ItemCollectionDistance {
			if p.Collect(item) {
				collectedItems = append(collectedItems, item)
			}
		}
	}
	return collectedItems
}
