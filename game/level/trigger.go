package level

import "github.com/batariloa/lino/game/view"

func init() {
	Triggers = map[int]func(){}
}

func calcTriggerIndex(x, y int) int {

	width := MaxLevelWidth / view.TileSize
	index := y*width + x

	return index
}

func addTrigger(x, y int, triggerFunc func()) {

	index := calcTriggerIndex(x, y)
	Triggers[index] = triggerFunc
}
