package main

import (
	"log"

	"github.com/batariloa/lino/game"
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/interactables"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/view"
	"github.com/batariloa/lino/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	resources.LoadResources()

	ebiten.SetWindowSize(680, 680)
	ebiten.SetWindowTitle("Lino Walks Alone")

	player := entity.NewPlayer(100, 2, 90, 180)
	tiler := view.NewDrawer()
	holder := level.NewLevelHolder()

	iexecutor := interactables.NewInteractablExecutor()
	interactor := interactables.NewInteractor(iexecutor)

	game := game.NewGame(player, tiler, holder, interactor)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
