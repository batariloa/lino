package artPath

import "path/filepath"

var (
	FaceFront string
	FaceLeft  string
	FaceRight string
)

func init() {

	FaceFront = filepath.Join(".", "art", "face32-front.png")
	FaceLeft = filepath.Join(".", "art", "face32-left.png")
	FaceRight = filepath.Join(".", "art", "face32-right.png")
}
