package interactables

import (
	"fmt"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
)

type Interactor struct {
	executor *InteractableExecutor
}

func NewInteractor(executor *InteractableExecutor) *Interactor {
	return &Interactor{
		executor: executor,
	}
}

func (i *Interactor) HandlePlayerInteractions(p *entity.Player, h *level.LevelHolder) {

	interactions := i.findNearInteractables(p, h)
	i.handleInteractCodes(h, p, interactions)
}

func (i *Interactor) HandlePlayerTriggers(h *level.LevelHolder, p *entity.Player) {

	code := h.GetTriggerAtPos(p.GetPositionX(), p.GetPositionY())

	if code == 0 {
		return
	}

	it := model.TileInteraction{
		PosX: 0,
		PosY: 0,
		Code: code,
	}

	i.executor.Execute(h, p, it)
}

func (i *Interactor) findNearInteractables(p *entity.Player, h *level.LevelHolder) []model.TileInteraction {

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

	if t := h.GetInteractionRight(p, view.TileSize); t.Code != 0 {
		result = append(result, *t)
	}

	fmt.Printf("Found near interactables %v", result)

	return result
}

func (i *Interactor) handleInteractCodes(h *level.LevelHolder, p *entity.Player, its []model.TileInteraction) {

	for _, it := range its {
		it := it
		i.executor.Execute(h, p, it)
	}
}
