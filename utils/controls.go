package utils

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

func SetupPlayerControls(p *entity.Player) {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.MoveUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.MoveDown()
	}
}
