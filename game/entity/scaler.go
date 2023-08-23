package entity

import "github.com/hajimehoshi/ebiten/v2"

type Scaler interface {
	GetBaseModelSize() float64
	GetPositionY() float64
	GetVisual() *ebiten.Image
}
