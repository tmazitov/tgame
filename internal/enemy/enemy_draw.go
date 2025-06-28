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
	var anime *gm_anime.Anime = e.anime.GetCurrentAnime(e.actionState, e.lastAction)
	if anime == nil {
		return (e.anime.IdleBot.GetTile())
	}
	return anime.GetTile()
}

func FlipVertical(source *ebiten.Image) *ebiten.Image {
	result := ebiten.NewImage(source.Bounds().Dx(), source.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(result.Bounds().Dx()), 0)
	result.DrawImage(source, op)
	return result
}

func (p *Enemy) Draw(screen *ebiten.Image, camera *gm_camera.Camera) {

	var (
		relativeX, relativeY float64
		tile                 *ebiten.Image = p.GetNextTile()
	)

	relativeX, relativeY, _ = camera.GetRelativeCoords(p.X, p.Y)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	if p.actionState == Left_EnemyAction || p.lastAction == Left_EnemyAction {
		tile = FlipVertical(tile)
	}

	screen.DrawImage(tile, op)
	p.drawShadow(screen, relativeX, relativeY)
}
