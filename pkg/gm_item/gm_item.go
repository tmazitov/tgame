package gm_item

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Item struct {
	ID           uint
	Name         string
	MaxStackSize uint
	Amount       uint
	image        *gm_layer.Image
	smallImage   *gm_layer.Image
	description  *ItemDescription
	X            float64
	Y            float64
	isMoving     bool
	IsDropped    bool
	shape        *gm_geometry.Rect
	lastDropTime time.Time
	dropProcess  *ItemDropPath
}

type ItemOptions struct {
	ImagePath      string
	ImageSize      int
	SmallImagePath string
	SmallImageSize int
	MaxStackSize   uint
	Amount         uint
	X              float64
	Y              float64
}

func NewItem(id uint, name string, opt ItemOptions) (*Item, error) {

	var (
		shape      *gm_geometry.Rect
		image      *gm_layer.Image
		smallImage *gm_layer.Image
		err        error
	)

	if opt.ImagePath == "" || name == "" || opt.SmallImagePath == "" {
		return nil, ErrInvalidParams
	}

	if image, err = gm_layer.NewImageByPath(opt.ImagePath, opt.ImageSize); err != nil {
		return nil, err
	}
	if smallImage, err = gm_layer.NewImageByPath(opt.SmallImagePath, opt.SmallImageSize); err != nil {
		return nil, err
	}

	if opt.Amount == 0 {
		opt.Amount = 1
	}

	shape = gm_geometry.NewRect(&opt.X, &opt.Y, smallImage.Width(), smallImage.Height())

	return &Item{
		ID:           id,
		Name:         name,
		MaxStackSize: opt.MaxStackSize,
		image:        image,
		smallImage:   smallImage,
		X:            opt.X,
		Y:            opt.Y,
		isMoving:     false,
		IsDropped:    true,
		Amount:       opt.Amount,
		description:  nil,
		shape:        shape,
		lastDropTime: time.Time{},
	}, nil
}

func (i *Item) SetupDescription(source *gm_layer.Image, opt ItemDescriptionOpt) error {
	var err error
	if i.description, err = NewItemDescription(i.Name, source, opt); err != nil {
		return err
	}
	return nil
}

func (i *Item) InDropProcess() bool {
	return i.dropProcess != nil
}

func (i *Item) Clone(amount uint) *Item {

	if amount == 0 {
		amount = 1
	}

	var item Item = Item{
		ID:           i.ID,
		Name:         i.Name,
		MaxStackSize: i.MaxStackSize,
		Amount:       amount,
		image:        i.image,
		X:            i.X,
		Y:            i.Y,
		isMoving:     false,
		smallImage:   i.smallImage,
		IsDropped:    i.IsDropped,
		description:  i.description,
		lastDropTime: i.lastDropTime,
	}

	item.shape = gm_geometry.NewRect(&item.X, &item.Y, item.smallImage.Width(), item.smallImage.Height())

	return &item
}

func (i *Item) Update() {

	var (
		step gm_geometry.Point
	)

	if i.dropProcess == nil {
		return
	}
	if i.dropProcess.IsFinished() {
		i.dropProcess = nil
		return
	}
	step = i.dropProcess.Increment()
	i.SetPosition(step.X, step.Y)
}

func (i *Item) Drop(source *gm_geometry.Point, target *gm_geometry.Point) {
	i.dropProcess = NewItemDropPath(source, target)
	i.IsDropped = true
	i.lastDropTime = time.Now()
}

func (i *Item) AutoDrop(x, y float64) {
	i.SetPosition(x, y)
	i.IsDropped = true
}

func (i *Item) IsCollectable() bool {
	var now time.Time = time.Now()
	diff := now.Sub(i.lastDropTime)
	return diff.Seconds() > 3
}

func (i *Item) Draw(screen *ebiten.Image, camera *gm_camera.Camera) {
	var (
		positionX  float64
		positionY  float64
		isInCamera bool
	)

	if !i.IsDropped || camera == nil {
		op := &ebiten.DrawImageOptions{}
		if i.image.TileSize == 16 {
			op.GeoM.Scale(2, 2)
		}
		op.GeoM.Translate(i.X, i.Y)
		screen.DrawImage(i.image.Inst, op)
		return
	}

	positionX, positionY, isInCamera = camera.GetRelativeCoordsByRect(i.shape)
	if !isInCamera {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(positionX, positionY)
	screen.DrawImage(i.smallImage.Inst, op)
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
	i.X = x
	i.Y = y
}

func (i *Item) GetPosition() (float64, float64) {
	return i.X, i.Y
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

func (i *Item) DrawDescription(screen *ebiten.Image) {
	if i.description != nil {
		i.description.Draw(i.X, i.Y+32, screen)
	}
}

func (i *Item) Shape() *gm_geometry.Rect {
	return i.shape
}

func (i *Item) Size() int {
	if i.IsDropped {
		return i.smallImage.Height()
	}
	return i.image.Height()
}
