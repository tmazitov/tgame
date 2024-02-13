package gm_geometry

type Collider struct {
	topLeft  *Point
	topRight *Point
	botLeft  *Point
	botRight *Point

	top    *Line
	right  *Line
	bottom *Line
	left   *Line
}

func NewCollider(x, y, height, width float64) *Collider {
	var coll Collider = Collider{}
	coll.topLeft = NewPoint(x, y)
	coll.topRight = NewPoint(x+width, y)
	coll.botLeft = NewPoint(x, y+height)
	coll.botRight = NewPoint(x+width, y+height)

	coll.top = NewLine(coll.topLeft, coll.topRight)
	coll.right = NewLine(coll.topRight, coll.botRight)
	coll.bottom = NewLine(coll.botRight, coll.botLeft)
	coll.left = NewLine(coll.botLeft, coll.topLeft)
	return &coll
}

func (c *Collider) Points() []*Point {
	return []*Point{
		c.topLeft,
		c.topRight,
		c.botRight,
		c.botLeft,
	}
}

func (c *Collider) GetBorders() []*Line {
	var lines []*Line = []*Line{
		NewLine(c.topLeft, c.topRight),
		NewLine(c.topRight, c.botRight),
		NewLine(c.botRight, c.botLeft),
		NewLine(c.botLeft, c.topLeft),
	}
	return lines
}
func (c *Collider) IsIntersect(coll *Collider) bool {
	var (
		borders     []*Line = coll.GetBorders()
		selfBorders []*Line = c.GetBorders()
	)
	for _, line := range borders {
		for _, selfLine := range selfBorders {
			if line.IsIntersect(selfLine) {
				return true
			}
		}
	}
	return false
}

func (c *Collider) IsIntersectWithVector(coll *Collider, x, y float64) bool {
	var (
		borders     []*Line = coll.GetBorders()
		selfBorders []*Line = c.GetBorders()
	)
	for _, line := range borders {
		for _, selfLine := range selfBorders {
			if line.IsIntersect(selfLine) {
				return true
			}
		}
	}
	return false
}
