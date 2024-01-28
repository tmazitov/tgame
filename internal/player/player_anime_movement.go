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
			TileCount:    6,
			TileSize:     48,
			TileCoords: [][2]int{
				{0, 192},
				{48, 192},
				{96, 192},
				{144, 192},
				{192, 192},
				{240, 192},
			},
		})
}

func makeMoveBotAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 16,
			TileCount:    6,
			TileSize:     48,
			TileCoords: [][2]int{
				{0, 144},
				{48, 144},
				{96, 144},
				{144, 144},
				{192, 144},
				{240, 144},
			},
		})
}

func makeMoveTopAnime(image *gm_layer.Image) *gm_anime.Anime {
	return gm_anime.NewAnime(
		gm_anime.AnimeOptions{
			TileImage:    image,
			TileLifeTime: 16,
			TileCount:    6,
			TileSize:     48,
			TileCoords: [][2]int{
				{0, 240},
				{48, 240},
				{96, 240},
				{144, 240},
				{192, 240},
				{240, 240},
			},
		})
}
