package gm_machine

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	entities []gm_entity.GameEntity
}

func NewGameMachine(title string) *GameMachine {
	return &GameMachine{
		title:    title,
		layers:   []*gm_layer.Layer{},
		sprites:  []int{},
		objs:     []*gm_obj.GameObj{},
		entities: []gm_entity.GameEntity{},
	}
}

func (g *GameMachine) Update() error {
	return nil
}

func (g *GameMachine) AddLayer(layer *gm_layer.Layer) {
	g.layers = append(g.layers, layer)
}

func (g *GameMachine) AddObj(obj *gm_obj.GameObj) {
	g.objs = append(g.objs, obj)
}

func (g *GameMachine) AddEntity(entity gm_entity.GameEntity) {
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
	ebiten.SetWindowSize(stgs.ScreenWidth*3, stgs.ScreenHeight*3)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	return ebiten.RunGame(g)
}
