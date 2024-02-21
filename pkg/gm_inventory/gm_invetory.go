package gm_inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type InventoryOpt struct {
	Height        int
	Width         int
	SlotImagePath string
	SlotImageSize int
	X             float64
	Y             float64
}

type Inventory struct {
	slots         [][]*Slot
	Height        int
	Width         int
	IsVisible     bool
	slotImage     *gm_layer.Image
	replaceTouch  *Touch
	slotImageSize int
	x             float64
	y             float64
}

func NewInventory(opt InventoryOpt) (*Inventory, error) {

	if opt.Height <= 0 || opt.Width <= 0 {
		return nil, ErrZeroSize
	}

	if opt.SlotImagePath == "" {
		return nil, ErrSlotImagePathEmpty
	}

	if opt.SlotImageSize == 0 {
		return nil, ErrSlotImageSizeZero
	}

	var (
		slots [][]*Slot = [][]*Slot{}
		image *gm_layer.Image
		err   error
	)

	if image, err = gm_layer.NewImageByPath(opt.SlotImagePath); err != nil {
		return nil, err
	}

	for i := 0; i < opt.Height; i++ {
		slots = append(slots, []*Slot{})
		for j := 0; j < opt.Width; j++ {
			slots[i] = append(slots[i], NewSlot(image))
		}
	}

	return &Inventory{
		slots:         slots,
		Height:        opt.Height,
		Width:         opt.Width,
		slotImage:     image,
		IsVisible:     false,
		slotImageSize: opt.SlotImageSize,
		x:             opt.X,
		y:             opt.Y,
	}, nil
}

func (i *Inventory) findFreeSlot() *Slot {

	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			if i.slots[y][x].IsFree() {
				return i.slots[y][x]
			}
		}
	}
	return nil
}

func (i *Inventory) PutItemToFreeSlot(item *gm_item.Item) bool {

	var freeSlot *Slot

	if freeSlot = i.findFreeSlot(); freeSlot == nil {
		return false
	}

	freeSlot.SetItem(item)

	return true
}

func (i *Inventory) PutItem(item *gm_item.Item, x, y uint) bool {
	if !i.slots[y][x].IsFree() {
		return false
	}

	i.slots[y][x].SetItem(item)

	return true
}

func (i *Inventory) CheckTouchOnSlot(touch *Touch) (*Slot, float64, float64) {
	var (
		slot                 *Slot
		slotX                float64
		slotY                float64
		slotSize             float64
		touchX, touchY       float64
		touchXint, touchYint int
	)

	slotSize = float64(i.slotImageSize)
	touchXint, touchYint = touch.Position()
	touchX = float64(touchXint)
	touchY = float64(touchYint)
	for row := 0; row < i.Height; row++ {
		for column := 0; column < i.Width; column++ {
			slot = i.slots[row][column]
			slotX = i.x + float64(column*i.slotImageSize)
			slotY = i.y + float64(row*i.slotImageSize)
			if touchX >= slotX && touchX <= slotX+slotSize &&
				touchY >= slotY && touchY <= slotY+slotSize {
				return slot, touchX - slotX, touchY - slotY
			}
		}
	}
	return nil, 0, 0
}

func (i *Inventory) Draw(screen *ebiten.Image) {

	var (
		slot  *Slot
		slotX float64
		slotY float64
	)

	for row := 0; row < i.Height; row++ {
		for column := 0; column < i.Width; column++ {
			slot = i.slots[row][column]
			slotX = i.x + float64(column*i.slotImageSize)
			slotY = i.y + float64(row*i.slotImageSize)
			slot.Draw(slotX, slotY, screen)
		}
	}

	if i.replaceTouch != nil {
		i.replaceTouch.draggingItem.Draw(screen)
	}
}
