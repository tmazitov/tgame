package gm_item

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Item struct {
	ID           uint
	Name         string
	MaxStackSize int
	image        *gm_layer.Image
	x            float64
	y            float64
	isMoving     bool
}

type ItemOptions struct {
	MaxStackSize int
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

	return &Item{
		ID:           id,
		Name:         name,
		MaxStackSize: opt.MaxStackSize,
		image:        image,
		x:            opt.X,
		y:            opt.Y,
		isMoving:     false,
	}, nil
}

func (i *Item) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(i.x, i.y)
	screen.DrawImage(i.image.Inst, op)
}

func (i *Item) GetIsMoving() bool {
	return i.isMoving
}

func (i *Item) SetPosition(x, y float64) {
	i.x = x
	i.y = y
}

func (i *Item) GetPosition() (float64, float64) {
	return i.x, i.y
}

func (i *Item) Drag() {
	i.isMoving = true
}

func (i *Item) Move(x, y float64) {
	if i.isMoving {
		i.x = x
		i.y = y
	}
}

func (i *Item) Drop() {
	i.isMoving = false
}
