package interactor

import (
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/util"
)

type InteractFunc func(h *level.LevelHolder)

func interactLightSwitch(h *level.LevelHolder) {

	h.Status.LightOn = !h.Status.LightOn

	if h.Status.LightOn {
		util.SwitchLightOff(&h.Level)
	} else {
		util.SwitchLightOn(&h.Level)
	}
}
