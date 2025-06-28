package robber

import (
	"github.com/tmazitov/tgame.git/internal/enemy"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
)

func NewRobber(x, y float64, location *gm_map.Map) (*enemy.Enemy, error) {

	var (
		err  error
		en   *enemy.Enemy
		behavior *RobberBehavior
	)

	if en, err = enemy.NewEnemy(x, y, enemy.EnemyImagesPaths{
		Tiles:  "assets/textures/characters/Humans_Thief.png",
		Shadow: "assets/textures/characters/shadow.png",
	}); err != nil {
		return nil, err
	}

	behavior = &RobberBehavior{location: location, enemy: en}
	en.SetBehavior(behavior)

	return en, nil
}
