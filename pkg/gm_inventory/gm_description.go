package gm_inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (i *Inventory) HandleHoverSlot() {
	var (
		cursorX, cursorY int
		slot             *Slot
	)

	cursorX, cursorY = ebiten.CursorPosition()
	slot, _, _ = i.CheckTouchOnSlot(cursorX, cursorY)
	if slot == nil {
		i.hoveredSlot = nil
		return
	}
	i.hoveredSlot = slot
}

func (i *Inventory) HandleDrawItemDescription(screen *ebiten.Image) {
	if i.hoveredSlot == nil || i.hoveredSlot.IsFree() {
		return
	}

	i.hoveredSlot.Item.DrawDescription(screen)
}
