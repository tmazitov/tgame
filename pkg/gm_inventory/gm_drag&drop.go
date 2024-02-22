package gm_inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

func (i *Inventory) HandleDragAndDrop(touches []ebiten.TouchID) {

	var (
		touchedSlot *Slot
	)

	// Find touch on slot
	if i.replaceTouch == nil {
		i.replaceTouch = i.findTouchOnSlot(touches)
	}
	if i.replaceTouch == nil {
		return
	}

	// Update dragged item position
	i.replaceTouch.Update()

	// Check if touch is released
	if i.replaceTouch.IsReleased() {
		touchedSlot, _, _ = i.CheckTouchOnSlot(i.replaceTouch.Position())
		if i.putItemIsPossible(i.replaceTouch.draggingItem, touchedSlot) {
			i.putItem(i.replaceTouch, touchedSlot)
		} else {
			i.putItem(i.replaceTouch, i.replaceTouch.draggingItemOriginSlot)
		}
		i.replaceTouch = nil
	}
}

func (i *Inventory) findTouchOnSlot(touches []ebiten.TouchID) *Touch {

	var (
		touch       *Touch
		touchedSlot *Slot
		relX, relY  float64
	)

	// Mouse click right
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		touch = NewMouseTouch(ebiten.MouseButtonRight)
		touchedSlot, relX, relY = i.CheckTouchOnSlot(touch.Position())
		if touchedSlot != nil && !touchedSlot.IsFree() {
			touch.relX, touch.relY = relX, relY
			return i.takeHalfItem(touch, touchedSlot)
		}
	}

	// Mouse click left
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		touch = NewMouseTouch(ebiten.MouseButtonLeft)
		touchedSlot, relX, relY = i.CheckTouchOnSlot(touch.Position())
		if touchedSlot != nil && !touchedSlot.IsFree() {
			touch.relX, touch.relY = relX, relY
			return i.takeItem(touch, touchedSlot)
		}
	}

	// Touch
	for _, id := range touches {
		touch = NewOriginTouch(id)
		touchedSlot, relX, relY = i.CheckTouchOnSlot(touch.Position())
		if touchedSlot != nil && !touchedSlot.IsFree() {
			touch.relX, touch.relY = relX, relY
			return i.takeItem(touch, touchedSlot)
		}
	}
	return (nil)
}

func (i *Inventory) takeItem(touch *Touch, srcSlot *Slot) *Touch {
	touch.SetDraggingItem(srcSlot.Item)
	touch.SetDraggingItemSlot(srcSlot)
	srcSlot.SetItem(nil)
	return touch
}

func (i *Inventory) takeHalfItem(touch *Touch, srcSlot *Slot) *Touch {

	var (
		item      *gm_item.Item = srcSlot.Item
		itemClone *gm_item.Item = item.Clone(item.GetAmount() / 2)
	)

	if item.GetAmount() == itemClone.GetAmount() {
		touch.SetDraggingItem(item)
		srcSlot.SetItem(nil)
	} else {
		touch.SetDraggingItem(itemClone)
		srcSlot.Item.SetAmount(item.GetAmount() - itemClone.GetAmount())
	}
	touch.SetDraggingItemSlot(srcSlot)
	return touch
}

func (i *Inventory) putItemIsPossible(item *gm_item.Item, destSlot *Slot) bool {

	if destSlot == nil {
		return false
	}

	if destSlot.IsFree() {
		return true
	}

	if item.GetID() != destSlot.Item.GetID() {
		return false
	}

	var (
		destSlotItemStackSize uint = destSlot.Item.GetStackSize()
		destSlotItemAmount    uint = destSlot.Item.GetAmount()
	)

	return destSlotItemAmount < destSlotItemStackSize
}

func (i *Inventory) putItem(touch *Touch, destSlot *Slot) {
	if destSlot.Item == nil {
		destSlot.SetItem(touch.draggingItem)
	} else if destSlot.Item.GetID() == touch.draggingItem.GetID() {
		destSlot.Item.SetAmount(destSlot.Item.GetAmount() + touch.draggingItem.GetAmount())
	}
}
