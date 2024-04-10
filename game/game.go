package game

import (
	"github.com/batariloa/lino/game/controller"
	"github.com/batariloa/lino/game/entity"
	interact "github.com/batariloa/lino/game/interact"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *entity.Player
	Drawer *view.Drawer
}

func NewGame(p *entity.Player, d *view.Drawer) *Game {

	level.NewLevelOne()

	return &Game{
		Player: p,
		Drawer: d,
	}
}

func (g *Game) Update() error {

	interact.HandlePlayerTriggers(g.Player)

	eKeyIsPressed := ebiten.IsKeyPressed(ebiten.KeyE)

	if eKeyIsPressed && !controller.PrevEKeyState {
		interact.HandlePlayerInteractions(g.Player)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		controller.MoveLeft(g.Player, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		controller.MoveRight(g.Player, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		controller.MoveUp(g.Player, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		controller.MoveDown(g.Player, view.TileSize)
	}

	controller.PrevEKeyState = eKeyIsPressed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	li := model.LevelInfo{
		MaxLevelWidth:  level.MaxLevelWidth,
		MaxLevelHeight: level.MaxLevelHeight,
		Level:          *level.LevelMap,
	}
	g.Drawer.DrawTiles(screen, g.Player, li)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return view.ScreenWidth, view.ScreenHeight
}
