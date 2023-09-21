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
	return h.GetTileAtPos(p.GetPositionX()+float64(tileSize), p.GetPositionY(), tileSize)
}

func (h *LevelHolder) GetTileTop(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(p.GetPositionX(), p.GetPositionY()-float64(tileSize), tileSize)
}

func (h *LevelHolder) GetTileBottom(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(p.GetPositionX(), p.GetPositionY()+float64(tileSize), tileSize)
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

	posX := p.GetPositionX() + float64(tileSize)
	log.Printf("Curr position is %d but right is %d", p.GetPositionX(), posX)
	val := h.GetTileAtPos(posX, p.GetPositionY(), tileSize)
	log.Println("before passing Pos X %d", posX)
	log.Println("Val is %d", val)
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

	// Convert 2D coordinates to 1D index.
	index := yPos*width + xPos
	return h.Interactables[index]
}

func (h *LevelHolder) ReplaceTile(posX, posY, oldTile int, newTile int) {
	log.Printf("Try to replace %d %d %d %d", posX, posY, oldTile, newTile)
	// Calculate the indices of the tile within the level array.
	xPos := posX / view.TileSize
	yPos := posY / view.TileSize

	width := h.MaxLevelWidth / view.TileSize

	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len((*h.Level)[0]) {
		log.Printf("Coordinates or index out of bounds!")
		return
	}

	index := yPos*width + xPos

	// Loop through all layers and replace the tile in each one.
	for i := range *h.Level {
		if index < len((*h.Level)[i]) && (*h.Level)[i][index] == oldTile {
			(*h.Level)[i][index] = newTile
		}
	}
}

func (h *LevelHolder) GenerateLevelOne() {

	LightOn := true

	h.Status = &Status{
		LightOn: &LightOn,
	}

	h.Level = &[][]int{

		{
			63, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
			63, 40, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 40, 40,
			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,

			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,
			63, 62, 42, 42, 42, 42, 42, 44, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,

			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 45, 43, 63, 40,
			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 45, 43, 43, 66, 41, 41, 41, 41, 41, 41, 41, 40, 40,

			63, 62, 43, 43, 43, 45, 43, 43, 63, 40, 40, 40, 40, 40, 40, 40, 40, 40,
		},

		{
			63, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
			63, 40, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 46, 47, 48, 60, 40, 40,
			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 49, 50, 51, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 52, 53, 54, 43, 63, 40,

			63, 62, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,
			63, 62, 42, 42, 42, 42, 42, 44, 42, 42, 42, 42, 42, 42, 42, 42, 63, 40,
			63, 62, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 63, 40,
		},

		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 67, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}

	h.Interactables = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	h.MaxLevelWidth = 18 * 20
	h.MaxLevelHeight = 13 * 20
}
