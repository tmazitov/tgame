package gm_geometry

type Line struct {
	Start *Point
	End   *Point
}

func NewLine(start, end *Point) *Line {
	return &Line{
		Start: start,
		End:   end,
	}
}

func NewLineByCoords(x1, y1, x2, y2 float64) *Line {
	return &Line{
		Start: NewPoint(x1, y1),
		End:   NewPoint(x2, y2),
	}
}

func (l *Line) Shift(x, y float64) *Line {
	l.Start.X += x
	l.Start.Y += y
	l.End.X += x
	l.End.Y += y
	return l
}

func (l *Line) Update(x1, y1, x2, y2 float64) *Line {
	l.Start.X = x1
	l.Start.Y = y1
	l.End.X = x2
	l.End.Y = y2
	return l
}

func (l *Line) IsIntersect(line *Line) bool {
	var (
		x1, x2, x3, x4 float64 = l.Start.X, l.End.X, line.Start.X, line.End.X
		y1, y2, y3, y4 float64 = l.Start.Y, l.End.Y, line.Start.Y, line.End.Y
	)
	var (
		denom float64 = (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
		numA  float64 = (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
		numB  float64 = (x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)
	)
	if denom == 0 {
		return false
	}
	var ua float64 = numA / denom
	var ub float64 = numB / denom
	return ua >= 0 && ua <= 1 && ub >= 0 && ub <= 1
}

func (l *Line) IsIntersectWithVector(line *Line, x, y float64) bool {
	var (
		x1, x2, x3, x4 float64 = l.Start.X, l.End.X, line.Start.X + x, line.End.X + x
		y1, y2, y3, y4 float64 = l.Start.Y, l.End.Y, line.Start.Y + y, line.End.Y + y
	)
	var (
		denom float64 = (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
		numA  float64 = (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
		numB  float64 = (x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)
	)
	if denom == 0 {
		return false
	}
	var ua float64 = numA / denom
	var ub float64 = numB / denom
	return ua >= 0 && ua <= 1 && ub >= 0 && ub <= 1
}
