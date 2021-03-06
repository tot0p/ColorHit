package dep

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	X, Y      float64 // au centre du Projectile
	Img       *ebiten.Image
	Mouv      *Mouv
	Angle     int
	RigidBody RigidBody
	Tier      int
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	//body
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(16)/2, -float64(16)/2)
	op.GeoM.Rotate(float64(p.Angle%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.Img, op)
	screen.DrawImage(p.Img, op)
}

func (p *Projectile) Update() bool {
	t := p.Mouv.Apply(&p.X, &p.Y)
	p.RigidBody.Update(p.X-2, p.Y-2)
	return t
}
