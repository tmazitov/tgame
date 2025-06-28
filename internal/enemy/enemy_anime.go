package enemy

import (
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type EnemyAnime struct {
	IdleBot   *gm_anime.Anime
	IdleTop   *gm_anime.Anime
	IdleRight *gm_anime.Anime
	MoveRight *gm_anime.Anime
	MoveBot   *gm_anime.Anime
	MoveTop   *gm_anime.Anime
}

func NewEnemyAnime(image *gm_layer.Image) *EnemyAnime {
	return &EnemyAnime{
		IdleBot:   makeIdleBotAnime(image),
		IdleTop:   makeIdleTopAnime(image),
		IdleRight: makeIdleRightAnime(image),
		MoveRight: makeMoveRightAnime(image),
		MoveTop:   makeMoveTopAnime(image),
		MoveBot:   makeMoveBotAnime(image),
	}
}

func (ea *EnemyAnime) GetCurrentAnime(action EnemyAction, last EnemyAction) *gm_anime.Anime {
	if action == Idle_EnemyAction {
		if last == Bot_EnemyAction {
			return ea.IdleBot
		}
		if last == Top_EnemyAction {
			return ea.IdleTop
		}
		return ea.IdleRight
	}
	if action == Right_EnemyAction || action == Left_EnemyAction {
		return ea.MoveRight
	}
	if action == Top_EnemyAction {
		return ea.MoveTop
	}
	if action == Bot_EnemyAction {
		return ea.MoveBot
	}
	return (nil)
}
