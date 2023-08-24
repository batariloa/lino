package game

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  int = 320
	ScreenHeight int = 240
)

type Game struct {
	Player *entity.Player
}

func NewGame(p *entity.Player) *Game {

	return &Game{
		Player: p,
	}
}

func (g *Game) Update() error {

	utils.SetupPlayerControls(g.Player)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Lino, where are you going?")

	g.Player.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
