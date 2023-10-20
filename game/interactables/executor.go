package interactables

import (
	"log"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
)

type InteractableExecutor struct{}

func NewInteractablExecutor() *InteractableExecutor {
	return &InteractableExecutor{}
}

func (*InteractableExecutor) Execute(h *level.LevelHolder, p *entity.Player, ti model.TileInteraction) {
	log.Printf("Trying to execute interactable with code %d", ti.Code)
	if f, ok := InteractionMap[ti.Code]; ok && f != nil {
		f(ti, h)
	}
}
