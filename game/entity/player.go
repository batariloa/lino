package entity

import (
	"image/color"

	"github.com/batariloa/lino/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const BaseModelSize = 15

type Player struct {
	Health        int
	Speed         float64
	positionXPrev float64
	PositionX     float64
	PositionY     float64
}

func NewPlayer(health int, speed float64, posX float64, posY float64) *Player {

	return &Player{
		Health:        health,
		Speed:         speed,
		PositionX:     posX,
		positionXPrev: posX,
		PositionY:     posY,
	}
}

func (p *Player) DrawPlayerModel(screen *ebiten.Image, offsetX float64, offsetY float64) {

	op := ebiten.DrawImageOptions{}
	faceImage := p.GetVisual()

	ebitenutil.DrawCircle(
		screen,
		p.PositionX-offsetX,
		p.PositionY-offsetY,
		BaseModelSize,
		color.Opaque)

	scalingFactorX := BaseModelSize / float64(faceImage.Bounds().Dx())
	scalingFactorY := BaseModelSize / float64(faceImage.Bounds().Dy())

	op.GeoM.Scale(scalingFactorX, scalingFactorY)
	op.GeoM.Translate(p.PositionX-BaseModelSize/2, p.PositionY-BaseModelSize/2-float64(offsetY))

	op.Filter = ebiten.FilterLinear
	screen.DrawImage(faceImage, &op)

	p.positionXPrev = p.PositionX
}

func (p *Player) GetPositionX() float64 {
	return p.PositionX
}

func (p *Player) GetPositionY() float64 {
	return p.PositionY
}

func (p *Player) GetVisual() *ebiten.Image {

	if p.PositionX > p.positionXPrev {
		return resources.GetFaceImageRight()
	}
	if p.PositionX < p.positionXPrev {
		return resources.GetFaceImageLeft()
	}

	return resources.GetFaceImageFront()
}

func (p *Player) GetBaseModelSize() float64 {
	return BaseModelSize
}

func (p *Player) MoveLeft() {
	p.PositionX -= p.Speed
}

func (p *Player) MoveRight() {
	p.PositionX += p.Speed
}

func (p *Player) MoveUp() {
	p.PositionY -= p.Speed
}

func (p *Player) MoveDown() {
	p.PositionY += p.Speed
}
