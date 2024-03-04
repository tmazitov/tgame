// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "image/png"

	"github.com/tmazitov/tgame.git/internal/items"
	"github.com/tmazitov/tgame.git/internal/maps"
	"github.com/tmazitov/tgame.git/internal/player"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	gm_machine "github.com/tmazitov/tgame.git/pkg/gm_machine"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func main() {
	var (
		m                     *gm_map.Map
		err                   error
		f                     *gm_font.Font
		pl                    *player.Player
		itemCollectionStorage *gm_item.ItemCollectionStorage
	)

	if f, err = gm_font.NewFont("assets/fonts/pipel.png"); err != nil {
		panic(err)
	}

	pl, err = player.NewPlayer(0, 0, player.PlayerImagesPaths{
		Tiles:  "assets/textures/characters/Humans_Smith.png",
		Shadow: "assets/textures/characters/shadow.png",
	}, f)
	if err != nil {
		panic(err)
	}
	if pl == nil {
		panic("Player is nil!")
	}

	game := gm_machine.NewGameMachine("Title")
	if game == nil {
		panic("Game is nil!")
	}

	if m, err = maps.MainMap(); err != nil {
		panic(err)
	}

	itemCollectionStorage, err = gm_item.NewItemCollectionStorage("items/collectionsConfig.json", 32, f)
	if err != nil {
		panic(err)
	}

	m.AddCamera(gm_camera.NewCamera(stgs.ScreenHeight, stgs.ScreenWidth))
	game.SetupItemStorage(itemCollectionStorage)
	game.SetupPlayer(pl)

	item := game.ItemStorage.GetItem(items.MaterialsCollection, items.Stick).Clone(5)
	item.AutoDrop(10, 25)
	m.AddDropItem(item)

	item = game.ItemStorage.GetItem(items.FoodCollection, items.CherryPie).Clone(3)
	item.AutoDrop(70, 15)
	m.AddDropItem(item)

	game.AddMap(m)
	game.Run()
}
