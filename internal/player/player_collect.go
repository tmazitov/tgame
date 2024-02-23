package player

import "github.com/tmazitov/tgame.git/pkg/gm_item"

func (p *Player) Collect(item *gm_item.Item) {
	p.inventory.PutItemToFreeSlot(item)
}
