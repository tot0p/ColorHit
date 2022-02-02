package dep

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	X, Y, W, H int
	Img        *ebiten.Image
	Hover      bool
}

func (b *Button) Draw(screen *ebiten.Image) {

}

func (b *Button) Update() {

}
