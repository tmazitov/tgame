package gm_machine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type GameMachine struct {
	title      string
	currentMap int
	maps       []*gm_map.Map
	sprites    []int
	keys       []ebiten.Key
	player     gm_entity.Player
}

func NewGameMachine(title string) *GameMachine {
	return &GameMachine{
		title:      title,
		maps:       []*gm_map.Map{},
		currentMap: 0,
		sprites:    []int{},
		player:     nil,
		keys:       []ebiten.Key{},
	}
}

func (g *GameMachine) Update() error {
	if g.player == nil {
		return nil
	}

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	cameraIsMoved, err := g.maps[g.currentMap].MoveCamera(g.keys)
	if err != nil {
		return err
	}
	g.player.MovementHandler(g.keys, cameraIsMoved)
	g.player.AttackHandler(g.keys)
	return nil
}

func (g *GameMachine) AddPlayer(player gm_entity.Player) {
	if player == nil {
		return
	}
	g.player = player
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
