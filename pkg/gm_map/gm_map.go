package gm_map

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	"github.com/tmazitov/tgame.git/pkg/gm_obj"
)

type Map struct {
	ground   *Ground
	player   gm_entity.Player
	entities []gm_entity.GameEntity
	objs     []*gm_obj.GameObj
	width    int
	height   int
	raw      []int
}

type MapOpt struct {
	GroundRawPath   string
	GroundImagePath string
}

func NewMap(opt MapOpt) (*Map, error) {

	var (
		width  int   = 0
		height int   = 0
		raw    []int = []int{}
	)

	if opt.GroundImagePath == "" || opt.GroundRawPath == "" {
		return nil, ErrMapWithoutBackground
	}

	background, err := gm_layer.NewLayer("background", opt.GroundRawPath, opt.GroundImagePath)
	if err != nil {
		return nil, err
	}

	return &Map{
		ground:   NewGround(background),
		raw:      raw,
		width:    width,
		height:   height,
		player:   nil,
		objs:     []*gm_obj.GameObj{},
		entities: []gm_entity.GameEntity{},
	}, nil
}

func (m *Map) AddPlayer(player gm_entity.Player) {
	m.player = player
	m.entities = append(m.entities, player)
}

func (m *Map) AddLayer(level MapLevel, layer *gm_layer.Layer) {
	if level == MapGroundLevel {
		m.ground.AddLayer(layer)
	}
}

func (m *Map) Draw(screen *ebiten.Image) {
	m.ground.Draw(screen)

	// for _, layer := range g.layers {
	// 	layer.Draw(screen)
	// }
	// for _, obj := range g.objs {
	// 	obj.Draw(screen)
	// }
	for _, entity := range m.entities {
		entity.Draw(screen)
	}
}
