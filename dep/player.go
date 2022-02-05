package dep

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const velocity = 2

type Player struct {
	RigidBody         *RigidBody
	RW, RH            int //X et Y au centre du player
	Map               *Map
	ImgData, ImgData1 *ebiten.Image
	Angle             int
	Angle2            int
	NBBall            int
	VB                int
	Time              int
}

func (p *Player) Update(count int) {
	if p.VB != 0 {
		p.Time--
		if p.Time <= 0 {
			p.VB = 0
		}
	}
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) && p.RigidBody.X > 0 {
		p.RigidBody.X -= velocity + p.VB
		p.Angle = 270
		if p.Map.Collide(p.RigidBody) {
			p.RigidBody.X += velocity + p.VB
		}
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && p.RigidBody.X < p.RW-p.RigidBody.Width {
		p.RigidBody.X += velocity + p.VB
		p.Angle = 90
		if p.Map.Collide(p.RigidBody) {
			p.RigidBody.X -= velocity + p.VB
		}
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) && p.RigidBody.Y > 0 {
		p.RigidBody.Y -= velocity + p.VB
		p.Angle = 0
		if p.Map.Collide(p.RigidBody) {
			p.RigidBody.Y += velocity + p.VB
		}
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) && p.RigidBody.Y < p.RH-p.RigidBody.Height {
		p.RigidBody.Y += velocity + p.VB
		p.Angle = 180
		if p.Map.Collide(p.RigidBody) {
			p.RigidBody.Y -= velocity + p.VB
		}
	}
	posx, posy := ebiten.CursorPosition()
	vx, vy := float64(p.RigidBody.X-posx), float64(p.RigidBody.Y-posy)
	p.Angle2 = int(math.Atan2(vy, vx)*180/math.Pi) - 90
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.Map.Proj == nil && count >= 15 && p.NBBall > 0 {
		p.NBBall--
		t := math.Sqrt(math.Pow(float64(p.RigidBody.X+8-posx), 2)+math.Pow(float64(p.RigidBody.Y+8-posy), 2)) * 0.15
		p.Map.NewProjectile(float64(p.RigidBody.X+8), float64(p.RigidBody.Y+8), float64(posx), float64(posy), float64(p.RigidBody.X+8-posx)/t, float64(p.RigidBody.Y+8-posy)/t, p.Angle2)
	}
	if p.Angle == 90 || p.Angle == 270 {
		for i := 0; i < p.RigidBody.Width; i++ {
			for y := 0; y < 3; y++ {
				p.Map.Set(p.RigidBody.X+i, p.RigidBody.Y+y, Chen)
			}
			for y := 13; y < 16; y++ {
				p.Map.Set(p.RigidBody.X+i, p.RigidBody.Y+y, Chen)
			}
		}
	} else {
		for i := 0; i < p.RigidBody.Height; i++ {
			for y := 0; y < 3; y++ {
				p.Map.Set(p.RigidBody.X+y, p.RigidBody.Y+i, Chen)
			}
			for y := 13; y < 16; y++ {
				p.Map.Set(p.RigidBody.X+y, p.RigidBody.Y+i, Chen)
			}
		}
	}
	p.Map.Point += p.Map.CoinCheck(p.RigidBody)
}

func (p *Player) Draw(screen *ebiten.Image) {
	//body
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(16)/2, -float64(16)/2)
	op.GeoM.Rotate(float64(p.Angle%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(p.RigidBody.X+p.RigidBody.Width/2), float64(p.RigidBody.Y+p.RigidBody.Height/2))
	screen.DrawImage(p.ImgData, op)
	//canon
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(16)/2, -float64(16)/2)
	op.GeoM.Rotate(float64(p.Angle2%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(p.RigidBody.X+p.RigidBody.Width/2), float64(p.RigidBody.Y+p.RigidBody.Height/2))
	screen.DrawImage(p.ImgData1, op)
}

func (p *Player) DrawGUI(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(2.5, 75)
	op.GeoM.Scale(2, 2)
	var img *ebiten.Image
	switch p.VB {
	case -1:
		img = LoadImg("data/img/icons.png").SubImage(image.Rect(144, 32, 160, 48)).(*ebiten.Image)
	case 1:
		img = LoadImg("data/img/icons.png").SubImage(image.Rect(112, 32, 128, 48)).(*ebiten.Image)
	case 2:
		img = LoadImg("data/img/icons.png").SubImage(image.Rect(128, 32, 144, 48)).(*ebiten.Image)
	}
	if p.VB != 0 {
		screen.DrawImage(img, op)
		op.GeoM.Translate(0, 9)
	}
}
