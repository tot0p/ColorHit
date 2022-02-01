package dep

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const velocity = 2

type Player struct {
	X, Y, Width, Height, RW, RH int
	Map                         *Map
	//imgData *ebiten.Image
}

func (p *Player) Update() {
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) && p.X > 0 {
		p.X -= velocity
	}
	if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) && p.X < p.RW-p.Width {
		p.X += velocity
	}
	if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) && p.Y > 0 {
		p.Y -= velocity
	}
	if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) && p.Y < p.RH-p.Height {
		p.Y += velocity
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		for i := 0; i < p.Width; i++ {
			for y := 0; y < p.Height; y++ {
				p.Map.Set(p.X+y, p.Y+i, Pal[rand.Intn(5)])
			}
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	for i := 0; i < p.Width; i++ {
		for y := 0; y < p.Height; y++ {
			screen.Set(p.X+y, p.Y+i, color.White)
		}
	}

}
