package level

import "github.com/batariloa/lino/game/entity"

type LevelHolder struct {
	Level          [][]int
	Interactables  []int
	MaxLevelWidth  int
	MaxLevelHeight int
}

func NewLevelHolder() *LevelHolder {

	return &LevelHolder{}
}

func (h *LevelHolder) GetTileAtPos(x, y, tileSize int) int {
	// Convert pixel values to map coordinates.
	xPos := x / tileSize
	yPos := y / tileSize

	// Calculate width in terms of tiles.
	width := h.MaxLevelWidth / tileSize

	// Ensure the position is within the bounds of the map.
	if xPos < 0 || yPos < 0 || xPos >= width || yPos*width+xPos >= len(h.Level[0]) {
		return 0 // some default value indicating no tile or error.
	}

	// Convert 2D coordinates to 1D index.
	index := yPos*width + xPos
	return h.Level[0][index]
}

func (h *LevelHolder) GetTileLeft(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(int(p.GetPositionX())-int(p.GetBaseModelSize()), int(p.GetPositionY()), tileSize)
}

func (h *LevelHolder) GetTileRight(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(int(p.GetPositionX())-int(p.GetBaseModelSize()), int(p.GetPositionY()), tileSize)
}

func (h *LevelHolder) GetTileTop(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(int(p.GetPositionX())-int(p.GetBaseModelSize()), int(p.GetPositionY()), tileSize)
}

func (h *LevelHolder) GetTileBottom(p *entity.Player, tileSize int) int {
	return h.GetTileAtPos(int(p.GetPositionX())-int(p.GetBaseModelSize()), int(p.GetPositionY()), tileSize)
}

func (h *LevelHolder) GenerateLevelOne() {
	h.Level = [][]int{

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
		0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	h.MaxLevelWidth = 18 * 20
	h.MaxLevelHeight = 13 * 20
}
