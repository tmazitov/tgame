package gm_item

import (
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type ItemRaw struct {
	Id             uint               `json:"id"`
	Name           string             `json:"name"`
	ImagePath      string             `json:"imagePath"`
	SmallImagePath string             `json:"smallImagePath"`
	ImageSize      int                `json:"imageSize"`
	SmallImageSize int                `json:"smallImageSize"`
	MaxStack       uint               `json:"maxStack"`
	Description    ItemDescriptionRaw `json:"description"`
}

type ItemDescriptionRaw struct {
	Height      int `json:"height"`
	Width       int `json:"width"`
	TextPadding int `json:"textPadding"`
}

func (ir *ItemRaw) ToItem(font *gm_font.Font, descriptionSourceImage *gm_layer.Image) (*Item, error) {

	var (
		item *Item
		err  error
	)

	item, err = NewItem(ir.Id, ir.Name, ItemOptions{
		MaxStackSize:   ir.MaxStack,
		Amount:         1,
		ImageSize:      ir.ImageSize,
		ImagePath:      ir.ImagePath,
		SmallImagePath: ir.SmallImagePath,
		SmallImageSize: ir.SmallImageSize,
	})

	if err != nil {
		return nil, err
	}

	if err = item.SetupDescription(descriptionSourceImage, ItemDescriptionOpt{
		TextPadding: ir.Description.TextPadding,
		Height:      ir.Description.Height,
		Width:       ir.Description.Width,
		Font:        font,
	}); err != nil {
		return nil, err
	}

	return item, err
}
