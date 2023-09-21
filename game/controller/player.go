package controller

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
)

type PlayerController struct{}

func NewPlayerController() *PlayerController {
	return &PlayerController{}
}

func (*PlayerController) MoveLeft(p *entity.Player, h *level.LevelHolder, tileSize int) {

	leftTile := h.GetTileAtPos(p.GetPositionX()-float64(p.GetBaseModelSize()), p.GetPositionY(), tileSize)
	isNoPassTile := level.IsNoPassTile(leftTile)

	if !isNoPassTile &&
		p.PositionX > p.GetBaseModelSize() {
		p.MoveLeft()
	}
}

func (*PlayerController) MoveRight(p *entity.Player, h *level.LevelHolder, tileSize int) {

	rightTile := h.GetTileAtPos(p.GetPositionX()+float64(p.GetBaseModelSize()), p.GetPositionY(), tileSize)
	isNoPassTile := level.IsNoPassTile(rightTile)

	if !isNoPassTile &&
		float64(h.MaxLevelWidth-int(p.GetBaseModelSize())) > p.PositionX {
		p.MoveRight()
	}
}

func (*PlayerController) MoveUp(p *entity.Player, h *level.LevelHolder, tileSize int) {

	topTile := h.GetTileAtPos(p.GetPositionX(), p.GetPositionY()-float64(p.GetBaseModelSize()), tileSize)
	isNoPassTile := level.IsNoPassTile(topTile)

	if !isNoPassTile {
		p.MoveUp()
	}
}

func (*PlayerController) MoveDown(p *entity.Player, h *level.LevelHolder, tileSize int) {

	topTile := h.GetTileAtPos(p.GetPositionX(), p.GetPositionY()+float64(p.GetBaseModelSize()), tileSize)
	isNoPassTile := level.IsNoPassTile(topTile)
	if !isNoPassTile &&
		p.PositionY+p.GetBaseModelSize() < float64(h.MaxLevelHeight) {
		p.MoveDown()
	}
}
