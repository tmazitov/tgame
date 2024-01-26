package ground

import stgs "github.com/tmazitov/tgame.git/settings"

func GroundRaw(tileNumber int) []int {
	var ground []int = []int{}

	for h := 0; h < stgs.ScreenHeight; h += stgs.TileSize {
		for w := 0; w < stgs.ScreenWidth; w += stgs.TileSize {
			ground = append(ground, tileNumber)
		}
	}
	return ground
}
