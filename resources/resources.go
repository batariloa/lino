package resources

import (
	"io/ioutil"
	"log"

	"github.com/batariloa/lino/const"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	TileArtBytes   []byte
	faceImageFront *ebiten.Image
	faceImageLeft  *ebiten.Image
	faceImageRight *ebiten.Image
)

func init() {
	var err error
	TileArtBytes, err = ioutil.ReadFile(artPath.TileArt)
	if err != nil {
		log.Fatal(err)
	}
}

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
