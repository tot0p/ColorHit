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

func LoadScore() int {
	t, err := LoadJson("data/save.json")
	if err != nil {
		t = map[string]int{"BestScore": 0}
		WriteJson("data/save.json", t)
	}
	if _, exist := t["BestScore"]; !exist {
		t = map[string]int{"BestScore": 0}
		WriteJson("data/save.json", t)
	}
	return t["BestScore"]
}

func (s *Sco) SetScore(Score int) {
	if Score > s.BestScore {
		s.BestScore = Score / 10 * 100
		t := map[string]int{"BestScore": Score}
		WriteJson("data/save.json", t)
	}
	s.CurrentScore = Score / 10 * 100
}

func (s *Sco) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	msg := fmt.Sprintf(`Point : %d`, s.CurrentScore)
	msg2 := fmt.Sprintf(`Best Point : %d`, s.BestScore)
	screen.DrawImage(s.Img, op)
	ebitenutil.DebugPrintAt(screen, msg, 224, 172)
	ebitenutil.DebugPrintAt(screen, msg2, 224, 188)
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
