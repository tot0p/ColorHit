package dep

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImg(s string) *ebiten.Image {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
func LoadImgImage(s string) image.Image {
	file, _ := os.Open(s)
	defer file.Close()
	img, _, _ := image.Decode(file)
	return img
}
