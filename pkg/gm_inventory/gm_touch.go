package gm_inventory

import "github.com/tmazitov/tgame.git/pkg/gm_item"

type Touch struct {
	initX    int
	initY    int
	currentX int
	currentY int
	relX     float64
	relY     float64
	source   TouchSource

	released bool

	// draggingItem represents a object (sprite in this case)
	// that is being dragged.
	draggingItemOriginSlot *Slot
	draggingItem           *gm_item.Item
}

func NewTouch(source TouchSource) *Touch {
	var (
		x int
		y int
	)

	x, y = source.Position()
	return &Touch{
		initX:    x,
		initY:    y,
		currentX: x,
		currentY: y,
		source:   source,
		released: false,
	}
}

func (t *Touch) Update() {
	if t.released {
		return
	}
	if t.source.IsJustReleased() {
		t.released = true
		t.draggingItem.SetIsMoving(false)
		return
	}
	x, y := t.source.Position()
	t.currentX = x
	t.currentY = y
	t.draggingItem.SetPosition(float64(x)-t.relX, float64(y)-t.relY)
}

func (t *Touch) IsReleased() bool {
	return t.released
}

func (t *Touch) Position() (int, int) {
	return t.currentX, t.currentY
}

func (t *Touch) PositionDiff() (int, int) {
	dx := t.currentX - t.initX
	dy := t.currentY - t.initY
	return dx, dy
}

func (t *Touch) SetDraggingItemSlot(slot *Slot) {
	t.draggingItemOriginSlot = slot
}
func (t *Touch) SetDraggingItem(item *gm_item.Item) {
	t.draggingItem = item
	item.SetIsMoving(true)
}
