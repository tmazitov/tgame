package player

import (
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func (p *Player) Collect(item *gm_item.Item) bool {
	return p.inventory.PutItemToFreeSlot(item)
}

func (p *Player) CollectItemsHandler(items []*gm_item.Item, camera *gm_camera.Camera) []*gm_item.Item {

	var (
		playerPos      gm_geometry.Point
		playerRelX     float64
		playerRelY     float64
		isInCamera     bool
		itemPos        gm_geometry.Point
		collectedItems []*gm_item.Item
		relX, relY     float64
	)

	playerRelX, playerRelY, isInCamera = camera.GetRelativeCoords(p.X, p.Y)
	if !isInCamera {
		return nil
	}
	playerPos = gm_geometry.Point{
		X: playerRelX - float64(p.images.Tiles.TileSize/2),
		Y: playerRelY - float64(p.images.Tiles.TileSize/2),
	}
	for _, item := range items {
		relX, relY, isInCamera = camera.GetRelativeCoords(item.X, item.Y)
		// fmt.Printf("relX: %v, relY: %v, isInCamera: %v\n", relX, relY, isInCamera)
		if !isInCamera {
			continue
		}

		itemPos = gm_geometry.Point{
			X: relX - float64(item.Size()/2),
			Y: relY - float64(item.Size()/2),
		}
		if gm_geometry.Length(itemPos, playerPos) <= stgs.ItemCollectionDistance {
			if p.Collect(item) {
				collectedItems = append(collectedItems, item)
			}
		}
	}
	return collectedItems
}
