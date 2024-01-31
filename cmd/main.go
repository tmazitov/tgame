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
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/tmazitov/tgame.git/internal/ground"
	gm_layer "github.com/tmazitov/tgame.git/pkg/gm_layer"
	gm_machine "github.com/tmazitov/tgame.git/pkg/gm_machine"
	gm_obj "github.com/tmazitov/tgame.git/pkg/gm_obj"
)

func main() {

	var tilesImage *ebiten.Image
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
	gameImage := gm_layer.NewImage(tilesImage)
	game := gm_machine.NewGameMachine("Title")
	if game == nil {
		panic("Game is nil!")
	}
	game.AddLayer(gm_layer.NewLayer("Layer 1", ground.GroundRaw(243), gameImage))
	game.AddObj(gm_obj.NewGameObj(
		"Test building",
		gm_obj.GameObjOptions{
			X:      1,
			Y:      2,
			Width:  6,
			Height: 14,
			Raw: []int{
				26, 27, 28, 29, 30, 31,
				51, 52, 53, 54, 55, 56,
				76, 77, 78, 79, 80, 81,
				101, 102, 103, 104, 105, 106,

				126, 127, 128, 129, 130, 131,
				303, 303, 245, 242, 303, 303,
				0, 0, 245, 242, 0, 0,
				0, 0, 245, 242, 0, 0,
				0, 0, 245, 242, 0, 0,

				0, 0, 245, 242, 0, 0,
				0, 0, 245, 242, 0, 0,
				0, 0, 245, 242, 0, 0,
				0, 0, 245, 242, 0, 0,
				0, 0, 245, 242, 0, 0,
			},
			Image: gameImage,
		},
	))
	game.Run()
}
