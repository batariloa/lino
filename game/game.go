package game

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *entity.Player
	Tiler  *level.Tiler
}

func StartLevelOne(game *ebiten.Game) {
}

func NewGame(p *entity.Player, t *level.Tiler) *Game {

	t.GenerateLevelOne()
	return &Game{
		Player: p,
	}
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {

		if g.Player.PositionX > g.Player.GetBaseModelSize() {
			g.Player.MoveLeft()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Player.MoveUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Player.MoveDown()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	mapManager := level.NewMapManager(g.Tiler, screen, g.Player)

	mapManager.GenerateLevelOne()
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return level.ScreenWidth, level.ScreenHeight
}
