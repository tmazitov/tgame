package gm_map

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	tileSize int
	camera   *Camera
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
		ground:   NewGround(background),
		width:    width,
		height:   height,
		player:   nil,
		tileSize: opt.TileSize,
		objs:     []*gm_obj.GameObj{},
		entities: []gm_entity.GameEntity{},
		camera:   nil,
	}, nil
}

func (m *Map) AddPlayer(player gm_entity.Player) {
	if m.camera != nil {
		m.camera.SetSpeed(player.GetSpeed())
	}

	m.player = player
	m.entities = append(m.entities, player)
}

func (m *Map) AddCamera(camera *Camera) {
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

func (m *Map) MoveCamera(keys []ebiten.Key) (bool, error) {
	return m.camera.MovementHandler(keys)
}

func (m *Map) Draw(screen *ebiten.Image) {

	var border gm_layer.LayerBorder = gm_layer.LayerBorder{
		X:      m.camera.X,
		Y:      m.camera.Y,
		Width:  m.camera.width,
		Height: m.camera.height,
	}

	m.ground.Draw(screen, border)

	// for _, layer := range g.layers {
	// 	layer.Draw(screen)
	// }
	// for _, obj := range g.objs {
	// 	obj.Draw(screen)
	// }
	for _, entity := range m.entities {
		entity.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("camera %f %f %f \n", m.camera.X, m.camera.Y, m.camera.limitX))
}
