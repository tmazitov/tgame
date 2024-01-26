package gm_machine

import (
	"github.com/hajimehoshi/ebiten/v2"
	gm_layer "github.com/tmazitov/tgame.git/pkg/gm_layer"
	gm_obj "github.com/tmazitov/tgame.git/pkg/gm_obj"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type GameMachine struct {
	title   string
	layers  []*gm_layer.Layer
	sprites []int
	objs    []*gm_obj.GameObj
}

func NewGameMachine(title string) *GameMachine {
	return &GameMachine{
		title:   title,
		layers:  []*gm_layer.Layer{},
		sprites: []int{},
		objs:    []*gm_obj.GameObj{},
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
}

func (g *GameMachine) Run() error {
	ebiten.SetWindowSize(stgs.ScreenWidth*2, stgs.ScreenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	return ebiten.RunGame(g)
}
