package dep

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImg(s string) *ebiten.Image {
	img, err := ebitenutil.NewImageFromURL(s)
	if err != nil {
		panic(err)
	}
	return img
}
func LoadImgImage(s string) image.Image {
	file, _ := os.Open(s)
	defer file.Close()
	img, _, _ := image.Decode(file)
	return img
}
