package player

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type FireballRemoveCallback func(*PlayerFireball)
type PlayerFireball struct {
	X              float32
	Y              float32
	tileCounter    int
	tileIterator   int
	dest           *gm_geometry.Point
	images         []*ebiten.Image
	removeCallback FireballRemoveCallback
}

func NewPlayerFireball(x, y float32, dest *gm_geometry.Point, images []*ebiten.Image, removeCallback FireballRemoveCallback) *PlayerFireball {

	var fireball *PlayerFireball = &PlayerFireball{
		X:              x,
		Y:              y,
		dest:           dest,
		images:         images,
		tileCounter:    0,
		tileIterator:   0,
		removeCallback: removeCallback,
	}

	return fireball
}

func (pf *PlayerFireball) Move() {
	pf.X += stgs.FireballSpeed
	if pf.X > pf.dest.X {
		pf.removeCallback(pf)
	}
}

func (pf *PlayerFireball) Draw(screen *ebiten.Image) {
	if pf.tileIterator == 8 {
		pf.tileCounter++
		pf.tileIterator = 0
	}
	if pf.tileCounter == 5 {
		pf.tileCounter = 0
	}
	pf.tileIterator++

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(pf.X), float64(pf.Y))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Fireball: %f %f\n", pf.X, pf.Y))
	screen.DrawImage(pf.images[pf.tileCounter], op)
}
