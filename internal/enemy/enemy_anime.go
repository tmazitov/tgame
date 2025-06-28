package enemy

import (
	"github.com/tmazitov/tgame.git/internal/enemy/behavior"
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type EnemyAnime struct {
	IdleBot   *gm_anime.Anime
	IdleTop   *gm_anime.Anime
	IdleRight *gm_anime.Anime
	// MoveRight *gm_anime.Anime
	// MoveBot   *gm_anime.Anime
	// MoveTop   *gm_anime.Anime
}

func NewEnemyAnime(image *gm_layer.Image) *EnemyAnime {
	return &EnemyAnime{
		IdleBot:   makeIdleBotAnime(image),
		IdleTop:   makeIdleTopAnime(image),
		IdleRight: makeIdleRightAnime(image),
		// MoveRight: makeMoveRightAnime(image),
		// MoveTop:   makeMoveTopAnime(image),
		// MoveBot:   makeMoveBotAnime(image),
	}
}

func (ea *EnemyAnime) GetCurrentAnime(action behavior.EnemyAction) *gm_anime.Anime {
	if action == behavior.EnemyStay {
		// if last == Bot_PlayerAction {
		// 	return pa.IdleBot
		// }
		// if last == Top_PlayerAction {
		// 	return pa.IdleTop
		// }
		return ea.IdleRight
	}
	// if action == Right_PlayerAction || action == Left_PlayerAction {
	// 	return pa.MoveRight
	// }
	// if action == Top_PlayerAction {
	// 	return pa.MoveTop
	// }
	// if action == Bot_PlayerAction {
	// 	return pa.MoveBot
	// }
	return (nil)
}
