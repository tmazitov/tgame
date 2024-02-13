package prefabs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
)

type GameObj interface {
}

type Tree struct {
	X      float64
	Y      float64
	Width  int
	Height int
	Image  *ebiten.Image
	coll   *gm_geometry.Collider
}

func NewTree(x, y float64, image *ebiten.Image) *Tree {
	return &Tree{
		X:      x,
		Y:      y,
		Width:  48,
		Height: 64,
		Image:  image,
		coll:   gm_geometry.NewCollider(x, y, x+48, y+64),
	}
}

func (t *Tree) Draw(screen *ebiten.Image, camera *gm_map.Camera) {

	var (
		relativeX  float64
		relativeY  float64
		isOnScreen bool
	)

	relativeX, relativeY, isOnScreen = camera.GetRelativeCoords(t.X, t.Y)
	if !isOnScreen {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	screen.DrawImage(t.Image, op)
}
