package view

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/batariloa/lino/game/animator"
	"github.com/batariloa/lino/game/entity"
	"github.com/batariloa/lino/game/model"
	"github.com/batariloa/lino/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  int = 240
	ScreenHeight int = 240
)

const (
	TileSize = 20
)

type Drawer struct {
	RainAnimator *animator.RainAnimator
	screen       *ebiten.Image
	OffsetX      float64
	OffsetY      float64
}

var tilesImage *ebiten.Image

func NewDrawer() *Drawer {
	ra := animator.NewRainAnimator()
	go ra.StartTimer()
	drawer := &Drawer{
		RainAnimator: ra,
	}

	return drawer
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.TileArtBytes))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func (d *Drawer) DrawTiles(screen *ebiten.Image, p *entity.Player, li model.LevelDrawInfo) {
	halfScreenWidth := float64(ScreenWidth / 2)
	halfScreenHeight := float64(ScreenHeight / 2)

	d.OffsetX = 0.0
	if p.PositionX > halfScreenWidth {
		if p.PositionX+halfScreenWidth < float64(li.MaxLevelWidth) {
			d.OffsetX = p.PositionX - halfScreenWidth
		} else {
			d.OffsetX = float64(li.MaxLevelWidth - ScreenWidth)
		}
	}

	d.OffsetY = 0.0
	if p.PositionY > halfScreenHeight {
		if p.PositionY+halfScreenHeight < float64(li.MaxLevelHeight) {
			d.OffsetY = p.PositionY - halfScreenHeight
		} else {
			d.OffsetY = float64(li.MaxLevelHeight - ScreenHeight)
		}
	}

	w := tilesImage.Bounds().Dx()
	tileXCount := w / TileSize
	totalTileCount := li.MaxLevelWidth / TileSize

	// Draw the first two layers.
	for layerIndex := 0; layerIndex < 2 && layerIndex < len(li.Level); layerIndex++ {
		for i, t := range li.Level[layerIndex] {
			drawTile(i, t, tileXCount, totalTileCount, screen, d)
		}
	}

	d.DrawPlayerModel(screen, p)

	// Draw the rest of the layers.
	for layerIndex := 2; layerIndex < len(li.Level); layerIndex++ {
		for i, t := range li.Level[layerIndex] {
			drawTile(i, t, tileXCount, totalTileCount, screen, d)
		}
	}
}

func drawTile(i, t, tileXCount, totalTileCount int, screen *ebiten.Image, d *Drawer) {
	if t > 379 && t < 387 {
		d.drawRainTile(i, tileXCount, totalTileCount, screen)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64((i%totalTileCount)*TileSize)-d.OffsetX,
			float64((i/totalTileCount)*TileSize)-d.OffsetY)

		sx := (t % tileXCount) * TileSize
		sy := (t / tileXCount) * TileSize
		screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+TileSize, sy+TileSize)).(*ebiten.Image), op)
	}
}

func (d *Drawer) drawRainTile(i, tileXCount, totalTileCount int, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((i%totalTileCount)*TileSize)-d.OffsetX,
		float64((i/totalTileCount)*TileSize)-d.OffsetY)

	currentRainFrame := d.RainAnimator.CurrentRainFrame()
	log.Printf("Current rainf rame %d", currentRainFrame)

	sx := (currentRainFrame % tileXCount) * TileSize
	sy := (currentRainFrame / tileXCount) * TileSize
	screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+TileSize, sy+TileSize)).(*ebiten.Image), op)
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
