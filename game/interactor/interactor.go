package interactor

import (
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

	if t := h.GetTileTop(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	if t := h.GetTileBottom(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	if t := h.GetTileLeft(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	if t := h.GetTileRight(p, view.TileSize); t != 0 {
		result = append(result, t)
	}

	return result
}

func handleInteractCodes(c *[]int, h *level.LevelHolder) {

	for code := range *c {
		executeInteraction(code, h)
	}
}
