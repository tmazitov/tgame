package scene

import "github.com/tmazitov/tgame.git/internal/enemy"

type SceneEnemiesStorage struct {
	enemies []*enemy.Enemy
}

func NewSceneEnemiesStorage() *SceneEnemiesStorage {
	return &SceneEnemiesStorage{
		enemies: []*enemy.Enemy{},
	}
}

func (es *SceneEnemiesStorage) Add(en *enemy.Enemy) {
	es.enemies = append(es.enemies, en)
}

func (es *SceneEnemiesStorage) Get() []*enemy.Enemy {
	return es.enemies
}
