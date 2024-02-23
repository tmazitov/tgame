package gm_item

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

func (ir *ItemRaw) ToItem() (*Item, error) {
	return NewItem(ir.Id, ir.Name, ir.ImagePath, ItemOptions{
		MaxStackSize: ir.MaxStack,
		Amount:       1,
		TileSize:     ir.ImageSize,
	})
}
