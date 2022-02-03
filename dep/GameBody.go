package dep

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameBody struct {
	Player *Player
	M      *Map
	Count  int
}

func (g *GameBody) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{31, 31, 31, 255})
	g.M.Draw(screen)
	g.Player.Draw(screen)
	point := fmt.Sprintf(`Points : %d`, g.M.Point/10*100)
	ebitenutil.DebugPrint(screen, point)
	time := fmt.Sprintf(`time : %d `, 60-g.Count/60)
	ebitenutil.DebugPrintAt(screen, time, 200, 5)

}

func (g *GameBody) Update() bool {
	rand.Seed(time.Now().Unix())
	g.Count++
	g.Player.Update(g.Count)
	g.M.Update()
	if g.Count%(60*5) == 0 {
		g.M.AddCoin(rand.Intn(500), rand.Intn(300), rand.Intn(3)+1)
	}
	return !(g.Count/60 == 60)
}
