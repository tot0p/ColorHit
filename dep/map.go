package dep

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Map struct {
	Img       *ebiten.Image
	Point     int
	Proj      *Projectile
	Structure []*Structure
	Coin      []*Coin
}

func (m *Map) AddCoin() {
	rand.Seed(time.Now().Unix())
	m.Coin = append(m.Coin, CreateCoin(rand.Intn(3)+1, 0, 0))
}

func (m *Map) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(m.Img, op)
	for _, i := range m.Structure {
		i.Draw(screen)
	}
	for _, i := range m.Coin {
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

func (m *Map) Collide(r *RigidBody) bool {
	for _, i := range m.Structure {
		if i.Collide(r) {
			return true
		}
	}
	return false
}

func (m *Map) Set(x, y int, color color.Color) {
	if m.Img.At(x, y) != color {
		m.Img.Set(x, y, color)
		m.Point++
	}
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
