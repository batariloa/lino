package utils

import (
	entity "github.com/batariloa/lino/game/entity/contract"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

func ScaleImage(model entity.Scaler, screen *ebiten.Image) {
	modelSize := getModelSize(model)
	drawCircle(screen, model.GetPositionX(), model.GetPositionY(), modelSize)

	op := ebiten.DrawImageOptions{}
	scaleImage(model, &op)
	screen.DrawImage(model.GetVisual(), &op)
}

func getModelSize(model entity.Scaler) float64 {
	return 8 + 0.1*model.GetPositionY()
}

func drawCircle(screen *ebiten.Image, x, y, radius float64) {
	ebitenutil.DrawCircle(screen, x, y, radius, color.Opaque)
}

func scaleImage(model entity.Scaler, op *ebiten.DrawImageOptions) {
	modelSize := getModelSize(model)
	faceImage := model.GetVisual()

	scalingFactorX := modelSize / float64(faceImage.Bounds().Dx())
	scalingFactorY := modelSize / float64(faceImage.Bounds().Dy())

	op.GeoM.Scale(scalingFactorX, scalingFactorY)
	op.GeoM.Translate(model.GetPositionX()-modelSize/2, model.GetPositionY()-modelSize/2)

	op.Filter = ebiten.FilterLinear
}
