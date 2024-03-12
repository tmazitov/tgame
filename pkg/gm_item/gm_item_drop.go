package gm_item

import (
	"math"

	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
)

type ItemDropPath struct {
	source         *gm_geometry.Point
	target         *gm_geometry.Point
	current        *gm_geometry.Point
	hypotenuseApex *gm_geometry.Point
	isHorizontal   bool
	iterations     int
	difference     float64
}

func NewItemDropPath(source, target *gm_geometry.Point) *ItemDropPath {

	var (
		hypoX        float64
		hypoY        float64
		isHorizontal bool
		difference   float64
	)

	if source.Y == target.Y {
		hypoX = ((target.X - source.X) / 2) + source.X
		hypoY = target.Y - 12
		isHorizontal = true
		difference = math.Abs(target.X - source.X)
	} else if source.X == target.X {
		hypoX = source.X
		hypoY = ((target.Y - source.Y) / 2) + source.Y
		isHorizontal = false
		difference = math.Abs(target.Y - source.Y)
	} else {
		return nil
	}

	return &ItemDropPath{
		source:         source,
		target:         target,
		current:        gm_geometry.NewPoint(source.X, source.Y),
		hypotenuseApex: gm_geometry.NewPoint(hypoX, hypoY),
		isHorizontal:   isHorizontal,
		iterations:     0,
		difference:     difference / 2,
	}
}

func (idp *ItemDropPath) calcNextYValue(x float64) float64 {
	return math.Pow(x, 2) / 20
}

func (idp *ItemDropPath) calcSpeed(value float64) float64 {

	return (2/math.Pow(idp.difference, 2))*math.Pow(value, 2) + 1
}

func (idp *ItemDropPath) incrementX() float64 {

	var speed float64 = idp.calcSpeed(idp.current.Y - idp.hypotenuseApex.Y)
	// var speed float64 = 1

	if idp.source.X < idp.target.X {
		idp.current.X += speed
	} else {
		idp.current.X -= speed
	}
	return idp.current.X
}

func (idp *ItemDropPath) incrementY() float64 {
	var speed float64 = idp.calcSpeed(idp.current.Y - idp.hypotenuseApex.Y)

	if idp.source.Y < idp.target.Y {
		idp.current.Y += speed
	} else {
		idp.current.Y -= speed
	}
	return idp.current.Y
}
func (idp *ItemDropPath) IsFinished() bool {

	if idp.iterations == 0 {
		return false
	}

	if idp.isHorizontal {
		return idp.source.X < idp.target.X && idp.current.X > idp.target.X ||
			idp.source.X > idp.target.X && idp.current.X < idp.target.X
	}

	if !idp.isHorizontal {
		return idp.source.Y > idp.target.Y && idp.current.Y < idp.target.Y ||
			idp.source.Y < idp.target.Y && idp.current.Y > idp.target.Y
	}

	return false
}

func (idp *ItemDropPath) Increment() gm_geometry.Point {

	var (
		stepX float64 = 0
		stepY float64 = 0
	)

	idp.iterations += 1
	if idp.isHorizontal {
		stepX = idp.incrementX()
		stepY = idp.source.Y + idp.calcNextYValue(stepX-idp.hypotenuseApex.X)
		idp.current.Y = stepY
	} else {
		stepX = idp.source.X
		stepY = idp.incrementY()
		idp.current.X = stepX
	}
	return gm_geometry.Point{X: stepX, Y: stepY}
}
