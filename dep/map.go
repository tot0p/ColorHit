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

func (m *Map) AddCoin(x, y, t int) {
	m.Coin = append(m.Coin, CreateCoin(t, x, y))
}

func (m *Map) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(m.Img, op)
	for _, i := range m.Structure {
		i.Draw(screen)
	}
	for _, i := range m.Coin {
		if i != nil {
			i.Draw(screen)
		}
	}
	if m.Proj != nil {
		m.Proj.Draw(screen)
	}
}

func (m *Map) Update() {
	if m.Proj != nil {
		t := m.Proj.Update()
		if t {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(m.Proj.X-8, m.Proj.Y-8)
			m.Point += 188
			m.Img.DrawImage(LoadImg("data/img/props.png").SubImage(image.Rect(144+16*m.Proj.Tier, 48, 160+16*m.Proj.Tier, 64)).(*ebiten.Image), op)
			m.Proj = nil
		}
	}
	for k, i := range m.Coin {
		if i != nil {
			i.Time -= 1
			if i.Time <= 0 {
				m.Coin[k] = nil
			}
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

func (m *Map) CoinCheck(r *RigidBody) int {
	for k, i := range m.Coin {
		if i != nil {
			if r.CollideCenter(&i.RigidBody) {
				t := i.GetPoint()
				m.Coin[k] = nil
				return t
			}
		}
	}
	return 0
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
		rand.Seed(time.Now().Unix())
		t := rand.Intn(5)
		m.Proj = &Projectile{x, y, LoadImg("data/img/tank.png").SubImage(image.Rect(32+16*t, 0, 48+16*t, 16)).(*ebiten.Image), &Mouv{speedx, speedy, destX, destY, x, y, xt, yt}, angle, RigidBody{int(x - 2), int(y - 2), 4, 4}, t}
	}
}
