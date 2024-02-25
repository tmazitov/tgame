package gm_geometry

type IRect interface {
	Points() [4]*Point
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
	return &rect
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
