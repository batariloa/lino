package game

import (
	"github.com/batariloa/lino/game/controller"
	"github.com/batariloa/lino/game/entity"
	interact "github.com/batariloa/lino/game/interactables"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *entity.Player
	Drawer *view.Drawer
	Holder *level.LevelHolder
}

func NewGame(p *entity.Player, d *view.Drawer, h *level.LevelHolder) *Game {

	level.NewLevelOne(h)

	return &Game{
		Player: p,
		Drawer: d,
		Holder: h,
	}
}

func (g *Game) Update() error {

	interact.HandlePlayerTriggers(g.Holder, g.Player)

	eKeyIsPressed := ebiten.IsKeyPressed(ebiten.KeyE)

	if eKeyIsPressed && !controller.PrevEKeyState {
		interact.HandlePlayerInteractions(g.Player, g.Holder)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		controller.MoveLeft(g.Player, g.Holder, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		controller.MoveRight(g.Player, g.Holder, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		controller.MoveUp(g.Player, g.Holder, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		controller.MoveDown(g.Player, g.Holder, view.TileSize)
	}

	controller.PrevEKeyState = eKeyIsPressed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	li := model.LevelInfo{
		MaxLevelWidth:  g.Holder.MaxLevelWidth,
		MaxLevelHeight: g.Holder.MaxLevelHeight,
		Level:          *g.Holder.Level,
	}
	g.Drawer.DrawTiles(screen, g.Player, li)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return view.ScreenWidth, view.ScreenHeight
}
