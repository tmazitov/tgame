package gm_map

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

type Map struct {
	ground   *Ground
	player   gm_entity.Player
	entities []gm_entity.GameEntity
	objs     []IMapObj
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
		objs:     []IMapObj{},
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

func (m *Map) AddObj(obj IMapObj) {
	m.objs = append(m.objs, obj)
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

func (m *Map) MoveCamera(keys []ebiten.Key, area CameraArea) (bool, error) {
	return m.camera.MovementHandler(keys, area)
}

func (m *Map) GetCameraArea(x, y float64) CameraArea {
	return m.camera.GetPointArea(x, y)
}

func (m *Map) Draw(screen *ebiten.Image) {

	var border gm_layer.LayerBorder = gm_layer.LayerBorder{
		X:      m.camera.X,
		Y:      m.camera.Y,
		Width:  m.camera.Width,
		Height: m.camera.Height,
	}

	m.ground.Draw(screen, border)

	// for _, layer := range g.layers {
	// 	layer.Draw(screen)
	// }

	for _, entity := range m.entities {
		entity.Draw(screen)
	}

	for _, obj := range m.objs {
		obj.Draw(screen, m.camera)
	}

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("area %d\n", m.GetCameraArea(m.player.GetPosition())))
	// relX, relY, inCamera := m.camera.GetRelativeCoords(m.player.GetPosition())
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("area %f %f %t \n", relX, relY, inCamera), 300, 0)
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("camera %f %f %f \n", m.camera.X, m.camera.Y, m.camera.limitX), 0, 30)
}
