package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerAction int

const (
	Idle_PlayerAction  PlayerAction = 0
	Right_PlayerAction PlayerAction = 1
	Left_PlayerAction  PlayerAction = 2
	Top_PlayerAction   PlayerAction = 3
	Bot_PlayerAction   PlayerAction = 4
)

func checkIsDiagonalMovement(keys []bool) bool {
	return keys[0] && keys[3] ||
		keys[0] && keys[1] ||
		keys[2] && keys[3] ||
		keys[2] && keys[1]
}

func (p *Player) MovementHandler(keys []ebiten.Key) {

	var (
		pressedKeyFound    bool   = false
		isDiagonalMovement bool   = false
		pressedKeyArray    []bool = []bool{false, false, false, false}
	)

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
		p.handleDiagonalMove(pressedKeyArray)
	} else if pressedKeyFound {
		p.handleSimpleMove(pressedKeyArray)
	} else if p.actionState != Idle_PlayerAction {
		p.idle()
	}
}

func (p *Player) handleSimpleMove(pressedKeyArray []bool) {
	if pressedKeyArray[0] {
		p.moveTop(p.Speed)
	}
	if pressedKeyArray[1] {
		p.moveLeft(p.Speed)
	}
	if pressedKeyArray[2] {
		p.moveBot(p.Speed)
	}
	if pressedKeyArray[3] {
		p.moveRight(p.Speed)
	}
}

func (p *Player) handleDiagonalMove(pressedKeyArray []bool) {
	if pressedKeyArray[0] {
		p.moveTop(p.Speed / math.Sqrt2)
	}
	if pressedKeyArray[1] {
		p.moveLeft(p.Speed / math.Sqrt2)
	}
	if pressedKeyArray[2] {
		p.moveBot(p.Speed / math.Sqrt2)
	}
	if pressedKeyArray[3] {
		p.moveRight(p.Speed / math.Sqrt2)
	}
}

func (p *Player) idle() {
	p.lastAction = p.actionState
	p.actionState = Idle_PlayerAction
}

func (p *Player) moveLeft(speed float32) {
	p.X -= speed
	if p.actionState != Left_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Left_PlayerAction
	}
}
func (p *Player) moveRight(speed float32) {
	p.X += speed
	if p.actionState != Right_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Right_PlayerAction
	}
}
func (p *Player) moveTop(speed float32) {
	p.Y -= speed
	if p.actionState != Top_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Top_PlayerAction
	}
}
func (p *Player) moveBot(speed float32) {
	p.Y += speed
	if p.actionState != Bot_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Bot_PlayerAction
	}
}
