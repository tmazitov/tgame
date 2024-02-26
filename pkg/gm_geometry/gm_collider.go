package gm_geometry

type Collider struct {
	X           *float64
	Y           *float64
	height      float64
	width       float64
	paddingTop  float64
	paddingLeft float64

	// points []*Point
	topLeft  *Point
	topRight *Point
	botRight *Point
	botLeft  *Point

	// borders []*Line
	top   *Line
	right *Line
	bot   *Line
	left  *Line
}

type ColliderOptions struct {
	Height      float64
	Width       float64
	PaddingTop  float64
	PaddingLeft float64
}

func NewCollider(x, y *float64, opt ColliderOptions) *Collider {
	return &Collider{
		X:           x,
		Y:           y,
		height:      opt.Height,
		width:       opt.Width,
		paddingTop:  opt.PaddingTop,
		paddingLeft: opt.PaddingLeft,
		topLeft:     NewPoint(*x, *y),
		topRight:    NewPoint(*x+opt.Width, *y),
		botRight:    NewPoint(*x+opt.Width, *y+opt.Height),
		botLeft:     NewPoint(*x, *y+opt.Height),
		top:         NewLineByCoords(*x, *y, *x+opt.Width, *y),
		right:       NewLineByCoords(*x+opt.Width, *y, *x+opt.Width, *y+opt.Height),
		bot:         NewLineByCoords(*x, *y+opt.Height, *x+opt.Width, *y+opt.Height),
		left:        NewLineByCoords(*x, *y, *x, *y+opt.Height),
	}
}

func (c *Collider) Points() [4]*Point {

	var (
		collX float64 = *c.X
		collY float64 = *c.Y
	)

	if c.paddingLeft > 0 {
		collX += c.paddingLeft
	}
	if c.paddingTop > 0 {
		collY += c.paddingTop
	}

	return [4]*Point{
		c.topLeft.Update(collX, collY),
		c.topRight.Update(collX+c.width, collY),
		c.botRight.Update(collX+c.width, collY+c.height),
		c.botLeft.Update(collX+c.width, collY+c.height),
	}
}

func (c *Collider) Borders() [4]*Line {

	var (
		collX float64 = *c.X
		collY float64 = *c.Y
	)

	if c.paddingLeft > 0 {
		collX += c.paddingLeft
	}
	if c.paddingTop > 0 {
		collY += c.paddingTop
	}

	return [4]*Line{
		c.top.Update(collX, collY, collX+c.width, collY),
		c.right.Update(collX+c.width, collY, collX+c.width, collY+c.height),
		c.bot.Update(collX, collY+c.height, collX+c.width, collY+c.height),
		c.left.Update(collX, collY, collX, collY+c.height),
	}
}
func (c *Collider) IsIntersect(rect IRect) bool {
	var (
		borders     [4]*Line = rect.Borders()
		selfBorders [4]*Line = c.Borders()
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

func (c *Collider) IsIntersectWithVector(rect IRect, x, y float64) bool {
	var (
		borders     [4]*Line = rect.Borders()
		selfBorders [4]*Line = c.Borders()
	)
	for _, line := range borders {
		line.Shift(x, y)
		for _, selfLine := range selfBorders {
			if line.IsIntersect(selfLine) {
				return true
			}
		}
	}
	return false
}
