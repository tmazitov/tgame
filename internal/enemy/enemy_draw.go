package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
)

func (e *Enemy) drawShadow(screen *ebiten.Image, relativeX, relativeY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	screen.DrawImage(e.images.Shadow.Inst, op)
}

func (e *Enemy) GetNextTile() *ebiten.Image {
	var anime *gm_anime.Anime = e.anime.GetCurrentAnime(e.actionState)
	if anime == nil {
		return (e.anime.IdleBot.GetTile())
	}
	return anime.GetTile()
}

func (p *Enemy) Draw(screen *ebiten.Image, camera *gm_camera.Camera) {

	var (
		relativeX, relativeY float64
		tile                 *ebiten.Image = p.GetNextTile()
	)

	relativeX, relativeY, _ = camera.GetRelativeCoords(p.X, p.Y)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	// if p.actionState == Left_PlayerAction || p.lastAction == Left_PlayerAction {
	// 	tile = FlipVertical(tile)
	// }

	screen.DrawImage(tile, op)
	p.drawShadow(screen, relativeX, relativeY)
}
