package player

import (
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

func makeMoveRightAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 16,
			TileCount:    4,
			TileSize:     32,
			TileCoords: [][2]int{
				{0, 128},
				{32, 128},
				{64, 128},
				{96, 128},
			},
		})
}

func makeMoveBotAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 16,
			TileCount:    4,
			TileSize:     32,
			TileCoords: [][2]int{
				{0, 96},
				{32, 96},
				{64, 96},
				{96, 96},
			},
		})
}

func makeMoveTopAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 16,
			TileCount:    4,
			TileSize:     32,
			TileCoords: [][2]int{
				{0, 160},
				{32, 160},
				{64, 160},
				{96, 160},
			},
		})
}
