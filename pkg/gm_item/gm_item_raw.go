package gm_item

import "github.com/tmazitov/tgame.git/pkg/gm_layer"

type ItemRaw struct {
	Id          uint               `json:"id"`
	Name        string             `json:"name"`
	ImagePath   string             `json:"imagePath"`
	ImageSize   int                `json:"imageSize"`
	MaxStack    uint               `json:"maxStack"`
	Description ItemDescriptionRaw `json:"description"`
}

type ItemDescriptionRaw struct {
	Height      int `json:"height"`
	Width       int `json:"width"`
	TextPadding int `json:"textPadding"`
}

func (ir *ItemRaw) ToItem(descriptionSourceImage *gm_layer.Image) (*Item, error) {

	var (
		item *Item
		err  error
	)

	item, err = NewItem(ir.Id, ir.Name, ir.ImagePath, ItemOptions{
		MaxStackSize: ir.MaxStack,
		Amount:       1,
		TileSize:     ir.ImageSize,
	})

	if err != nil {
		return nil, err
	}

	if err = item.SetupDescription(descriptionSourceImage, ItemDescriptionOpt{
		TextPadding: ir.Description.TextPadding,
		Height:      ir.Description.Height,
		Width:       ir.Description.Width,
	}); err != nil {
		return nil, err
	}

	return item, err
}
