package gm_map

import (
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Map struct {
	ground       *Ground
	player       gm_entity.Player
	entities     []gm_entity.GameEntity
	objs         []IMapObj
	width        int
	height       int
	tileSize     int
	camera       *gm_camera.Camera
	droppedItems []*gm_item.Item
}

type MapOpt struct {
	GroundRawPath   string
	GroundImagePath string
	TileSize        int
}

func NewMap(opt MapOpt) (*Map, error) {

	var (
		width  int = 0
		height int = 0
	)

	if opt.TileSize == 0 {
		return nil, ErrMapWithZeroTileSize
	}

	if opt.GroundImagePath == "" || opt.GroundRawPath == "" {
		return nil, ErrMapWithoutBackground
	}

	background, err := gm_layer.NewLayer("background", opt.GroundRawPath, opt.GroundImagePath, opt.TileSize)
	if err != nil {
		return nil, err
	}

	height, width = background.GetSizes()

	return &Map{
		ground:       NewGround(background),
		width:        width,
		height:       height,
		player:       nil,
		tileSize:     opt.TileSize,
		objs:         []IMapObj{},
		entities:     []gm_entity.GameEntity{},
		droppedItems: []*gm_item.Item{},
		camera:       nil,
	}, nil
}

func (m *Map) AddPlayer(player gm_entity.Player) {
	if m.camera != nil {
		m.camera.SetSpeed(player.GetSpeed())
	}

	m.player = player
	m.entities = append(m.entities, player)
}

func (m *Map) AddObj(obj IMapObj) {
	m.objs = append(m.objs, obj)
}

func (m *Map) AddCamera(camera *gm_camera.Camera) {
	if camera == nil {
		return
	}

	var (
		limitX float64 = float64(m.tileSize * m.width)
		limitY float64 = float64(m.tileSize * m.height)
	)
	if m.player != nil {
		camera.SetSpeed(m.player.GetSpeed())
	}
	camera.SetLimits(limitX, limitY)
	m.camera = camera
}

func (m *Map) AddLayer(level MapLevel, layer *gm_layer.Layer) {
	if level == MapGroundLevel {
		m.ground.AddLayer(layer)
	}
}
