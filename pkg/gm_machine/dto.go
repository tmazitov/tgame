package gm_machine

import "github.com/hajimehoshi/ebiten/v2"

type ILayer interface {
	Draw(screen *ebiten.Image)
}
