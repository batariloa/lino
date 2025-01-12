package constants

import "path/filepath"

var (
	PathTileArt   string
	PathFaceFront string
	PathFaceLeft  string
	PathFaceRight string
)

func init() {

	PathTileArt = filepath.Join(".", "art", "tiles.png")
	PathFaceFront = filepath.Join(".", "art", "face32-front.png")
	PathFaceLeft = filepath.Join(".", "art", "face32-left.png")
	PathFaceRight = filepath.Join(".", "art", "face32-right.png")
}
