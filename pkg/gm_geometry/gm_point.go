package gm_geometry

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Shift(x, y float64) *Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) Update(x, y float64) *Point {
	p.X = x
	p.Y = y
	return p
}
