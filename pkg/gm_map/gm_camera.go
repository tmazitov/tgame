package gm_map

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X      float64
	Y      float64
	limitX float64
	limitY float64
	width  int
	height int
	speed  *float64
}

func NewCamera(height, width int) *Camera {
	return &Camera{
		X:      0,
		Y:      0,
		width:  width,
		height: height,
		speed:  nil,
	}
}

func (c *Camera) SetSpeed(speed *float64) {
	c.speed = speed
}

func (c *Camera) SetLimits(limitX, limitY float64) {
	c.limitX = limitX
	c.limitY = limitY
}

func checkIsDiagonalMovement(keys []bool) bool {
	return keys[0] && keys[3] ||
		keys[0] && keys[1] ||
		keys[2] && keys[3] ||
		keys[2] && keys[1]
}

func (c *Camera) MovementHandler(keys []ebiten.Key) (bool, error) {

	var (
		pressedKeyFound    bool   = false
		isDiagonalMovement bool   = false
		pressedKeyArray    []bool = []bool{false, false, false, false}
	)

	if c.speed == nil {
		return false, ErrZeroCameraSpeed
	}

	for _, key := range keys {
		if key == ebiten.KeyW {
			pressedKeyArray[0] = true
		}
		if key == ebiten.KeyA {
			pressedKeyArray[1] = true
		}
		if key == ebiten.KeyS {
			pressedKeyArray[2] = true
		}
		if key == ebiten.KeyD {
			pressedKeyArray[3] = true
		}
	}

	if pressedKeyArray[0] && pressedKeyArray[2] {
		pressedKeyArray[0] = false
		pressedKeyArray[2] = false
	}

	if pressedKeyArray[1] && pressedKeyArray[3] {
		pressedKeyArray[1] = false
		pressedKeyArray[3] = false
	}

	for _, key := range pressedKeyArray {
		pressedKeyFound = pressedKeyFound || key
	}

	isDiagonalMovement = checkIsDiagonalMovement(pressedKeyArray)

	if pressedKeyFound && isDiagonalMovement {
		return c.handleDiagonalMove(pressedKeyArray)
	} else if pressedKeyFound {
		return c.handleSimpleMove(pressedKeyArray)
	}
	return false, nil
}

func (c *Camera) handleSimpleMove(pressedKeyArray []bool) (bool, error) {
	if c.speed == nil {
		return false, ErrZeroCameraSpeed
	}

	if pressedKeyArray[0] {
		c.moveTop(*c.speed)
	}
	if pressedKeyArray[1] {
		c.moveLeft(*c.speed)
	}
	if pressedKeyArray[2] {
		c.moveBot(*c.speed)
	}
	if pressedKeyArray[3] {
		c.moveRight(*c.speed)
	}
	return true, nil
}

func (c *Camera) handleDiagonalMove(pressedKeyArray []bool) (bool, error) {
	if c.speed == nil {
		return false, ErrZeroCameraSpeed
	}

	if pressedKeyArray[0] {
		c.moveTop(*c.speed / math.Sqrt2)
	}
	if pressedKeyArray[1] {
		c.moveLeft(*c.speed / math.Sqrt2)
	}
	if pressedKeyArray[2] {
		c.moveBot(*c.speed / math.Sqrt2)
	}
	if pressedKeyArray[3] {
		c.moveRight(*c.speed / math.Sqrt2)
	}
	return true, nil
}

func (c *Camera) moveLeft(speed float64) {
	if c.X-speed < 0 {
		c.X = 0
		return
	}
	c.X -= speed
}
func (c *Camera) moveRight(speed float64) {
	if c.X+speed > c.limitX {
		c.X = c.limitX
		return
	}
	c.X += speed
}
func (c *Camera) moveTop(speed float64) {
	if c.Y-speed < 0 {
		c.Y = 0
		return
	}
	c.Y -= speed
}
func (c *Camera) moveBot(speed float64) {
	if c.Y+speed > c.limitY {
		c.Y = c.limitY
		return
	}
	c.Y += speed
}
