package dep

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const velocity = 2

type Player struct {
	X, Y, Width, Height, RW, RH int //X et Y au centre du player
	Map                         *Map
	ImgData, ImgData1           *ebiten.Image
	Angle                       int
	Angle2                      int
}

func (p *Player) Update() {
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) && p.X > p.Width/2 {
		p.X -= velocity
		p.Angle = 270
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && p.X < p.RW-p.Width/2 {
		p.X += velocity
		p.Angle = 90
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) && p.Y > p.Height/2 {
		p.Y -= velocity
		p.Angle = 0
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) && p.Y < p.RH-p.Height/2 {
		p.Y += velocity
		p.Angle = 180
	}
	posx, posy := ebiten.CursorPosition()
	vx, vy := float64(p.X-posx), float64(p.Y-posy)
	p.Angle2 = int(math.Atan2(vy, vx)*180/math.Pi) - 90
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.Map.NewProjectile(float64(p.X), float64(p.Y), float64(posx), float64(posy), float64(p.X-posx)/30, float64(p.Y-posy)/30, p.Angle2)
	}
	if p.Angle == 90 || p.Angle == 270 {
		for i := 1; i < p.Width-1; i++ {
			for y := 1; y < 4; y++ {
				p.Map.Set(p.X+i-8, p.Y+y-8, Chen)
			}
			for y := 12; y < 15; y++ {
				p.Map.Set(p.X+i-8, p.Y+y-8, Chen)
			}
		}
	} else {
		for i := 1; i < p.Height-1; i++ {
			for y := 1; y < 4; y++ {
				p.Map.Set(p.X+y-8, p.Y+i-8, Chen)
			}
			for y := 12; y < 15; y++ {
				p.Map.Set(p.X+y-8, p.Y+i-8, Chen)
			}
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	//body
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(16)/2, -float64(16)/2)
	op.GeoM.Rotate(float64(p.Angle%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.ImgData, op)
	//canon
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(16)/2, -float64(16)/2)
	op.GeoM.Rotate(float64(p.Angle2%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.ImgData1, op)
}
