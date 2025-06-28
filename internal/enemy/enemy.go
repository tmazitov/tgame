package enemy

import (
	"fmt"
	"log"

	"github.com/tmazitov/tgame.git/internal/enemy/behavior"
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
	actionState behavior.EnemyAction
	coll        *gm_geometry.Collider
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
		actionState: behavior.EnemyStay,
		coll:        nil,
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

func (e *Enemy) GetCollider() *gm_geometry.Collider {
	return e.coll
}
