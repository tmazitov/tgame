package robber

import (
	"fmt"
	"math"

	"github.com/tmazitov/tgame.git/internal/enemy"
	"github.com/tmazitov/tgame.git/pkg/gm_map"
	stgs "github.com/tmazitov/tgame.git/settings"
)

type RobberBehavior struct {
	location *gm_map.Map
	enemy    *enemy.Enemy
}

func calcFutureTarget(differenceX, differenceY float64) enemy.EnemyAction {
	fmt.Printf("diff: %v %v\n", differenceX, differenceY)
	if differenceX > 1 {
		return enemy.Left_EnemyAction
	} else if differenceX < -1 {
		return enemy.Right_EnemyAction
	}

	if differenceY > 1 {
		return enemy.Top_EnemyAction
	} else if differenceY < 1 {
		return enemy.Bot_EnemyAction
	}

	return enemy.Idle_EnemyAction
}

func (b *RobberBehavior) Move() (float64, float64, enemy.EnemyAction) {

	currentX, currentY := b.enemy.GetPosition()
	currentSpeed := b.enemy.GetSpeed()

	player := b.location.GetPlayer()
	playerX, playerY := player.GetOppositeMoveSidePosition()

	differenceX := currentX - playerX
	differenceY := currentY - playerY

	if differenceX >= 1 {
		differenceX -= stgs.PlayerSize
	} else if differenceX <= -1 {
		differenceX += stgs.PlayerSize
	} else {
		differenceX = 0
	}

	if differenceY > 1 {
		differenceY += stgs.PlayerSize
	} else if differenceY < -1 {
		differenceY -= stgs.PlayerSize
	} else {
		differenceY = 0
	}

	differenceX *= -1
	differenceY *= -1

	vectorX := math.Max(math.Min(differenceX, currentSpeed), currentSpeed*-1)
	vectorY := math.Max(math.Min(differenceY, currentSpeed), currentSpeed*-1)

	fmt.Printf("robber move: %v %v | pos : %v %v | diff : %v %v\n", vectorX, vectorY, playerX, currentX, differenceX, differenceY)
	if !b.location.EntityMayMove(b.enemy, vectorX, vectorY) {
		return 0, 0, calcFutureTarget(currentX-playerX, currentY-playerY)
	}

	return vectorX, vectorY, calcFutureTarget(currentX-playerX, currentY-playerY)
}
