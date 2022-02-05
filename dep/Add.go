package dep

import (
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Add interface {
	Update(p *Player) bool
	Draw(screen *ebiten.Image)
}

//Ammo
type Ammo struct {
	Nb        int
	RigidBody RigidBody
	Img       *ebiten.Image
	Time      int
}

func (a *Ammo) Update(p *Player) bool {
	a.Time--
	if a.RigidBody.CollideCenter(p.RigidBody) {
		p.NBBall += a.Nb
		return true
	} else if a.Time <= 0 {
		return true
	}
	return false
}

func (a *Ammo) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(a.RigidBody.X), float64(a.RigidBody.Y))
	screen.DrawImage(a.Img, op)
}

func CreateAmmo(x, y, w, h int) Add {
	rand.Seed(time.Now().Unix())
	t := rand.Intn(10)
	var img *ebiten.Image
	if t > 6 {
		t = 10
		temp := []image.Rectangle{image.Rect(64, 48, 80, 64), image.Rect(80, 48, 96, 64), image.Rect(64, 64, 80, 80), image.Rect(80, 64, 96, 80)}[rand.Intn(3)]
		img = LoadImg("data/img/props.png").SubImage(temp).(*ebiten.Image)
	} else {
		t = 3
		temp := []image.Rectangle{image.Rect(64, 16, 80, 32), image.Rect(80, 16, 96, 32), image.Rect(64, 32, 80, 48), image.Rect(80, 32, 96, 48)}[rand.Intn(3)]
		img = LoadImg("data/img/props.png").SubImage(temp).(*ebiten.Image)
	}
	var temp Add = &Ammo{t, RigidBody{x, y, w, h}, img, 600}
	return temp
}