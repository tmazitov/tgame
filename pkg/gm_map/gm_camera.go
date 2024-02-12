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
		limitX: 0,
		limitY: 0,
	}
}

func (c *Camera) SetSpeed(speed *float64) {
	c.speed = speed
}

func (c *Camera) SetLimits(limitX, limitY float64) {
	c.limitX = limitX - float64(c.width)
	c.limitY = limitY - float64(c.height)
}

func (c *Camera) GetRelativeCoords(x, y float64) (float64, float64, bool) {

	var (
		relativeX  float64
		relativeY  float64
		isInCamera bool
	)

	relativeX = x
	relativeY = y
	isInCamera = relativeX >= 0 && relativeX <= float64(c.width) &&
		relativeY >= 0 && relativeY <= float64(c.height)
	return relativeX, relativeY, isInCamera
}

func (c *Camera) GetPointArea(x, y float64) CameraArea {

	var (
		relX       float64
		relY       float64
		width      float64 = float64(c.width)
		height     float64 = float64(c.height)
		isInCamera bool
	)

	if relX, relY, isInCamera = c.GetRelativeCoords(x, y); !isInCamera {
		return NoneCameraArea
	}

	// Top Left corner
	if relX >= 0 && relX <= CameraBorderSize &&
		relY >= 0 && relY <= CameraBorderSize {
		return TopLeftCornerArea
	}

	// Top Right corner
	if relX >= width-CameraBorderSize && relX <= width &&
		relY >= 0 && relY <= CameraBorderSize {
		return TopRightCornerArea
	}

	// Bot Left corner
	if relX >= 0 && relX <= CameraBorderSize &&
		relY >= height-CameraBorderSize && relY <= height {
		return BotLeftCornerArea
	}

	// Bot Right corner
	if relX >= width-CameraBorderSize && relX <= width &&
		relY >= height-CameraBorderSize && relY <= height {
		return BotRightCornerArea
	}

	// Top border
	if relX > CameraBorderSize && relX < width-CameraBorderSize &&
		relY >= 0 && relY <= CameraBorderSize {
		return TopBorderCameraArea
	}

	// Bot border
	if relX > CameraBorderSize && relX < width-CameraBorderSize &&
		relY >= height-CameraBorderSize && relY <= height {
		return BotBorderCameraArea
	}

	// Left border
	if relX >= 0 && relX <= CameraBorderSize &&
		relY > CameraBorderSize && relY < height-CameraBorderSize {
		return LeftBorderCameraArea
	}

	// Right border
	if relX >= width-CameraBorderSize && relX <= width &&
		relY > CameraBorderSize && relY < height-CameraBorderSize {
		return RightBorderCameraArea
	}

	return FreeCameraArea
}

func checkIsDiagonalMovement(keys []bool) bool {
	return keys[0] && keys[3] ||
		keys[0] && keys[1] ||
		keys[2] && keys[3] ||
		keys[2] && keys[1]
}

func (c *Camera) MovementHandler(keys []ebiten.Key, area CameraArea) (bool, error) {

	var (
		pressedKeyFound    bool   = false
		isDiagonalMovement bool   = false
		pressedKeyArray    []bool = []bool{false, false, false, false}
	)

	if c.speed == nil {
		return false, ErrZeroCameraSpeed
	}

	if area == NoneCameraArea {
		return false, nil
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
		return c.handleDiagonalMove(pressedKeyArray, area)
	} else if pressedKeyFound {
		return c.handleSimpleMove(pressedKeyArray, area)
	}
	return false, nil
}

func (c *Camera) handleSimpleMove(pressedKeyArray []bool, area CameraArea) (bool, error) {

	if c.speed == nil {
		return false, ErrZeroCameraSpeed
	}

	if pressedKeyArray[0] && (area == TopBorderCameraArea || area == TopLeftCornerArea || area == TopRightCornerArea) {
		return c.moveTop(*c.speed)
	}
	if pressedKeyArray[1] && (area == LeftBorderCameraArea || area == TopLeftCornerArea || area == BotLeftCornerArea) {
		return c.moveLeft(*c.speed)
	}
	if pressedKeyArray[2] && (area == BotBorderCameraArea || area == BotRightCornerArea || area == BotLeftCornerArea) {
		return c.moveBot(*c.speed)
	}
	if pressedKeyArray[3] && (area == RightBorderCameraArea || area == TopRightCornerArea || area == BotRightCornerArea) {
		return c.moveRight(*c.speed)
	}
	return false, nil
}

func (c *Camera) handleDiagonalMove(pressedKeyArray []bool, area CameraArea) (bool, error) {

	var (
		isMoved bool  = false
		err     error = nil
	)

	if c.speed == nil {
		return false, ErrZeroCameraSpeed
	}

	if pressedKeyArray[0] && ((area == TopLeftCornerArea && pressedKeyArray[1]) || (area == TopRightCornerArea && pressedKeyArray[3])) {
		isMoved, err = c.moveTop(*c.speed / math.Sqrt2)
	}
	if pressedKeyArray[1] && ((area == TopLeftCornerArea && pressedKeyArray[0]) || (area == BotLeftCornerArea && pressedKeyArray[2])) {
		isMoved, err = c.moveLeft(*c.speed / math.Sqrt2)
	}
	if pressedKeyArray[2] && ((area == BotLeftCornerArea && pressedKeyArray[1]) || (area == BotRightCornerArea && pressedKeyArray[3])) {
		isMoved, err = c.moveBot(*c.speed / math.Sqrt2)
	}
	if pressedKeyArray[3] && ((area == TopRightCornerArea && pressedKeyArray[0]) || (area == BotRightCornerArea && pressedKeyArray[2])) {
		isMoved, err = c.moveRight(*c.speed / math.Sqrt2)
	}
	return isMoved, err
}

func (c *Camera) moveLeft(speed float64) (bool, error) {
	if c.limitX != 0 && c.X-speed < 0 {
		c.X = 0
		return false, nil
	}
	c.X -= speed
	return true, nil
}
func (c *Camera) moveRight(speed float64) (bool, error) {
	if c.limitX != 0 && c.X+speed > c.limitX {
		c.X = c.limitX
		return false, nil
	}
	c.X += speed
	return true, nil
}
func (c *Camera) moveTop(speed float64) (bool, error) {
	if c.limitY != 0 && c.Y-speed < 0 {
		c.Y = 0
		return false, nil
	}
	c.Y -= speed
	return true, nil
}
func (c *Camera) moveBot(speed float64) (bool, error) {
	if c.limitY != 0 && c.Y+speed > c.limitY {
		c.Y = c.limitY
		return false, nil
	}
	c.Y += speed
	return true, nil
}
