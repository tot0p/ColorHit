package dep

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	Colide    bool
	RigidBody *RigidBody
	Img       *ebiten.Image
}

func (t *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.RigidBody.X), float64(t.RigidBody.Y))
	screen.DrawImage(t.Img, op)
}
