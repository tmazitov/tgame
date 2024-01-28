package player

import (
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type PlayerAnime struct {
	IdleBot   *gm_anime.Anime
	IdleTop   *gm_anime.Anime
	IdleRight *gm_anime.Anime
	MoveRight *gm_anime.Anime
	MoveBot   *gm_anime.Anime
	MoveTop   *gm_anime.Anime
}

func NewPlayerAnime(image *gm_layer.Image) *PlayerAnime {
	return &PlayerAnime{
		IdleBot:   makeIdleBotAnime(image),
		IdleTop:   makeIdleTopAnime(image),
		IdleRight: makeIdleRightAnime(image),
		MoveRight: makeMoveRightAnime(image),
		MoveTop:   makeMoveTopAnime(image),
		MoveBot:   makeMoveBotAnime(image),
	}
}

func (pa *PlayerAnime) GetCurrentAnime(action PlayerAction, last PlayerAction) *gm_anime.Anime {
	if action == Idle_PlayerAction {
		if last == Bot_PlayerAction {
			return pa.IdleBot
		}
		if last == Top_PlayerAction {
			return pa.IdleTop
		}
		return pa.IdleRight
	}
	if action == Right_PlayerAction || action == Left_PlayerAction {
		return pa.MoveRight
	}
	if action == Top_PlayerAction {
		return pa.MoveTop
	}
	if action == Bot_PlayerAction {
		return pa.MoveBot
	}
	return (nil)
}
