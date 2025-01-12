package controller

type Mover interface {
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
}
