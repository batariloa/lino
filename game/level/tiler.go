package level

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	ScreenWidth  int = 240
	ScreenHeight int = 240
)

const (
	tileSize = 16
)

type Tiler struct {
}

var (
	tilesImage *ebiten.Image
)

func NewTiler() *Tiler {

	return &Tiler{}
}
func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func (t *Tiler) DrawTiles(screen *ebiten.Image, layers [][]int) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	const xCount = ScreenWidth / tileSize
	for _, l := range layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}
}
