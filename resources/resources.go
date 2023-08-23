package resources

import (
	"log"

	"github.com/batariloa/lino/const"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	faceImageFront *ebiten.Image
	faceImageLeft  *ebiten.Image
	faceImageRight *ebiten.Image
)

func LoadResources() error {
	var err error
	faceImageFront, _, err = ebitenutil.NewImageFromFile(artPath.FaceFront)
	if err != nil {
		log.Printf("Error loading resources: %s", err)
	}

	faceImageLeft, _, err = ebitenutil.NewImageFromFile(artPath.FaceLeft)
	if err != nil {
		log.Printf("Error loading resources: %s", err)
	}

	faceImageRight, _, err = ebitenutil.NewImageFromFile(artPath.FaceRight)
	if err != nil {
		log.Printf("Error loading resources: %s", err)
	}

	return err
}

func GetFaceImageFront() *ebiten.Image {
	return faceImageFront
}

func GetFaceImageLeft() *ebiten.Image {
	return faceImageLeft
}

func GetFaceImageRight() *ebiten.Image {
	return faceImageRight
}
