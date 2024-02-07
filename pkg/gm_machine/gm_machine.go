package gm_machine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tmazitov/tgame.git/internal/player"
	"github.com/tmazitov/tgame.git/pkg/gm_entity"
	gm_layer "github.com/tmazitov/tgame.git/pkg/gm_layer"
	gm_obj "github.com/tmazitov/tgame.git/pkg/gm_obj"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type GameMachine struct {
	title    string
	layers   []*gm_layer.Layer
	sprites  []int
	objs     []*gm_obj.GameObj
	keys     []ebiten.Key
	entities []gm_entity.GameEntity
	player   *player.Player
}

func NewGameMachine(title string) *GameMachine {

	var player *player.Player = player.NewPlayer(0, 0, "../assets/textures/characters/Humans_Smith.png")

	if player == nil {
		return nil
	}

	return &GameMachine{
		title:    title,
		layers:   nil,
		sprites:  []int{},
		objs:     []*gm_obj.GameObj{},
		entities: []gm_entity.GameEntity{player},
		player:   player,
		keys:     []ebiten.Key{},
	}
}

func (g *GameMachine) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.player.MovementHandler(g.keys)
	return nil
}

func (g *GameMachine) AddLayer(layer *gm_layer.Layer) {
	if layer == nil {
		return
	}
	g.layers = append(g.layers, layer)
}

func (g *GameMachine) AddObj(obj *gm_obj.GameObj) {
	if obj == nil {
		return
	}
	g.objs = append(g.objs, obj)
}

func (g *GameMachine) AddEntity(entity gm_entity.GameEntity) {
	if entity == nil {
		return
	}
	g.entities = append(g.entities, entity)
}

func (g *GameMachine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return stgs.ScreenWidth, stgs.ScreenHeight
}

func (g *GameMachine) Draw(screen *ebiten.Image) {
	for _, layer := range g.layers {
		layer.Draw(screen)
	}
	for _, obj := range g.objs {
		obj.Draw(screen)
	}
	for _, entity := range g.entities {
		entity.Draw(screen)
	}
}

func (g *GameMachine) Run() error {
	ebiten.SetWindowSize(stgs.ScreenWidth*4, stgs.ScreenHeight*4)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	return ebiten.RunGame(g)
}
