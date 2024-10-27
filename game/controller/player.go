package controller

import (
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/level"
)

var PrevEKeyState bool

func MoveLeft(p *entity.Player, tileSize int) {
	leftTile := level.GetTileAtPos(p.GetPositionX()-float64(p.GetBaseModelSize()), p.GetPositionY(), tileSize)
	isNoPassTile := level.IsNoPassTile(leftTile)

	if !isNoPassTile &&
		p.PositionX > p.GetBaseModelSize() {
		p.MoveLeft()
	}
}

func MoveRight(p *entity.Player, tileSize int) {
	rightTile := level.GetTileAtPos(p.GetPositionX()+float64(p.GetBaseModelSize()), p.GetPositionY(), tileSize)
	isNoPassTile := level.IsNoPassTile(rightTile)

	if !isNoPassTile &&
		float64(level.MaxLevelWidth-int(p.GetBaseModelSize())) > p.PositionX {
		p.MoveRight()
	}
}

func MoveUp(p *entity.Player, tileSize int) {
	topTile := level.GetTileAtPos(p.GetPositionX(), p.GetPositionY()-float64(p.GetBaseModelSize()), tileSize)
	isNoPassTile := level.IsNoPassTile(topTile)

	if !isNoPassTile {
		p.MoveUp()
	}
}

func MoveDown(p *entity.Player, tileSize int) {
	topTile := level.GetTileAtPos(p.GetPositionX(), p.GetPositionY()+float64(p.GetBaseModelSize()), tileSize)
	isNoPassTile := level.IsNoPassTile(topTile)
	if !isNoPassTile &&
		p.PositionY+p.GetBaseModelSize() < float64(level.MaxLevelHeight) {
		p.MoveDown()
	}
}
