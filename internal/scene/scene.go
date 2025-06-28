package scene

import (
	"github.com/tmazitov/tgame.git/internal/enemy"
	"github.com/tmazitov/tgame.git/internal/items"
	"github.com/tmazitov/tgame.git/internal/maps"
	"github.com/tmazitov/tgame.git/internal/player"
	"github.com/tmazitov/tgame.git/pkg/gm_camera"
	"github.com/tmazitov/tgame.git/pkg/gm_font"
	"github.com/tmazitov/tgame.git/pkg/gm_item"
	"github.com/tmazitov/tgame.git/pkg/gm_machine"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type Scene struct {
	location *gm_map.Map
	game     *gm_machine.GameMachine
	font     *gm_font.Font
	enemies  *SceneEnemiesStorage
	player   *player.Player
}

func NewScene(game *gm_machine.GameMachine, font *gm_font.Font) (*Scene, error) {

	var (
		err      error
		location *gm_map.Map
		scene    *Scene
		enemies  *SceneEnemiesStorage
	)

	if location, err = maps.MainMap(); err != nil {
		return nil, err
	}

	enemies = NewSceneEnemiesStorage()

	scene = &Scene{
		location: location,
		game:     game,
		font:     font,
		enemies:  enemies,
		player:   nil,
	}

	return scene, nil
}

func (s *Scene) Load() error {
	var (
		err                   error
		itemCollectionStorage *gm_item.ItemCollectionStorage
	)

	s.location.AddCamera(gm_camera.NewCamera(stgs.ScreenHeight, stgs.ScreenWidth))
	s.player, err = player.NewPlayer(0, 0, player.PlayerImagesPaths{
		Tiles:  "assets/textures/characters/Humans_Smith.png",
		Shadow: "assets/textures/characters/shadow.png",
	}, s.font)

	if err != nil {
		return err
	}

	itemCollectionStorage, err = gm_item.NewItemCollectionStorage("items/collectionsConfig.json", 32, s.font)
	if err != nil {
		return err
	}

	s.game.SetupItemStorage(itemCollectionStorage)
	s.game.SetupPlayer(s.player)

	item := s.game.ItemStorage.GetItem(items.MaterialsCollection, items.Stick).Clone(5)
	item.AutoDrop(10, 25)
	s.location.AddDropItem(item)

	item = s.game.ItemStorage.GetItem(items.FoodCollection, items.CherryPie).Clone(3)
	item.AutoDrop(70, 15)
	s.location.AddDropItem(item)

	item = s.game.ItemStorage.GetItem(items.FoodCollection, items.Tomato).Clone(4)
	item.AutoDrop(30, 55)
	s.location.AddDropItem(item)

	s.game.AddMap(s.location)

	e, err := enemy.NewEnemy(50, 50, enemy.EnemyImagesPaths{
		Tiles:  "assets/textures/characters/Humans_Thief.png",
		Shadow: "assets/textures/characters/shadow.png",
	})
	if err != nil {
		return err
	}
	s.enemies.Add(e)

	for _, enemyItem := range s.enemies.Get() {
		s.location.AddEntity(enemyItem)
	}

	return nil
}
