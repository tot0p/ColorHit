package dep

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sco struct {
	Img                     *ebiten.Image
	ButtonStart             RigidBody
	ButtonQuit              RigidBody
	BestScore, CurrentScore int
}

func (s *Sco) SetScore(Score int) {
	if Score > s.BestScore {
		s.BestScore = Score
	}
	s.CurrentScore = Score
}

func (s *Sco) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	msg := fmt.Sprintf(`Point : %d`, s.CurrentScore/10*100)
	screen.DrawImage(s.Img, op)
	ebitenutil.DebugPrintAt(screen, msg, 224, 172)
}

func (s *Sco) Update(temp, count *int) bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tx, ty := ebiten.CursorPosition()
		temp := &RigidBody{tx, ty, 5, 5}
		if s.ButtonStart.Collide(temp) {
			return true
		} else if s.ButtonQuit.Collide(temp) {
			os.Exit(0)
			return false
		}
	}
	return false
}
