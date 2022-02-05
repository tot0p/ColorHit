package dep

import (
	"fmt"
	"image"
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 180)
	op.GeoM.Scale(2, 2)
	switch g.Player.NBBall {
	case 0:
		return
	case 1:
		screen.DrawImage(LoadImg("data/img/icons.png").SubImage(image.Rect(80, 96, 96, 112)).(*ebiten.Image), op)
	case 2:
		screen.DrawImage(LoadImg("data/img/icons.png").SubImage(image.Rect(80, 80, 96, 96)).(*ebiten.Image), op)
	case 3:
		screen.DrawImage(LoadImg("data/img/icons.png").SubImage(image.Rect(80, 64, 96, 80)).(*ebiten.Image), op)
	case 4:
		screen.DrawImage(LoadImg("data/img/icons.png").SubImage(image.Rect(80, 48, 96, 64)).(*ebiten.Image), op)
	default:
		screen.DrawImage(LoadImg("data/img/icons.png").SubImage(image.Rect(80, 32, 96, 48)).(*ebiten.Image), op)
	}
}

func (g *GameBody) Update() bool {
	rand.Seed(time.Now().Unix())
	g.Count++
	g.Player.Update(g.Count)
	g.M.Update(g.Player)
	if g.Count%(60*5) == 0 {
		g.M.AddCoin(rand.Intn(500), rand.Intn(300), rand.Intn(3)+1)
		g.M.AddAdd(rand.Intn(500), rand.Intn(300), rand.Intn(3)+1)
	}
	return !(g.Count/60 == 60)
}
