package gm_item

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type ItemCollectionStorage struct {
	collections      []*ItemCollection
	collectionsNames []string
}

func NewItemCollectionStorage(configPath string, itemSize int) (*ItemCollectionStorage, error) {

	var (
		collectionsRaw   []*ItemCollectionRaw
		collection       *ItemCollection
		collections      []*ItemCollection = []*ItemCollection{}
		collectionsNames []string          = []string{}
	)

	if configPath == "" {
		return nil, ErrEmptyConfigPath
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if json.Unmarshal(data, &collectionsRaw); err != nil {
		return nil, err
	}

	for _, raw := range collectionsRaw {
		collection, err = raw.ToItemCollection(itemSize)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
		collectionsNames = append(collectionsNames, collection.Name)
	}

	return &ItemCollectionStorage{
		collections:      collections,
		collectionsNames: collectionsNames,
	}, nil
}

func (ics *ItemCollectionStorage) GetItem(collectionName string, itemID uint) *Item {

	var (
		collectionIndex int = -1
	)

	for index, name := range ics.collectionsNames {
		if name == collectionName {
			collectionIndex = index
			break
		}
	}

	if collectionIndex == -1 {
		return nil
	}

	return ics.collections[collectionIndex].GetItem(itemID)
}
