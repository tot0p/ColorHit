package dep

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Map struct {
	Img       *ebiten.Image
	Color     []*Point
	Proj      *Projectile
	Structure []*Structure
}

func (m *Map) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(m.Img, op)
	for _, i := range m.Structure {
		i.Draw(screen)
	}
	if m.Proj != nil {
		m.Proj.Draw(screen)
	}
}

func (m *Map) Update() {
	if m.Proj != nil {
		t := m.Proj.Update()
		if t {
			m.Proj = nil
		}
	}
}

func (m *Map) Set(x, y int, color color.Color) {
	m.Img.Set(x, y, color)
	temp := &Point{x, y}
	if !VerifPoint(m.Color, temp) {
		m.Color = append(m.Color, temp)
	}
}

func VerifPoint(a []*Point, k *Point) bool {
	for _, i := range a {
		if *i == *k {
			return true
		}
	}
	return false
}

func (m *Map) NewProjectile(x, y, destX, destY, speedx, speedy float64, angle int) {
	if m.Proj == nil {
		xt, yt := false, false
		if x < destX {
			xt = true
		}
		if y < destY {
			yt = true
		}
		m.Proj = &Projectile{x, y, LoadImg("data/img/tank.png").SubImage(image.Rect(32, 0, 48, 16)).(*ebiten.Image), &Mouv{speedx, speedy, destX, destY, x, y, xt, yt}, angle}
	}
}

type Point struct {
	X, Y int
}
