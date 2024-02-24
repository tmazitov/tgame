package gm_machine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	if g.player == nil {
		return nil
	}

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.touches = inpututil.AppendJustPressedTouchIDs(g.touches[:0])

	return g.CurrentMap().Update(g.touches, g.keys)
}

func (g *GameMachine) CurrentMap() *gm_map.Map {
	return g.maps[g.currentMap]
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
	g.CurrentMap().Draw(screen)
}

func (g *GameMachine) Run() error {
	ebiten.SetWindowSize(stgs.ScreenWidth*3, stgs.ScreenHeight*3)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	return ebiten.RunGame(g)
}
