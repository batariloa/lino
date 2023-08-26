package game

import (
	"log"

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
		Tiler:  t,
	}
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {

		if g.Player.PositionX > g.Player.GetBaseModelSize() {
			g.Player.MoveLeft()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.Player.PositionX+g.Player.GetBaseModelSize() < float64(g.Tiler.MaxLevelWidth) {
			g.Player.MoveRight()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.Player.PositionY > g.Player.GetBaseModelSize() {
			g.Player.MoveUp()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if g.Player.PositionY+g.Player.GetBaseModelSize() < float64(g.Tiler.MaxLevelHeight) {
			g.Player.MoveDown()
		}
	}

	log.Printf("Current position: %f", g.Player.PositionX)
	log.Printf("Current position Y: %f", g.Player.PositionY)
	log.Printf("Current position Y offset: %f", g.Tiler.OffsetX)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Tiler.DrawTiles(screen, g.Player)
	g.Player.DrawPlayerModel(screen, g.Tiler.OffsetX, g.Tiler.OffsetY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return level.ScreenWidth, level.ScreenHeight
}
