package entity

import (
	"github.com/batariloa/lino/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const baseModelSize = 9

type Player struct {
	Health        int
	Speed         float64
	PositionX     float64
	positionXPrev float64
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

func (p *Player) Draw(screen *ebiten.Image) {
	modelImage := utils.ScaleModelByY(p)
	screen.DrawImage(modelImage, &ebiten.DrawImageOptions{})
}

func (p *Player) GetBaseModelSize() float64 {

	return baseModelSize
}

func (p *Player) MoveLeft() {

	p.PositionX -= p.Speed
}

func (p *Player) MoveRight() {

	p.PositionX += p.Speed
}

func (p *Player) MoveUp() {

	p.PositionX -= p.Speed
}

func (p *Player) MoveDown() {

	p.PositionX += p.Speed
}
