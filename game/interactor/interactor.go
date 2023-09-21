package interactor

import (
	"fmt"
	"log"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
)

func HandlePlayerInteractions(p *entity.Player, h *level.LevelHolder) {

	interactions := findNearInteractables(p, h)
	handleInteractCodes(interactions, h)
}

func findNearInteractables(p *entity.Player, h *level.LevelHolder) []model.TileInteraction {

	var result []model.TileInteraction

	if t := h.GetInteractionTop(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	if t := h.GetInteractionBottom(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	if t := h.GetInteractionLeft(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	if t := h.GetInteractionRight(p, view.TileSize); true {
		log.Println("Checking interaction right %d", t.PosX)
		result = append(result, *t)
	}

	fmt.Printf("Found near interactables %s", result)

	return result
}

func handleInteractCodes(its []model.TileInteraction, h *level.LevelHolder) {

	for _, it := range its {
		log.Println("Before execute", it.PosX)
		executeInteraction(it, h)
	}
}
