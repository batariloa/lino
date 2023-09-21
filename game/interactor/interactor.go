package interactor

import (
	"fmt"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/view"
)

func HandlePlayerInteractions(p *entity.Player, h *level.LevelHolder) {

	interactions := findNearInteractables(p, h)
	handleInteractCodes(&interactions, h)
}

func findNearInteractables(p *entity.Player, h *level.LevelHolder) []int {

	var result []int

	if t := h.GetInteractionTop(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	if t := h.GetInteractionBottom(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	if t := h.GetInteractionLeft(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	if t := h.GetInteractionRight(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	fmt.Printf("Found near interactables %d", result)

	return result
}

func handleInteractCodes(c *[]int, h *level.LevelHolder) {

	for _, code := range *c {
		executeInteraction(code, h)
	}
}
