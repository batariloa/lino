package game

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player     *entity.Player
	mapManager *level.Map
}

func StartLevelOne(game *ebiten.Game) {
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

	tiler := level.NewTiler()
	mapManager := level.NewMapManager(tiler, screen)

	mapManager.GenerateLevelOne()
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return level.ScreenWidth, level.ScreenHeight
}
