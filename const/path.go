package artPath

import "path/filepath"

var (
	TileArt   string
	FaceFront string
	FaceLeft  string
	FaceRight string
)

func init() {

	TileArt = filepath.Join(".", "art", "tiles.png")
	FaceFront = filepath.Join(".", "art", "face32-front.png")
	FaceLeft = filepath.Join(".", "art", "face32-left.png")
	FaceRight = filepath.Join(".", "art", "face32-right.png")
}
