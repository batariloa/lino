package game

import (
	"log"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *entity.Player
	Drawer *level.Drawer
}

func NewGame(p *entity.Player, t *level.Drawer) *Game {

	return &Game{
		Player: p,
		Drawer: t,
	}
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {

		if g.Player.PositionX > g.Player.GetBaseModelSize() {
			g.Player.MoveLeft()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.Player.PositionX+g.Player.GetBaseModelSize() < float64(g.Drawer.Holder.MaxLevelWidth) {
			g.Player.MoveRight()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.Player.PositionY > g.Player.GetBaseModelSize() {
			g.Player.MoveUp()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if g.Player.PositionY+g.Player.GetBaseModelSize() < float64(g.Drawer.Holder.MaxLevelHeight) {
			g.Player.MoveDown()
		}
	}

	log.Printf("Current position: %f", g.Player.PositionX)
	log.Printf("Current position Y: %f", g.Player.PositionY)
	log.Printf("Current position Y offset: %f", g.Drawer.OffsetY)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Drawer.DrawTiles(screen, g.Player)
	g.Drawer.DrawPlayerModel(screen, g.Player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return level.ScreenWidth, level.ScreenHeight
}
