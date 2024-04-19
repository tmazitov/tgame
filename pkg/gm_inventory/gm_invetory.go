package gm_inventory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type InventoryOpt struct {
	Height        int
	Width         int
	SlotImagePath string
	SlotSize      int
	X             float64
	Font          *gm_font.Font
	Y             float64
}

type Inventory struct {
	slots        []*Slot
	Height       int
	Width        int
	IsVisible    bool
	slotImage    *gm_layer.Image
	replaceTouch *Touch
	hoveredSlot  *Slot
	slotSize     int
	font         *gm_font.Font
	x            float64
	y            float64
}

func NewInventory(opt InventoryOpt) (*Inventory, error) {

	if opt.Height <= 0 || opt.Width <= 0 {
		return nil, ErrZeroSize
	}

	if opt.SlotImagePath == "" {
		return nil, ErrSlotImagePathEmpty
	}

	if opt.SlotSize == 0 {
		return nil, ErrSlotImageSizeZero
	}

	if opt.Font == nil {
		return nil, ErrNilFont
	}

	var (
		slots []*Slot = []*Slot{}
		image *gm_layer.Image
		slotX float64
		slotY float64
		err   error
	)

	if image, err = gm_layer.NewImageByPath(opt.SlotImagePath, opt.SlotSize); err != nil {
		return nil, err
	}

	for row := 0; row < opt.Height; row++ {
		for column := 0; column < opt.Width; column++ {
			slotX = opt.X + float64(column*opt.SlotSize)
			slotY = opt.Y + float64(row*opt.SlotSize)
			slots = append(slots, NewSlot(slotX, slotY, image, nil))
		}
	}

	return &Inventory{
		slots:       slots,
		Height:      opt.Height,
		Width:       opt.Width,
		slotImage:   image,
		IsVisible:   false,
		slotSize:    opt.SlotSize,
		hoveredSlot: nil,
		x:           opt.X,
		y:           opt.Y,
		font:        opt.Font,
	}, nil
}

func (i *Inventory) findFreeSlot(item *gm_item.Item) *Slot {
	for _, slot := range i.slots {
		if slot.IsFree() && slot.ItemCollection == "" {
			return slot
		} else if slot.IsFree() && slot.ItemCollection == item.Collection {
			return slot
		}
	}
	return nil
}

func (i *Inventory) findSlotWithSameItem(item *gm_item.Item) *Slot {

	var (
		slotItem *gm_item.Item
	)

	for _, slot := range i.slots {
		if slot.IsFree() {
			continue
		}
		slotItem = slot.Item
		if slotItem.ID == item.ID && slotItem.Amount < slotItem.MaxStackSize {
			return slot
		}
	}

	return nil
}

func (i *Inventory) PutItemToFreeSlot(item *gm_item.Item) bool {

	var freeSlot *Slot

	freeSlot = i.findSlotWithSameItem(item)

	if freeSlot != nil {
		if freeSlot.Item.Amount+item.Amount <= freeSlot.Item.MaxStackSize {
			freeSlot.Item.Amount += item.Amount
			return true
		} else {
			difference := freeSlot.Item.MaxStackSize - freeSlot.Item.Amount
			freeSlot.Item.Amount = freeSlot.Item.MaxStackSize
			item.Amount -= difference
		}
	}

	if freeSlot = i.findFreeSlot(item); freeSlot == nil {
		return false
	}

	freeSlot.SetItem(item)

	return true
}

func (i *Inventory) CheckTouchOnSlot(cursorX, cursorY int) (*Slot, float64, float64) {
	var (
		slotX          float64
		slotY          float64
		slotSize       float64
		touchX, touchY float64
	)

	slotSize = float64(i.slotSize)
	touchX = float64(cursorX)
	touchY = float64(cursorY)

	for _, slot := range i.slots {
		slotX, slotY = slot.GetPosition()
		if touchX >= slotX && touchX <= slotX+slotSize &&
			touchY >= slotY && touchY <= slotY+slotSize {
			return slot, touchX - slotX, touchY - slotY
		}
	}

	return nil, 0, 0
}

func (i *Inventory) Draw(screen *ebiten.Image) {

	// Draw all slots with inside items
	for _, slot := range i.slots {
		slot.Draw(i.font, screen)
	}

	// Draw dragging item
	if i.replaceTouch != nil {
		i.replaceTouch.draggingItem.Draw(screen, nil)
	}

	// Draw item description
	if i.replaceTouch == nil && i.hoveredSlot != nil && i.hoveredSlot.Item != nil {
		i.hoveredSlot.Item.DrawDescription(screen)
	}
}

func (i *Inventory) AddSlot(x, y float64, opt *SlotOptions) *Slot {
	var slot *Slot = NewSlot(x, y, i.slotImage, opt)
	i.slots = append(i.slots, slot)
	return slot
}
