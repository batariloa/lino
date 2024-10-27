package game

import (
	"fmt"

	"github.com/batariloa/lino/game/controller"
	"github.com/batariloa/lino/game/entity"
	interact "github.com/batariloa/lino/game/interact"
	"github.com/batariloa/lino/game/level"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/trigger"
	"github.com/batariloa/lino/game/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player *entity.Player
	Drawer *view.Drawer
}

func NewGame(p *entity.Player, d *view.Drawer) *Game {
	err := level.LoadLevel("bedroom")
	level.NewLevelOne()

	if err != nil {
		fmt.Sprint("no level data", err.Error())
		panic(err)
	}

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
	li := model.LevelDrawInfo{
		MaxLevelWidth:  level.MaxLevelWidth,
		MaxLevelHeight: level.MaxLevelHeight,
		Level:          *level.LevelMap,
	}

	g.Drawer.DrawTiles(screen, g.Player, li)

	width := level.MaxLevelWidth / view.TileSize

	fmt.Print("Max width %d", width)
	triggerIndex := g.Player.PositionY*float64(width) + g.Player.PositionX

	debugText := fmt.Sprintf("Pixel X %f Y %f", g.Player.PositionX, g.Player.PositionY)
	debugText2 := fmt.Sprintf("Tile X %f Y %f -- %f", g.Player.PositionX/view.TileSize,
		g.Player.PositionY/view.TileSize, triggerIndex)

	ebitenutil.DebugPrintAt(screen, debugText, 0, screen.Bounds().Max.Y-32)
	ebitenutil.DebugPrint(screen, debugText2)

	trigger.HandleTransitionToBlack(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return view.ScreenWidth, view.ScreenHeight
}
