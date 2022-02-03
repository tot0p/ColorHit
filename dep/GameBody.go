package dep

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameBody struct {
	Player *Player
	M      *Map
	Count  int
}

func (g *GameBody) Draw(screen *ebiten.Image) {
	g.M.Draw(screen)
	g.Player.Draw(screen)
	point := fmt.Sprintf(`Points : %d `, g.M.Point/10*100)
	ebitenutil.DebugPrint(screen, point)
	time := fmt.Sprintf(`time : %d `, 60-g.Count/60)
	ebitenutil.DebugPrintAt(screen, time, 200, 5)

}

func (g *GameBody) Update() bool {
	g.Count++
	g.Player.Update(g.Count)
	g.M.Update()
	return !(g.Count/60 == 60)
}
