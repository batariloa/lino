package interactables

import (
	"log"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/model"
)

func Execute(p *entity.Player, ti model.TileInteraction) {
	log.Printf("Trying to execute interactable with code %d", ti.Code)
	if f, ok := InteractionMap[ti.Code]; ok && f != nil {
		f(ti)
	}
}
