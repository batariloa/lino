package game

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *entity.Player
}

func NewGame(p *entity.Player) *Game {

	return &Game{
		player: p,
	}
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.MoveUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.MoveDown()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Lino, where are you going?")

	g.player.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
