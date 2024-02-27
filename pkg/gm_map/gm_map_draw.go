package gm_map

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
)

func (m *Map) Draw(screen *ebiten.Image) {

	var border gm_layer.LayerBorder = gm_layer.LayerBorder{
		X:      m.camera.X,
		Y:      m.camera.Y,
		Width:  m.camera.Width,
		Height: m.camera.Height,
	}

	m.ground.Draw(screen, border)

	for _, item := range m.droppedItems {
		if item.InDropProcess() {
			continue
		}
		item.Draw(screen, m.camera)
	}

	for _, entity := range m.entities {
		entity.Draw(screen, m.camera)
	}

	for _, item := range m.droppedItems {
		if item.InDropProcess() {
			item.Draw(screen, m.camera)
		}
	}

	for _, obj := range m.objs {
		obj.Draw(screen, m.camera)
	}

	// posX, posY := m.player.GetPosition()
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Player: %f %f | %f %f \n", posX, posY, m.camera.X, m.camera.Y))
}
