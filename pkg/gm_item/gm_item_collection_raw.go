package gm_item

import "github.com/tmazitov/tgame.git/pkg/gm_font"

type ItemCollectionRaw struct {
	Name                  string `json:"name"`
	ItemsPath             string `json:"itemsPath"`
	DescriptionSourcePath string `json:"descriptionSourcePath"`
}

func (icr *ItemCollectionRaw) ToItemCollection(font *gm_font.Font, itemSize int) (*ItemCollection, error) {

	var (
		collection *ItemCollection
		err        error
	)

	collection, err = NewItemCollection(icr.Name, ItemCollectionOpt{
		ItemSize:              itemSize,
		DescriptionSourcePath: icr.DescriptionSourcePath,
		Font:                  font,
	})
	if err != nil {
		return nil, err
	}

	if err = collection.FillByJson(icr.ItemsPath); err != nil {
		return nil, err
	}

	return collection, nil
}
