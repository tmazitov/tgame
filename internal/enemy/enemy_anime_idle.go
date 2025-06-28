package enemy

import (
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

func makeIdleBotAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 20,
			TileCount:    4,
			TileSize:     32,
			TileCoords: [][2]int{
				{0, 0},
				{32, 0},
				{64, 0},
				{96, 0},
			},
		})
}

func makeIdleTopAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 20,
			TileCount:    4,
			TileSize:     32,
			TileCoords: [][2]int{
				{0, 64},
				{32, 64},
				{64, 64},
				{96, 64},
			},
		})
}

func makeIdleRightAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 20,
			TileCount:    4,
			TileSize:     32,
			TileCoords: [][2]int{
				{0, 32},
				{32, 32},
				{64, 32},
				{96, 32},
			},
		})
}
