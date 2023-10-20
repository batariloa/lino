package level

import (
	"log"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
)

type LevelHolder struct {
	Level          *[][]int
	Interactables  []int
	Triggers       []int
	Status         *Status
	MaxLevelWidth  int
	MaxLevelHeight int
}

func NewLevelHolder() *LevelHolder {

	return &LevelHolder{}
}

func (h *LevelHolder) GetTileAtPos(x float64, y float64, tileSize int) int {
	xPos := int(x) / tileSize
	yPos := int(y) / tileSize

	width := h.MaxLevelWidth / tileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*h.Level)[0]) {
		return 0
	}

	// Convert 2D coordinates to 1D index.
	index := yPos*width + xPos
	return (*h.Level)[0][index]
}
func (h *LevelHolder) GetTileLeft(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(p.GetPositionX()-float64(tileSize), p.GetPositionY(), tileSize)
}

func (h *LevelHolder) GetTileRight(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(p.GetPositionX()+float64(p.GetBaseModelSize())+view.TileSize+5, p.GetPositionY(), tileSize)
}

func (h *LevelHolder) GetTileTop(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(p.GetPositionX(), p.GetPositionY()-float64(tileSize), tileSize)
}

func (h *LevelHolder) GetTileBottom(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(p.GetPositionX(), p.GetPositionY()+float64(tileSize), tileSize)
}

func (h *LevelHolder) GetTriggerAtPos(x, y float64) int {

	xPos := int(x) / view.TileSize
	yPos := int(y) / view.TileSize

	width := h.MaxLevelWidth / view.TileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*h.Level)[0]) {
		return 0
	}

	index := yPos*width + xPos
	return h.Triggers[index]
}

func (h *LevelHolder) GetInteractionLeft(p *entity.Player, tileSize int) *model.TileInteraction {
	posX := p.GetPositionX() - float64(tileSize)
	val := h.GetInteractionAtPos(posX, p.GetPositionY(), tileSize)
	return &model.TileInteraction{
		PosX: posX,
		PosY: p.PositionY,
		Code: val,
	}
}

func (h *LevelHolder) GetInteractionRight(p *entity.Player, tileSize int) *model.TileInteraction {

	posX := p.GetPositionX() + p.GetBaseModelSize()
	val := h.GetInteractionAtPos(posX, p.GetPositionY(), view.TileSize)
	return &model.TileInteraction{
		PosX: posX,
		PosY: p.PositionY,
		Code: val,
	}
}

func (h *LevelHolder) GetInteractionTop(p *entity.Player, tileSize int) *model.TileInteraction {
	val := h.GetInteractionAtPos(p.GetPositionX(), p.GetPositionY()-float64(tileSize), tileSize)
	return &model.TileInteraction{
		PosX: p.PositionX,
		PosY: p.PositionY,
		Code: val,
	}
}

func (h *LevelHolder) GetInteractionBottom(p *entity.Player, tileSize int) *model.TileInteraction {

	val := h.GetInteractionAtPos(p.GetPositionX(), p.GetPositionY()+float64(tileSize), tileSize)
	return &model.TileInteraction{
		PosX: p.PositionX,
		PosY: p.PositionY,
		Code: val,
	}
}

func (h *LevelHolder) GetInteractionAtPos(x float64, y float64, tileSize int) int {
	xPos := int(x) / tileSize
	yPos := int(y) / tileSize

	width := h.MaxLevelWidth / tileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*h.Level)[0]) {
		return 0
	}

	index := yPos*width + xPos
	return h.Interactables[index]
}

func (h *LevelHolder) ReplaceTile(posX, posY, oldTile int, newTile int) {
	xPos := posX / view.TileSize
	yPos := posY / view.TileSize

	width := h.MaxLevelWidth / view.TileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*h.Level)[0]) {
		log.Printf("Coordinates or index out of bounds!")
		return
	}

	index := yPos*width + xPos

	for i := range *h.Level {
		if index < len((*h.Level)[i]) &&
			(*h.Level)[i][index] == oldTile {
			(*h.Level)[i][index] = newTile
		}
	}
}
