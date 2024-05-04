package level

type LevelData struct {
	Level *[][]int
}

type Level interface {
	GenerateMap()
	MoveRoom()
}
