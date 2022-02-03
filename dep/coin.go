package dep

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Coin struct {
	Tier      int
	Img       *ebiten.Image
	RigidBody RigidBody
}

func CreateCoin(tier int, x, y int) *Coin {
	return &Coin{}
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
