package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Player) MouseHandler(touches []ebiten.TouchID) {
	p.inventory.HandleHoverSlot()
}
