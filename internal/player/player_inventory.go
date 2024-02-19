package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_inventory"
)

type PlayerInventory struct {
	inventory       *gm_inventory.Inventory
	x               float64
	pressedKeyIndex int
	y               float64
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
		SlotImageSize: 33,
	}); err != nil {
		return nil, err
	}

	return &PlayerInventory{
		inventory:       inventory,
		pressedKeyIndex: -1,
		x:               x,
		y:               y,
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

func (pi *PlayerInventory) SetVisible(visible bool) {
	pi.inventory.IsVisible = visible
}

func (pi *PlayerInventory) Draw(screen *ebiten.Image) {
	if pi.inventory.IsVisible {
		pi.inventory.Draw(pi.x, pi.y, screen)
	}
}
