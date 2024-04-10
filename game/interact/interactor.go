package interactables

import (
	"fmt"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
)

func HandlePlayerInteractions(p *entity.Player) {

	interactions := findNearInteractables(p)
	handleInteractCodes(p, interactions)
}

func HandlePlayerTriggers(p *entity.Player) {

	code := level.GetTriggerAtPos(p.GetPositionX(), p.GetPositionY())

	if code == 0 {
		return
	}

	it := model.TileInteraction{
		PosX: 0,
		PosY: 0,
		Code: code,
	}

	Execute(p, it)
}

func findNearInteractables(p *entity.Player) []model.TileInteraction {

	var result []model.TileInteraction

	if t := level.GetInteractionTop(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	if t := level.GetInteractionBottom(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	if t := level.GetInteractionLeft(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	if t := level.GetInteractionRight(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	fmt.Printf("Found near interactables %v", result)

	return result
}

func handleInteractCodes(p *entity.Player, its []model.TileInteraction) {

	for _, it := range its {
		it := it
		Execute(p, it)
	}
}
