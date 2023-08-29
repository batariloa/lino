package level

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/batariloa/lino/game/animator"
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  int = 240
	ScreenHeight int = 240
)

const (
	tileSize = 20
)

type Drawer struct {
	Holder       *LevelHolder
	RainAnimator *animator.RainAnimator
	screen       *ebiten.Image
	OffsetX      float64
	OffsetY      float64
}

var (
	tilesImage *ebiten.Image
)

func NewDrawer() *Drawer {
	ra := animator.NewRainAnimator()
	go ra.StartTimer()
	drawer := &Drawer{
		Holder:       NewLevelHolder(),
		RainAnimator: ra,
	}

	drawer.Holder.GenerateLevelOne()

	return drawer
}

func init() {

	img, _, err := image.Decode(bytes.NewReader(resources.TileArtBytes))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func (d *Drawer) DrawTiles(screen *ebiten.Image, p *entity.Player) {
	halfScreenWidth := float64(ScreenWidth / 2)
	halfScreenHeight := float64(ScreenHeight / 2)

	d.OffsetX = 0.0
	if p.PositionX > halfScreenWidth {
		if p.PositionX+halfScreenWidth < float64(d.Holder.MaxLevelWidth) {
			d.OffsetX = p.PositionX - halfScreenWidth
		} else {
			d.OffsetX = float64(d.Holder.MaxLevelWidth - ScreenWidth)
		}
	}

	d.OffsetY = 0.0
	if p.PositionY > halfScreenHeight {
		if p.PositionY+halfScreenHeight < float64(d.Holder.MaxLevelHeight) {
			d.OffsetY = p.PositionY - halfScreenHeight
		} else {
			d.OffsetY = float64(d.Holder.MaxLevelHeight - ScreenHeight)
		}
	}

	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	var xCount = d.Holder.MaxLevelWidth / tileSize
	for _, l := range d.Holder.Level {
		for i, t := range l {

			if t > 379 && t < 400 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64((i%xCount)*tileSize)-d.OffsetX,
					float64((i/xCount)*tileSize)-d.OffsetY)

				currentRainFrame := d.RainAnimator.CurrentRainFrame()
				log.Printf("Current rainf rame %d", currentRainFrame)

				sx := (currentRainFrame % tileXCount) * tileSize
				sy := (currentRainFrame / tileXCount) * tileSize
				screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)

			} else {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64((i%xCount)*tileSize)-d.OffsetX,
					float64((i/xCount)*tileSize)-d.OffsetY)

				sx := (t % tileXCount) * tileSize
				sy := (t / tileXCount) * tileSize
				screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)

			}
		}
	}
}

func (d *Drawer) DrawPlayerModel(screen *ebiten.Image, p *entity.Player) {

	op := ebiten.DrawImageOptions{}
	faceImage := p.GetVisual()

	vector.DrawFilledCircle(
		screen,
		float32(p.PositionX-d.OffsetX),
		float32(p.PositionY-d.OffsetY),
		float32(p.GetBaseModelSize()),
		color.Opaque,
		false)

	scalingFactorX := p.GetBaseModelSize() / float64(faceImage.Bounds().Dx())
	scalingFactorY := p.GetBaseModelSize() / float64(faceImage.Bounds().Dy())

	op.GeoM.Scale(scalingFactorX, scalingFactorY)
	op.GeoM.Translate(
		p.PositionX-p.GetBaseModelSize()/2-d.OffsetX,
		p.PositionY-p.GetBaseModelSize()/2-d.OffsetY)

	op.Filter = ebiten.FilterLinear
	screen.DrawImage(faceImage, &op)

	p.PositionXPrev = p.PositionX
}
