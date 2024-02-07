package player

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
)

type PlayerAttackSystem struct {
	fireballImages []*ebiten.Image
	playerX        *float32
	playerY        *float32
	playerLastMove *PlayerAction
	fireballs      []*PlayerFireball
	block          bool
}

func (p *Player) AttackHandler(keys []ebiten.Key) {
	p.attack.Handle(keys)
}

func NewPlayerAttackSystem(playerX *float32, playerY *float32, playerLastMove *PlayerAction) (*PlayerAttackSystem, error) {
	var (
		file        *os.File
		err         error
		pas         *PlayerAttackSystem
		img         image.Image
		images      []*ebiten.Image = []*ebiten.Image{}
		imagesPaths                 = []string{
			"../assets/textures/fireball/fireball_1.png",
			"../assets/textures/fireball/fireball_2.png",
			"../assets/textures/fireball/fireball_3.png",
			"../assets/textures/fireball/fireball_4.png",
			"../assets/textures/fireball/fireball_5.png",
		}
	)

	for _, path := range imagesPaths {
		if file, err = os.Open(path); err != nil {
			return nil, err
		}

		img, _, err = image.Decode(file)
		file.Close()
		if err != nil {
			return nil, err
		}
		images = append(images, ebiten.NewImageFromImage(img))
		fmt.Printf("image %s\t\tdone\n", path)
	}

	pas = &PlayerAttackSystem{
		fireballImages: images,
		fireballs:      []*PlayerFireball{},
		playerX:        playerX,
		playerY:        playerY,
		playerLastMove: playerLastMove,
		block:          false,
	}

	return pas, nil
}

func (pas *PlayerAttackSystem) Handle(keys []ebiten.Key) {
	if len(pas.fireballs) != 0 {
		return
	}
	for _, key := range keys {
		if key == ebiten.KeySpace {
			pas.fireballs = append(pas.fireballs, NewPlayerFireball(
				*pas.playerX,
				*pas.playerY,
				&gm_geometry.Point{
					X: *pas.playerX + 200,
					Y: *pas.playerY,
				},
				pas.fireballImages,
				pas.RemoveFireball,
			))
		}
	}
}

func (pas *PlayerAttackSystem) RemoveFireball(fireball *PlayerFireball) {
	for i, f := range pas.fireballs {
		if f == fireball {
			pas.fireballs = append(pas.fireballs[:i], pas.fireballs[i+1:]...)
			break
		}
	}
}

func (pas *PlayerAttackSystem) GetFireballs() []*PlayerFireball {
	return pas.fireballs
}
