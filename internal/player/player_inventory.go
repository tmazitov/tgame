package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_inventory"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

type PlayerInventory struct {
	inventory       *gm_inventory.Inventory
	pressedKeyIndex int
}

func NewPlayerInventory(x, y float64) (*PlayerInventory, error) {

	var (
		inventory *gm_inventory.Inventory
		err       error
	)

	if inventory, err = gm_inventory.NewInventory(gm_inventory.InventoryOpt{
		Height:        4,
		Width:         6,
		SlotImagePath: "assets/textures/inventory_slot_3.png",
		SlotSize:      33,
		X:             x,
		Y:             y,
	}); err != nil {
		return nil, err
	}

	return &PlayerInventory{
		inventory:       inventory,
		pressedKeyIndex: -1,
	}, nil
}

func (pi *PlayerInventory) HandleToggle(keys []ebiten.Key) {
	var (
		toggleKeyIndex int = -1
	)

	for keyIndex, key := range keys {
		if key == ebiten.KeyI {
			toggleKeyIndex = keyIndex
			break
		}
	}
	if toggleKeyIndex != -1 && pi.pressedKeyIndex == -1 {
		pi.pressedKeyIndex = toggleKeyIndex
		pi.SetVisible(!pi.inventory.IsVisible)
	} else if toggleKeyIndex == -1 && pi.pressedKeyIndex != -1 {
		pi.pressedKeyIndex = -1
	}
}

func (pi *PlayerInventory) HandleDragAndDrop(touches []ebiten.TouchID) *gm_item.Item {
	if !pi.inventory.IsVisible {
		return nil
	}
	return pi.inventory.HandleDragAndDrop(touches)
}

func (pi *PlayerInventory) HandleHoverSlot() {
	if !pi.inventory.IsVisible {
		return
	}
	pi.inventory.HandleHoverSlot()
}

func (pi *PlayerInventory) SetVisible(visible bool) {
	pi.inventory.IsVisible = visible
}

func (pi *PlayerInventory) Draw(screen *ebiten.Image) {
	if pi.inventory.IsVisible {
		pi.inventory.Draw(screen)
	}
}

func (pi *PlayerInventory) PutItemToFreeSlot(item *gm_item.Item) bool {
	return pi.inventory.PutItemToFreeSlot(item)
}

func (pi *PlayerInventory) PutItem(item *gm_item.Item, x, y uint) {
	pi.inventory.PutItem(item, x, y)
}
