package gm_machine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type GameMachine struct {
	title       string
	currentMap  int
	maps        []*gm_map.Map
	sprites     []int
	keys        []ebiten.Key
	touches     []ebiten.TouchID
	player      gm_entity.Player
	ItemStorage *gm_item.ItemCollectionStorage
}

func NewGameMachine(title string) *GameMachine {
	return &GameMachine{
		title:       title,
		maps:        []*gm_map.Map{},
		currentMap:  0,
		sprites:     []int{},
		player:      nil,
		keys:        []ebiten.Key{},
		touches:     []ebiten.TouchID{},
		ItemStorage: nil,
	}
}

func (g *GameMachine) Update() error {

	var (
		m                *gm_map.Map
		area             gm_camera.CameraArea
		playerX, playerY float64
		cameraIsMoved    bool = false
		collectedItems   []*gm_item.Item
		err              error
	)

	if g.player == nil {
		return nil
	}

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.touches = inpututil.AppendJustPressedTouchIDs(g.touches[:0])
	m = g.maps[g.currentMap]

	if m.PlayerMayMove(g.keys) {
		playerX, playerY = g.player.GetMoveSidePosition()
		area = m.GetCameraArea(playerX, playerY)
		cameraIsMoved, err = m.MoveCamera(g.keys, area)
		if err != nil {
			return err
		}
		g.player.MovementHandler(g.keys, cameraIsMoved)
	}
	g.player.MouseHandler(g.touches)
	g.player.StaffHandler(g.keys)
	g.player.AttackHandler(g.keys)
	collectedItems = g.player.CollectItemsHandler(g.maps[g.currentMap].GetDropItems())
	for _, item := range collectedItems {
		g.maps[g.currentMap].DelDropItem(item)
	}
	return nil
}

func (g *GameMachine) AddMap(m *gm_map.Map) {
	if m == nil {
		return
	}
	g.maps = append(g.maps, m)
	m.AddPlayer(g.player)
}

func (g *GameMachine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return stgs.ScreenWidth, stgs.ScreenHeight
}

func (g *GameMachine) Draw(screen *ebiten.Image) {
	g.maps[g.currentMap].Draw(screen)
}

func (g *GameMachine) Run() error {
	ebiten.SetWindowSize(stgs.ScreenWidth*3, stgs.ScreenHeight*3)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	return ebiten.RunGame(g)
}
