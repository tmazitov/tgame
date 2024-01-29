package player

import (
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

func (p *Player) MovementHandler(keys []ebiten.Key) {

	var (
		pressedKeyFound bool   = false
		pressedKeyArray []bool = []bool{false, false, false, false}
	)

	// TODO : bitmap to handle case then user press D and A

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

	if pressedKeyFound {
		if pressedKeyArray[0] {
			p.moveTop()
		}
		if pressedKeyArray[1] {
			p.moveLeft()
		}
		if pressedKeyArray[2] {
			p.moveBot()
		}
		if pressedKeyArray[3] {
			p.moveRight()
		}
	} else if p.actionState != Idle_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Idle_PlayerAction
	}
}

func (p *Player) moveLeft() {
	p.X -= p.Speed
	if p.actionState != Left_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Left_PlayerAction
	}
}
func (p *Player) moveRight() {
	p.X += p.Speed
	if p.actionState != Right_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Right_PlayerAction
	}
}
func (p *Player) moveTop() {
	p.Y -= p.Speed
	if p.actionState != Top_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Top_PlayerAction
	}
}
func (p *Player) moveBot() {
	p.Y += p.Speed
	if p.actionState != Bot_PlayerAction {
		p.lastAction = p.actionState
		p.actionState = Bot_PlayerAction
	}
}
