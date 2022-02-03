package dep

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coin struct {
	Tier      int
	Img       *ebiten.Image
	RigidBody RigidBody
}

func CreateCoin(tier int, x, y int) *Coin {
	var img *ebiten.Image = LoadImg("data/img/props.png")
	switch tier {
	case 1:
		img = img.SubImage(image.Rect(144, 16, 160, 32)).(*ebiten.Image)
	case 2:
		img = img.SubImage(image.Rect(160, 16, 176, 32)).(*ebiten.Image)
	default:
		img = img.SubImage(image.Rect(176, 16, 192, 32)).(*ebiten.Image)
	}
	return &Coin{tier, img, RigidBody{0, 0, 16, 16}}
}

func (c *Coin) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.RigidBody.X), float64(c.RigidBody.Y))
	screen.DrawImage(c.Img, op)

}

func (c *Coin) GetPoint() int {
	switch c.Tier {
	case 1:
		return 5000
	case 2:
		return 10000
	default:
		return 25000
	}
}
