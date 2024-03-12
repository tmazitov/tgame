package gm_map

import "github.com/tmazitov/tgame.git/pkg/gm_item"

func (m *Map) AddDropItem(item *gm_item.Item) {
	m.droppedItems = append(m.droppedItems, item)
}

func (m *Map) DelDropItem(item *gm_item.Item) {
	for i, droppedItem := range m.droppedItems {
		if droppedItem == item {
			// Remove the item from the array
			m.droppedItems = append(m.droppedItems[:i], m.droppedItems[i+1:]...)
			break
		}
	}
}
