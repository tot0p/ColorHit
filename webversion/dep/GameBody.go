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

const (
	xc = 31
	yc = 23
)

type GameBody struct {
	Player               *Player
	M                    *Map
	Count                int
	ForbidenX, ForbidenY []int
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
	g.Player.DrawGUI(screen)
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
	if g.Player.NBBall > 0 {
		msg := fmt.Sprintf(": %d", g.Player.NBBall)
		ebitenutil.DebugPrintAt(screen, msg, 42, 360)
	}
}

func (g *GameBody) Update() bool {
	rand.Seed(time.Now().Unix())
	g.Count++
	g.Player.Update(g.Count)
	g.M.Update(g.Player)
	if g.Count%(60*5) == 0 {
		x, y := g.CreateSpawn()
		g.M.AddCoin(x, y, rand.Intn(3)+1)
		x, y = g.CreateSpawn()
		g.M.AddAdd(x, y, 48)
		x, y = g.CreateSpawn()
		g.M.AddAdd(x, y, rand.Intn(6)+1)
	}
	return !(g.Count/60 == 60)
}

func (g *GameBody) CreateSpawn() (int, int) {
	x, y := rand.Intn(xc), rand.Intn(yc)
	for g.M.Collide(&RigidBody{x * 16, y * 16, 16, 16}) {
		x, y = rand.Intn(xc), rand.Intn(yc)
	}
	return x * 16, y * 16
}
