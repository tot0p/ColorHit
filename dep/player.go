package dep

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const velocity = 2

type Player struct {
	X, Y, Width, Height, RW, RH int
	Map                         *Map
	ImgData, ImgData1           *ebiten.Image
	Angle                       int
	Angle2                      int
}

func (p *Player) Update() {
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) && p.X > 0 {
		p.X -= velocity
		p.Angle = 270
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && p.X < p.RW-p.Width {
		p.X += velocity
		p.Angle = 90
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) && p.Y > 0 {
		p.Y -= velocity
		p.Angle = 0
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) && p.Y < p.RH-p.Height {
		p.Y += velocity
		p.Angle = 180
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		p.Map.NewProjectile(p.X+7, p.Y, p.X+7, p.Y-160, 4)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		for i := 0; i < p.Width; i++ {
			for y := 0; y < p.Height; y++ {
				p.Map.Set(p.X+y, p.Y+i, Pal[rand.Intn(5)])
			}
		}
	}
	posx, posy := ebiten.CursorPosition()
	fmt.Println(ebiten.CursorPosition())
	fmt.Println(float64(posx-p.X), " ", float64(p.Y-posy))
	vx, vy := float64(p.X-posx), float64(p.Y-posy)

	sca := -8*vx + -8*vy
	ra := math.Sqrt(math.Pow(vx, 2) + math.Pow(vy, 2))
	ra1 := math.Sqrt(math.Pow(-8, 2) + math.Pow(-8, 2))
	fmt.Println(int(math.Acos(sca / (ra * ra1))))
	p.Angle2 = int(math.Acos(sca / (ra * ra1)))

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
	//fmt.Println(p.Angle2)
	op.GeoM.Rotate(float64(p.Angle2%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.ImgData1, op)
}
