package enemy

type EnemyAction int

const (
	Idle_EnemyAction EnemyAction = iota
	Left_EnemyAction
	Right_EnemyAction
	Top_EnemyAction
	Bot_EnemyAction
)

type IEnemyBehavior interface {
	Move() (float64, float64, EnemyAction)
}
