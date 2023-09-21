package interactor

import (
	"log"

	"github.com/batariloa/lino/game/level"
)

var InteractionMap map[int]InteractFunc

func init() {
	InteractionMap = make(map[int]InteractFunc)
	InteractionMap[1] = interactLightSwitch
}

func executeInteraction(code int, h *level.LevelHolder) {

	log.Printf("Trying to execute interactor with code %d", code)

	if f, ok := InteractionMap[code]; ok && f != nil {
		f(h)
	}
}
