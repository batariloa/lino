package interactables

var InteractionMap map[int]InteractFunc

func init() {
	InteractionMap = make(map[int]InteractFunc)
	InteractionMap[1] = interactLightSwitch
}
