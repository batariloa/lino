package game

import (
	"github.com/batariloa/lino/game/controller"
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/interactor"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player    *entity.Player
	Drawer    *view.Drawer
	Holder    *level.LevelHolder
	KeyStates *controller.KeyStates
	pc        *controller.PlayerController
}

func NewGame(p *entity.Player, t *view.Drawer, h *level.LevelHolder) *Game {

	return &Game{
		Player:    p,
		Drawer:    t,
		Holder:    h,
		KeyStates: controller.NewKeyStates(),
		pc:        controller.NewPlayerController(),
	}
}

func (g *Game) Update() error {

	eKeyIsPressed := ebiten.IsKeyPressed(ebiten.KeyE)

	if eKeyIsPressed && !g.KeyStates.PrevEKeyState {
		interactor.HandlePlayerInteractions(g.Player, g.Holder)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.pc.MoveLeft(g.Player, g.Holder, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.pc.MoveRight(g.Player, g.Holder, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.pc.MoveUp(g.Player, g.Holder, view.TileSize)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.pc.MoveDown(g.Player, g.Holder, view.TileSize)
	}

	g.KeyStates.PrevEKeyState = eKeyIsPressed

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
