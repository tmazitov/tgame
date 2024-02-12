package maps

import (
	"log"

	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func MainMap() (*gm_map.Map, error) {

	var (
		groundImagePath string = "assets/textures/tilesets/grass.png"
		groundRawPath   string = "maps/map1/ground_1"
		grassImagePath  string = "assets/textures/tilesets/decor_16x16.png"
		grassRawPath    string = "maps/map1/ground_2"
		roadsImagePath  string = "assets/textures/tilesets/plains.png"
		roadsRawPath    string = "maps/map1/ground_3"
		m               *gm_map.Map
		grass           *gm_layer.Layer
		roads           *gm_layer.Layer
		err             error
		size            int = stgs.TileSize
	)

	m, err = gm_map.NewMap(gm_map.MapOpt{
		GroundRawPath:   groundRawPath,
		GroundImagePath: groundImagePath,
		TileSize:        stgs.TileSize,
	})
	if err != nil {
		return nil, err
	}
	if grass, err = gm_layer.NewLayer("grass", grassRawPath, grassImagePath, size); err != nil {
		return nil, err
	}
	if roads, err = gm_layer.NewLayer("roads", roadsRawPath, roadsImagePath, size); err != nil {
		return nil, err
	}

	m.AddLayer(gm_map.MapGroundLevel, grass)
	m.AddLayer(gm_map.MapGroundLevel, roads)

	if stgs.IsDebug {
		log.Println("MainMap create\t\tsuccess")
	}

	return m, nil
}
