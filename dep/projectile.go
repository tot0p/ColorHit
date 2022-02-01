package dep

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	X, Y int
	//img  *ebiten.Image
	Mouv *Mouv
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	for i := 0; i < 2; i++ {
		for y := 0; y < 2; y++ {
			screen.Set(p.X+y, p.Y+i, color.White)
		}
	}
}

func (p *Projectile) Update() bool {
	if p.Mouv.DestX == p.X && p.Mouv.DestY == p.Y {
		return true
	} else {
		p.Mouv.Apply(&p.X, &p.Y)
		return false
	}
}
