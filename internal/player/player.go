package player

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmazitov/tgame.git/pkg/gm_anime"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_inventory"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type PlayerImagesPaths struct {
	Tiles  string
	Shadow string
}

type PlayerImages struct {
	Tiles  *gm_layer.Image
	Shadow *gm_layer.Image
}

type Player struct {
	X           float64
	Y           float64
	Speed       float64
	coll        *gm_geometry.Collider
	anime       *PlayerAnime
	images      *PlayerImages
	lastAction  PlayerAction
	actionState PlayerAction
	attack      *PlayerAttackSystem
	inventory   *PlayerInventory
}

func NewPlayer(x, y float64, imagesPaths PlayerImagesPaths, font *gm_font.Font) (*Player, error) {

	var (
		err         error
		playerAnime *PlayerAnime
		pl          *Player
	)

	pl = &Player{
		X:           x,
		Y:           y,
		Speed:       stgs.PlayerSpeed,
		images:      &PlayerImages{},
		anime:       playerAnime,
		actionState: Idle_PlayerAction,
		attack:      nil,
		coll:        nil,
	}

	if pl.images.Tiles, err = gm_layer.NewImageByPath(imagesPaths.Tiles, stgs.TileSize); err != nil {
		return nil, err
	}

	if pl.images.Shadow, err = gm_layer.NewImageByPath(imagesPaths.Shadow, stgs.TileSize); err != nil {
		return nil, err
	}

	if pl.anime = NewPlayerAnime(pl.images.Tiles); err != nil {
		return nil, err
	}

	if pl.attack, err = NewPlayerAttackSystem(&pl.X, &pl.Y, &pl.lastAction); pl.attack == nil {
		return nil, err
	}

	if pl.inventory, err = NewPlayerInventory(stgs.ScreenWidth-226, 30, font); err != nil {
		return nil, err
	}

	if stgs.IsDebug {
		log.Println("Player create\t\tsuccess")
	}

	pl.coll = gm_geometry.NewCollider(&pl.X, &pl.Y, gm_geometry.ColliderOptions{
		Height:      16,
		Width:       16,
		PaddingTop:  8,
		PaddingLeft: 8,
	})

	return pl, nil
}

func (p *Player) GetNextTile() *ebiten.Image {
	var anime *gm_anime.Anime = p.anime.GetCurrentAnime(p.actionState, p.lastAction)
	if anime == nil {
		return (p.anime.IdleBot.GetTile())
	}
	return anime.GetTile()
}

func (p *Player) GetCollider() *gm_geometry.Collider {
	return p.coll
}

func (p *Player) GetInventory() *gm_inventory.Inventory {
	return p.inventory.inventory
}

func (p *Player) GetMoveVector(keys []ebiten.Key) (float64, float64) {

	var (
		moveTop         bool
		moveBot         bool
		moveLeft        bool
		moveRight       bool
		pressedKeyCount int     = 0
		vectorX         float64 = 0
		vectorY         float64 = 0
	)

	for _, key := range keys {
		if key == ebiten.KeyW {
			moveTop = true
			pressedKeyCount++
		}
		if key == ebiten.KeyA {
			moveLeft = true
			pressedKeyCount++
		}
		if key == ebiten.KeyS {
			moveBot = true
			pressedKeyCount++
		}
		if key == ebiten.KeyD {
			moveRight = true
			pressedKeyCount++
		}
	}

	if moveTop && moveBot {
		moveTop = false
		moveBot = false
		pressedKeyCount--
	}

	if moveLeft && moveRight {
		moveLeft = false
		moveRight = false
		pressedKeyCount--
	}

	if pressedKeyCount == 1 && moveTop {
		return 0, -p.Speed
	}
	if pressedKeyCount == 1 && moveBot {
		return 0, p.Speed
	}
	if pressedKeyCount == 1 && moveLeft {
		return -p.Speed, 0
	}
	if pressedKeyCount == 1 && moveRight {
		return p.Speed, 0
	}

	if pressedKeyCount == 2 {
		if moveTop {
			vectorY -= p.Speed / math.Sqrt2
		}
		if moveBot {
			vectorY += p.Speed / math.Sqrt2
		}
		if moveLeft {
			vectorX -= p.Speed / math.Sqrt2
		}
		if moveRight {
			vectorX += p.Speed / math.Sqrt2
		}
	}

	return vectorX, vectorY
}

func (p *Player) GetSpeed() *float64 {
	return &p.Speed
}

func (p *Player) GetPosition() (float64, float64) {
	return p.X, p.Y
}

func (p *Player) GetMoveSidePosition() (float64, float64) {
	if (p.actionState == Left_PlayerAction) || (p.lastAction == Left_PlayerAction) {
		return p.X, p.Y
	}
	if (p.actionState == Right_PlayerAction) || (p.lastAction == Right_PlayerAction) {
		return p.X + stgs.PlayerSize, p.Y
	}
	if (p.actionState == Top_PlayerAction) || (p.lastAction == Top_PlayerAction) {
		return p.X, p.Y
	}
	if (p.actionState == Bot_PlayerAction) || (p.lastAction == Bot_PlayerAction) {
		return p.X, p.Y + stgs.PlayerSize
	}
	return p.X, p.Y
}

func FlipVertical(source *ebiten.Image) *ebiten.Image {
	result := ebiten.NewImage(source.Bounds().Dx(), source.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(result.Bounds().Dx()), 0)
	result.DrawImage(source, op)
	return result
}

func (p *Player) drawShadow(screen *ebiten.Image, relativeX, relativeY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	screen.DrawImage(p.images.Shadow.Inst, op)
}

func (p *Player) Draw(screen *ebiten.Image, camera *gm_camera.Camera) {

	var (
		relativeX, relativeY float64
		tile                 *ebiten.Image = p.GetNextTile()
	)

	relativeX, relativeY, _ = camera.GetRelativeCoords(p.X, p.Y)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(relativeX, relativeY)
	if p.actionState == Left_PlayerAction || p.lastAction == Left_PlayerAction {
		tile = FlipVertical(tile)
	}

	screen.DrawImage(tile, op)
	p.drawShadow(screen, relativeX, relativeY)

	for _, fireball := range p.attack.GetFireballs() {
		fireball.Move()
		fireball.Draw(screen)
	}
	p.inventory.Draw(screen)
}

func (p *Player) StaffHandler(keys []ebiten.Key) {
	p.inventory.HandleToggle(keys)
}
