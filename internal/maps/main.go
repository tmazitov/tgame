package maps

import (
	"log"

	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

func MainMap() (*gm_map.Map, error) {

	var (
		groundImagePath string = "../assets/textures/tilesets/grass.png"
		groundRawPath   string = "../maps/map1/ground_1"
		m               *gm_map.Map
		err             error
	)

	m, err = gm_map.NewMap(gm_map.MapOpt{
		GroundRawPath:   groundRawPath,
		GroundImagePath: groundImagePath,
	})
	if err != nil {
		return nil, err
	}

	if stgs.IsDebug {
		log.Println("MainMap create\t\tsuccess")
	}

	return m, nil
}
