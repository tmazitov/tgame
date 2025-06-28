package enemy

import (
	"fmt"
	"log"

	"github.com/tmazitov/tgame.git/pkg/gm_geometry"
	"github.com/tmazitov/tgame.git/pkg/gm_layer"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type EnemyImages struct {
	Tiles  *gm_layer.Image
	Shadow *gm_layer.Image
}

type Enemy struct {
	X           float64
	Y           float64
	Speed       float64
	images      *EnemyImages
	anime       *EnemyAnime
	actionState EnemyAction
	lastAction  EnemyAction
	coll        *gm_geometry.Collider
	behavior    IEnemyBehavior
}

type EnemyImagesPaths struct {
	Tiles  string
	Shadow string
}

func NewEnemy(x, y float64, imagesPaths EnemyImagesPaths) (*Enemy, error) {

	var (
		err        error
		en         *Enemy
		enemyAnime *EnemyAnime
	)

	en = &Enemy{
		X:           x,
		Y:           y,
		Speed:       stgs.EnemySpeed,
		images:      &EnemyImages{},
		anime:       enemyAnime,
		actionState: Idle_EnemyAction,
		coll:        nil,
		behavior:    nil,
	}

	if en.images.Tiles, err = gm_layer.NewImageByPath(imagesPaths.Tiles, stgs.TileSize); err != nil {
		return nil, err
	}

	if en.images.Shadow, err = gm_layer.NewImageByPath(imagesPaths.Shadow, stgs.TileSize); err != nil {
		return nil, err
	}

	en.anime = NewEnemyAnime(en.images.Tiles)

	if stgs.IsDebug {
		log.Println("enemy create\t\tsuccess")
	}

	en.coll = gm_geometry.NewCollider(&en.X, &en.Y, gm_geometry.ColliderOptions{
		Height:      16,
		Width:       16,
		PaddingTop:  8,
		PaddingLeft: 8,
	})

	fmt.Printf("collider %v %v\n", *en.coll.X, *en.coll.Y)

	return en, nil
}

func (e *Enemy) Update() {
	e.move()
}

func (e *Enemy) GetCollider() *gm_geometry.Collider {
	return e.coll
}

func (e *Enemy) SetBehavior(behavior IEnemyBehavior) {
	e.behavior = behavior
}

func (e *Enemy) GetPosition() (float64, float64) {
	return e.X, e.Y
}
func (e *Enemy) GetSpeed() float64 {
	return e.Speed
}
