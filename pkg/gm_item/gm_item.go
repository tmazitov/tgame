package gm_item

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Item struct {
	ID           uint
	Name         string
	MaxStackSize uint
	Amount       uint
	image        *gm_layer.Image
	x            float64
	y            float64
	isMoving     bool
}

type ItemOptions struct {
	MaxStackSize uint
	Amount       uint
	X            float64
	Y            float64
}

func NewItem(id uint, name string, imagePath string, opt ItemOptions) (*Item, error) {

	var (
		image *gm_layer.Image
		err   error
	)

	if imagePath == "" || name == "" {
		return nil, ErrInvalidParams
	}

	if image, err = gm_layer.NewImageByPath(imagePath); err != nil {
		return nil, err
	}
	if opt.Amount == 0 {
		opt.Amount = 1
	}

	return &Item{
		ID:           id,
		Name:         name,
		MaxStackSize: opt.MaxStackSize,
		image:        image,
		x:            opt.X,
		y:            opt.Y,
		isMoving:     false,
		Amount:       opt.Amount,
	}, nil
}

func (i *Item) Clone(amount uint) *Item {

	if amount == 0 {
		amount = 1
	}

	return &Item{
		ID:           i.ID,
		Name:         i.Name,
		MaxStackSize: i.MaxStackSize,
		Amount:       amount,
		image:        i.image,
		x:            i.x,
		y:            i.y,
		isMoving:     false,
	}
}

func (i *Item) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(i.x, i.y)
	screen.DrawImage(i.image.Inst, op)
}

func (i *Item) GetID() uint {
	return i.ID
}

func (i *Item) GetIsMoving() bool {
	return i.isMoving
}

func (i *Item) SetIsMoving(value bool) {
	i.isMoving = value
}

func (i *Item) SetPosition(x, y float64) {
	i.x = x
	i.y = y
}

func (i *Item) GetPosition() (float64, float64) {
	return i.x, i.y
}

func (i *Item) GetStackSize() uint {
	return i.MaxStackSize
}

func (i *Item) GetAmount() uint {
	return i.Amount
}

func (i *Item) SetAmount(value uint) {
	i.Amount = value
}
