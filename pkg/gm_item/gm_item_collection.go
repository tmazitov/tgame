package gm_item

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type ItemCollection struct {
	Name              string
	items             map[uint]*Item
	descriptionSource *gm_layer.Image
	font              *gm_font.Font
}

type ItemCollectionOpt struct {
	DescriptionSourcePath string
	ItemSize              int
	Font                  *gm_font.Font
}

func NewItemCollection(name string, opt ItemCollectionOpt) (*ItemCollection, error) {

	var (
		descriptionSource *gm_layer.Image
		err               error
	)

	if opt.DescriptionSourcePath == "" {
		return nil, ErrEmptyDescriptionSourcePath
	}

	if opt.Font == nil {
		return nil, ErrNilFont
	}

	if opt.ItemSize == 0 {
		return nil, ErrZeroItemSize
	}

	if descriptionSource, err = gm_layer.NewImageByPath(opt.DescriptionSourcePath, 16); err != nil {
		return nil, err
	}

	return &ItemCollection{
		Name:              name,
		items:             map[uint]*Item{},
		font:              opt.Font,
		descriptionSource: descriptionSource,
	}, nil
}

func (c *ItemCollection) FillByJson(path string) error {
	var (
		item     *Item
		itemsRaw []ItemRaw
		err      error
	)

	if path == "" {
		return ErrEmptyJsonPath
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if json.Unmarshal(data, &itemsRaw); err != nil {
		return err
	}

	for _, itemRaw := range itemsRaw {
		item, err = itemRaw.ToItem(c.font, c.descriptionSource)
		if err != nil {
			return err
		}
		c.items[itemRaw.Id] = item
	}

	return nil
}

func (c *ItemCollection) AddItem(id uint, name string, imagePath string, iOpt ItemOptions, dOpt ItemDescriptionOpt) error {
	var (
		item *Item
		err  error
	)

	if item, err = NewItem(id, name, iOpt); err != nil {
		return err
	}
	if err = item.SetupDescription(c.descriptionSource, dOpt); err != nil {
		return err
	}
	c.items[id] = item
	return nil
}

func (c *ItemCollection) GetItem(id uint) *Item {

	item, ok := c.items[id]
	if !ok {
		log.Printf("error : undefined item type %d", id)
		return nil
	}

	return item
}
