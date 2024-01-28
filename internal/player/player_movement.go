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

	var pressedKeyFound bool = false

	// TODO : bitmap to handle case then user press D and A

	for _, key := range keys {
		if key == ebiten.KeyW {
			p.moveTop()
			pressedKeyFound = true
		}
		if key == ebiten.KeyS {
			p.moveBot()
			pressedKeyFound = true
		}
		if key == ebiten.KeyD {
			p.moveRight()
			pressedKeyFound = true
		}
		if key == ebiten.KeyA {
			p.moveLeft()
			pressedKeyFound = true
		}
	}

	if !pressedKeyFound && p.actionState != Idle_PlayerAction {
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
