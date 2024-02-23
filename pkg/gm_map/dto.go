package gm_map

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
)

type MapLevel int

const (
	MapGroundLevel MapLevel = 1
)

type IMapObj interface {
	Draw(screen *ebiten.Image, camera *gm_camera.Camera)
	GetCollider() *gm_geometry.Collider
	IntersectVector(obj gm_geometry.IMapIntersectable, x, y float64) bool
}

// type IDropedItem interface {
// 	SetPosition(x, y float64)
// 	Draw(screen *ebiten.Image, camera *gm_camera.Camera)
// }
