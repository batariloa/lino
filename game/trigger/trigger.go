package trigger

import (
	"image/color"

	"github.com/batariloa/lino/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Triggers          map[int]func()
	transitionToBlack bool
	transitionAlpha   uint8
)

func init() {
	Triggers = map[int]func(){}
}

func calcTriggerIndex(x, y, levelWidth int) int {

	width := levelWidth / view.TileSize
	index := y*width + x

	return index
}

func AddTrigger(x, y, width int, triggerFunc func()) {

	index := calcTriggerIndex(x, y, width)
	Triggers[index] = triggerFunc
}

func HandleTransitionToBlack(screen *ebiten.Image) {
	if transitionToBlack && transitionAlpha < 255 {
		transitionAlpha += 5
		gray := uint8(transitionAlpha)
		color := color.RGBA{gray, gray, gray, 255}
		screen.Fill(color)
	}
}

func StartTransitionToBlack() {
	transitionAlpha = 0
	transitionToBlack = true
}

func GetTriggerAtPos(xPixel, yPixel float64, levelWidth int) func() {
	tileX := int(xPixel / view.TileSize)
	tileY := int(yPixel / view.TileSize)
	index := calcTriggerIndex(tileX, tileY, int(levelWidth))
	return Triggers[index]
}
