package gm_geometry

type IRect interface {
	Points() [4]*Point
	Borders() [4]*Line
}

type Rect struct {
	X        *float64
	Y        *float64
	width    float64
	height   float64
	topLeft  *Point
	topRight *Point
	botRight *Point
	botLeft  *Point
	top      *Line
	right    *Line
	bot      *Line
	left     *Line
}

func NewRect(x, y *float64, width, height int) *Rect {
	var rect Rect = Rect{
		X:      x,
		Y:      y,
		width:  float64(width),
		height: float64(height),
	}

	rect.topLeft = NewPoint(*x, *y)
	rect.topRight = NewPoint(*x+rect.width, *y)
	rect.botRight = NewPoint(*x+rect.width, *y+rect.height)
	rect.botLeft = NewPoint(*x, *y+rect.height)
	rect.top = NewLineByCoords(rect.topLeft.X, rect.topLeft.Y, rect.topRight.X, rect.topRight.Y)
	rect.right = NewLineByCoords(rect.topRight.X, rect.topRight.Y, rect.botRight.X, rect.botRight.Y)
	rect.bot = NewLineByCoords(rect.botRight.X, rect.botRight.Y, rect.botLeft.X, rect.botLeft.Y)
	rect.left = NewLineByCoords(rect.botLeft.X, rect.botLeft.Y, rect.topLeft.X, rect.topLeft.Y)
	return &rect
}

func (c *Rect) Borders() [4]*Line {

	var (
		collX float64 = *c.X
		collY float64 = *c.Y
	)

	return [4]*Line{
		c.top.Update(collX, collY, collX+c.width, collY),
		c.right.Update(collX+c.width, collY, collX+c.width, collY+c.height),
		c.bot.Update(collX, collY+c.height, collX+c.width, collY+c.height),
		c.left.Update(collX, collY, collX, collY+c.height),
	}
}

func (r *Rect) Points() [4]*Point {

	var (
		collX float64 = *r.X
		collY float64 = *r.Y
	)

	return [4]*Point{
		r.topLeft.Update(collX, collY),
		r.topRight.Update(collX+r.width, collY),
		r.botRight.Update(collX+r.width, collY+r.height),
		r.botLeft.Update(collX+r.width, collY+r.height),
	}
}
