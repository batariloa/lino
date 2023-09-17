package level

var noPassTiles map[int]bool

func init() {

	noPassTiles = make(map[int]bool)

	noPassTileIds := []int{
		40, 41,
		60, 61, 62, 63, 64, 65, 66,
	}

	for _, tile := range noPassTileIds {
		noPassTiles[tile] = true
	}
}

func IsNoPassTile(tileId int) bool {

	if noPassTiles[tileId] {
		return true
	}

	return false
}
