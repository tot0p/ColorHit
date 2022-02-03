package dep

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Acc struct {
	Img         *ebiten.Image
	ButtonStart RigidBody
	ButtonQuit  RigidBody
}

func (a *Acc) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(a.Img, op)
}

func (a *Acc) Update(temp, count *int) bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tx, ty := ebiten.CursorPosition()
		temp := &RigidBody{tx, ty, 5, 5}
		if a.ButtonStart.Collide(temp) {
			return true
		} else if a.ButtonQuit.Collide(temp) {
			os.Exit(0)
			return false
		}
	}
	return false
}
