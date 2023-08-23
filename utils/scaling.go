package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func ScaleModelByY(model con.Scaler) *ebiten.Image {

	modelSize := model.GetBaseModelSize() + 0.1*model.GetPositionY()
	visual := model.GetVisual()

	scalingFactorX := modelSize / float64(visual.Bounds().Dx())
	scalingFactorY := modelSize / float64(visual.Bounds().Dy())

	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(scalingFactorX, scalingFactorY)
	op.GeoM.Translate(scalingFactorX, model.GetPositionY()-modelSize/2)

	op.Filter = ebiten.FilterLinear
	scaledImage := ebiten.NewImage(int(modelSize), int(modelSize))
	scaledImage.DrawImage(visual, &op)

	return scaledImage
}
