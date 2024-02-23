package gm_machine

import (
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
)

func (g *GameMachine) SetupItemStorage(storage *gm_item.ItemCollectionStorage) {
	if storage == nil {
		return
	}
	g.ItemStorage = storage
}

func (g *GameMachine) SetupPlayer(player gm_entity.Player) {
	if player == nil {
		return
	}
	g.player = player
}
