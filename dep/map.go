package dep

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Map struct {
	Img *ebiten.Image
}

func (m *Map) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(m.Img, op)
}

func (m *Map) Set(x, y int, color color.Color) {
	m.Img.Set(x, y, color)
}
