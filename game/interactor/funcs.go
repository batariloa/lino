package interactor

import (
	"log"

	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/util"
)

type InteractFunc func(ti model.TileInteraction, h *level.LevelHolder)

func interactLightSwitch(ti model.TileInteraction, h *level.LevelHolder) {

	*h.Status.LightOn = !(*h.Status.LightOn)
	log.Printf("Switching light to %v", *h.Status.LightOn)

	log.Println("Right of player %d", ti.PosX)

	h.ReplaceTile(int(ti.PosX), int(ti.PosY), 67, 68)

	if *h.Status.LightOn {
		util.SwitchLightOff(h.Level)
	} else {
		util.SwitchLightOn(h.Level)
	}
}
