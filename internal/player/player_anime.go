package player

import (
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type PlayerAnime struct {
	Idle *gm_anime.Anime
}

func MakeIdleAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 20,
			TileCount:    6,
			TileSize:     48,
			TileCoords: [][2]int{
				{0, 0},
				{48, 0},
				{96, 0},
				{144, 0},
				{192, 0},
				{240, 0},
			},
		})
}

func NewPlayerAnime(image *gm_layer.Image) *PlayerAnime {
	return &PlayerAnime{
		Idle: MakeIdleAnime(image),
	}
}

func (pa *PlayerAnime) GetCurrentAnime(action PlayerAction) *gm_anime.Anime {
	if action == Idle_PlayerAction {
		return pa.Idle
	}
	return (nil)
}
