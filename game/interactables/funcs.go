package interactables

import (
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/util"
)

type InteractFunc func(ti model.TileInteraction, h *level.LevelHolder)

func interactLightSwitch(ti model.TileInteraction, h *level.LevelHolder) {

	*h.Status.LightOn = !(*h.Status.LightOn)

	if *h.Status.LightOn {
		util.SwitchLightOff(h.Level)
		h.ReplaceTile(int(ti.PosX), int(ti.PosY), 68, 67)
	} else {
		util.SwitchLightOn(h.Level)
		h.ReplaceTile(int(ti.PosX), int(ti.PosY), 67, 68)
	}
}
