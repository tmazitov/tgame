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

	"github.com/tmazitov/tgame.git/internal/scene"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	gm_machine "github.com/tmazitov/tgame.git/pkg/gm_machine"
)

func main() {
	var (
		err  error
		font *gm_font.Font
		sc   *scene.Scene
	)

	if font, err = gm_font.NewFont("assets/fonts/pipel.png"); err != nil {
		panic(err)
	}

	game := gm_machine.NewGameMachine("Title")
	if game == nil {
		panic(err)
	}

	if sc, err = scene.NewScene(game, font); err != nil {
		panic(err)
	}

	if err = sc.Load(); err != nil {
		panic(err)
	}

	game.Run()
}
