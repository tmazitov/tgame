package enemy

func (e *Enemy) move() {
	if e.behavior == nil {
		return
	}

	vectorX, vectorY, futureTarget := e.behavior.Move()
	if vectorX == 0 && vectorY == 0 {
		e.idle(futureTarget)
		return
	}

	if vectorY < 0 {
		e.moveTop(vectorY)
	} else if vectorY > 0 {
		e.moveBot(vectorY)
	}

	if vectorX > 0 {
		e.moveRight(vectorX)
	} else if vectorX < 0 {
		e.moveLeft(vectorX)
	}
}

func (e *Enemy) idle(futureTarget EnemyAction) {
	if e.actionState == Idle_EnemyAction {
		return
	}
	if futureTarget != Idle_EnemyAction {
		e.lastAction = futureTarget
	} else {
		e.lastAction = e.actionState
	}
	e.actionState = Idle_EnemyAction
}

func (e *Enemy) moveLeft(speed float64) {
	e.X += speed
	if e.actionState != Left_EnemyAction {
		e.lastAction = e.actionState
		e.actionState = Left_EnemyAction
	}
}
func (e *Enemy) moveRight(speed float64) {
	e.X += speed
	if e.actionState != Right_EnemyAction {
		e.lastAction = e.actionState
		e.actionState = Right_EnemyAction
	}
}
func (e *Enemy) moveTop(speed float64) {
	e.Y += speed
	if e.actionState != Top_EnemyAction {
		e.lastAction = e.actionState
		e.actionState = Top_EnemyAction
	}
}
func (e *Enemy) moveBot(speed float64) {
	e.Y += speed
	if e.actionState != Bot_EnemyAction {
		e.lastAction = e.actionState
		e.actionState = Bot_EnemyAction
	}
}
