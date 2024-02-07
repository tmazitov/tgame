package maps

import (
	"log"

	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func MainMap() (*gm_map.Map, error) {

	var (
		groundImagePath string = "../assets/textures/tilesets/grass.png"
		groundRawPath   string = "../maps/map1/ground_1"
		grassImagePath  string = "../assets/textures/tilesets/decor_16x16.png"
		grassRawPath    string = "../maps/map1/ground_2"
		m               *gm_map.Map
		grass           *gm_layer.Layer
		err             error
	)

	m, err = gm_map.NewMap(gm_map.MapOpt{
		GroundRawPath:   groundRawPath,
		GroundImagePath: groundImagePath,
	})
	if err != nil {
		return nil, err
	}
	if grass, err = gm_layer.NewLayer("grass", grassRawPath, grassImagePath); err != nil {
		return nil, err
	}

	m.AddLayer(gm_map.MapGroundLevel, grass)

	if stgs.IsDebug {
		log.Println("MainMap create\t\tsuccess")
	}

	return m, nil
}
