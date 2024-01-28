package player

import (
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

func makeIdleBotAnime(image *gm_layer.Image) *gm_anime.Anime {
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

func makeIdleTopAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 20,
			TileCount:    6,
			TileSize:     48,
			TileCoords: [][2]int{
				{0, 96},
				{48, 96},
				{96, 96},
				{144, 96},
				{192, 96},
				{240, 96},
			},
		})
}

func makeIdleRightAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 20,
			TileCount:    6,
			TileSize:     48,
			TileCoords: [][2]int{
				{0, 48},
				{48, 48},
				{96, 48},
				{144, 48},
				{192, 48},
				{240, 48},
			},
		})
}
