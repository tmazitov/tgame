package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

func (p *Player) MouseHandler(touches []ebiten.TouchID) {
	p.inventory.HandleHoverSlot()
}

func (p *Player) DropItemHandler(touches []ebiten.TouchID) *gm_item.Item {
	return p.inventory.HandleDragAndDrop(touches)
}
