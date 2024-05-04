package util

import (
	"github.com/batariloa/lino/game/level"
)

func SwitchLightOn() {

	size := len((*level.LevelMap)[0])
	newLayer := make([]int, size)

	for i := range (*level.LevelMap)[0] {
		//dark shade
		newLayer[i] = 399
	}

	*level.LevelMap = append(*level.LevelMap, newLayer)
}

func SwitchLightOff() {

	if len(*level.LevelMap) > 0 {
		*level.LevelMap = (*level.LevelMap)[:len(*level.LevelMap)-1]
	}
}
