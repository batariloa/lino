package level

import "github.com/batariloa/lino/game/trigger"

func TeleportToRoom(roomTiles [][]int) {
	trigger.StartTransitionToBlack()
	LevelMap = &roomTiles
}
