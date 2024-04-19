package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/internal/items"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_inventory"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

type PlayerInventory struct {
	inventory       *gm_inventory.Inventory
	weaponSlot      *gm_inventory.Slot
	pressedKeyIndex int
}

func NewPlayerInventory(x, y float64, font *gm_font.Font) (*PlayerInventory, error) {

	var (
		inventory  *gm_inventory.Inventory
		slotSize   float64 = 33
		weaponSlot *gm_inventory.Slot
		err        error
	)

	if inventory, err = gm_inventory.NewInventory(gm_inventory.InventoryOpt{
		Height:        3,
		Width:         6,
		SlotImagePath: "assets/textures/inventory_slot_3.png",
		SlotSize:      int(slotSize),
		X:             x,
		Y:             y,
		Font:          font,
	}); err != nil {
		return nil, err
	}

	weaponSlot = inventory.AddSlot(x+slotSize*1.5, y-slotSize*2, &gm_inventory.SlotOptions{
		ItemCollection: items.WeaponCollection,
	})

	return &PlayerInventory{
		inventory:       inventory,
		weaponSlot:      weaponSlot,
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
