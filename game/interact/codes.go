package interactables

import (
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/util"
)

type InteractFunc func(ti model.TileInteraction)

var InteractionMap map[int]InteractFunc

func init() {
	InteractionMap = make(map[int]InteractFunc)
	InteractionMap[1] = interactLightSwitch
}

func interactLightSwitch(ti model.TileInteraction) {

	*level.LevelStatus.LightOn = !*level.LevelStatus.LightOn

	if *level.LevelStatus.LightOn {
		util.SwitchLightOff()
		level.ReplaceTile(int(ti.PosX), int(ti.PosY), 68, 67)
	} else {
		util.SwitchLightOn()
		level.ReplaceTile(int(ti.PosX), int(ti.PosY), 67, 68)
	}
}
