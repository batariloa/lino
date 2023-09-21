package interactor

import (
	"log"

	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
)

var InteractionMap map[int]InteractFunc

func init() {
	InteractionMap = make(map[int]InteractFunc)
	InteractionMap[1] = interactLightSwitch
}

func executeInteraction(ti model.TileInteraction, h *level.LevelHolder) {

	log.Printf("Trying to execute interactor with code %d", ti.Code)
	log.Println("Execute receives pos %d", ti.PosX)

	if f, ok := InteractionMap[ti.Code]; ok && f != nil {
		f(ti, h)
	}
}
