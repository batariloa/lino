package level

import (
	"log"

	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/game/view"
)

var (
	LevelMap       *[][]int
	Interactables  []int
	InteractionMap map[int]model.TileInteraction
	Triggers       []int
	LevelStatus    *Status
	MaxLevelWidth  int
	MaxLevelHeight int
)

func GetTileAtPos(x float64, y float64, tileSize int) int {
	xPos := int(x) / tileSize
	yPos := int(y) / tileSize

	width := MaxLevelWidth / tileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*LevelMap)[0]) {
		return 0
	}

	// Convert 2D coordinates to 1D index.
	index := yPos*width + xPos
	return (*LevelMap)[0][index]
}
func GetTileLeft(p *entity.Player, tileSize int) int {
	return GetTileAtPos(p.GetPositionX()-float64(tileSize), p.GetPositionY(), tileSize)
}

func GetTileRight(p *entity.Player, tileSize int) int {
	return GetTileAtPos(p.GetPositionX()+float64(p.GetBaseModelSize())+view.TileSize+5, p.GetPositionY(), tileSize)
}

func GetTileTop(p *entity.Player, tileSize int) int {
	return GetTileAtPos(p.GetPositionX(), p.GetPositionY()-float64(tileSize), tileSize)
}

func GetTileBottom(p *entity.Player, tileSize int) int {
	return GetTileAtPos(p.GetPositionX(), p.GetPositionY()+float64(tileSize), tileSize)
}

func GetTriggerAtPos(x, y float64) int {

	xPos := int(x) / view.TileSize
	yPos := int(y) / view.TileSize

	width := MaxLevelWidth / view.TileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*LevelMap)[0]) {
		return 0
	}

	index := yPos*width + xPos
	return Triggers[index]
}

func GetInteractionLeft(p *entity.Player, tileSize int) *model.TileInteraction {
	posX := p.GetPositionX() - float64(tileSize)
	val := GetInteractionAtPos(posX, p.GetPositionY(), tileSize)
	return &model.TileInteraction{
		PosX: posX,
		PosY: p.PositionY,
		Code: val,
	}
}

func GetInteractionRight(p *entity.Player, tileSize int) *model.TileInteraction {

	posX := p.GetPositionX() + p.GetBaseModelSize()
	val := GetInteractionAtPos(posX, p.GetPositionY(), view.TileSize)
	return &model.TileInteraction{
		PosX: posX,
		PosY: p.PositionY,
		Code: val,
	}
}

func GetInteractionTop(p *entity.Player, tileSize int) *model.TileInteraction {
	val := GetInteractionAtPos(p.GetPositionX(), p.GetPositionY()-float64(tileSize), tileSize)
	return &model.TileInteraction{
		PosX: p.PositionX,
		PosY: p.PositionY,
		Code: val,
	}
}

func GetInteractionBottom(p *entity.Player, tileSize int) *model.TileInteraction {

	val := GetInteractionAtPos(p.GetPositionX(), p.GetPositionY()+float64(tileSize), tileSize)
	return &model.TileInteraction{
		PosX: p.PositionX,
		PosY: p.PositionY,
		Code: val,
	}
}

func GetInteractionAtPos(x float64, y float64, tileSize int) int {
	xPos := int(x) / tileSize
	yPos := int(y) / tileSize

	width := MaxLevelWidth / tileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*LevelMap)[0]) {
		return 0
	}

	index := yPos*width + xPos
	return Interactables[index]
}

func ReplaceTile(posX, posY, oldTile int, newTile int) {
	xPos := posX / view.TileSize
	yPos := posY / view.TileSize

	width := MaxLevelWidth / view.TileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*LevelMap)[0]) {
		log.Printf("Coordinates or index out of bounds!")
		return
	}

	index := yPos*width + xPos

	for i := range *LevelMap {
		if index < len((*LevelMap)[i]) &&
			(*LevelMap)[i][index] == oldTile {
			(*LevelMap)[i][index] = newTile
		}
	}
}
