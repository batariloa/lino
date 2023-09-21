package interactor

import "github.com/batariloa/lino/game/level"

var InteractionMap map[int]InteractFunc

func init() {
	InteractionMap = make(map[int]InteractFunc)
	InteractionMap[1] = interactLightSwitch
}

func executeInteraction(code int, h *level.LevelHolder) {
	InteractionMap[code](h)
}
