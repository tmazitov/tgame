package prefabs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
)

type Tree struct {
	X      float64
	Y      float64
	Width  int
	Height int
	Image  *ebiten.Image
	coll   *gm_geometry.Collider
	shape  *gm_geometry.Rect
}

func NewTree(x, y float64, image *ebiten.Image) *Tree {

	var (
		width       int     = 48
		height      int     = 64
		collWidth   float64 = 16
		collHeight  float64 = 4
		paddingTop  float64 = float64(height) - 10
		paddingLeft float64 = float64(width)/2 - collWidth/2
		tree        *Tree
	)

	tree = &Tree{
		X:      x,
		Y:      y,
		Width:  48,
		Height: 64,
		Image:  image,
		coll:   nil,
		shape:  nil,
	}

	tree.coll = gm_geometry.NewCollider(&tree.X, &tree.Y, gm_geometry.ColliderOptions{
		Height:      collHeight,
		Width:       collWidth,
		PaddingTop:  paddingTop,
		PaddingLeft: paddingLeft,
	})
	tree.shape = gm_geometry.NewRect(&x, &y, width, height)
	return tree
}

func (t *Tree) GetCollider() *gm_geometry.Collider {
	return t.coll
}

func (t *Tree) Draw(screen *ebiten.Image, camera *gm_camera.Camera) {

	var (
		relativeX  float64
		relativeY  float64
		isOnScreen bool
	)

	relativeX, relativeY, isOnScreen = camera.GetRelativeCoordsByRect(t.shape)
	if !isOnScreen {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	screen.DrawImage(t.Image, op)
}

func (t *Tree) IntersectVector(obj gm_geometry.IRect, x, y float64) bool {
	return t.coll.IsIntersectWithVector(obj, x, y)
}

func (t *Tree) Intersect(obj gm_geometry.IRect) bool {
	return t.coll.IsIntersect(obj)
}
