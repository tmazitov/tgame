package enemy

import "github.com/tmazitov/tgame.git/pkg/gm_geometry"

func (e *Enemy) IntersectVector(obj gm_geometry.IRect, x, y float64) bool {
	return e.coll.IsIntersectWithVector(obj, x, y)
}
