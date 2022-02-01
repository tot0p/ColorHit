package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tot0p/JamOneMinute/dep"
)

const (
	screenWidth  = 512
	screenHeight = 384
	resolWidth   = 512
	resolHeight  = 384
)

var (
	game *Game
	m    dep.Map
)

type Game struct {
	Player *dep.Player
	count  int
}

func (g *Game) Update() error {
	g.count++
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	m.Draw(screen)
	g.Player.Draw(screen)
	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f
`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return resolWidth, resolHeight
}

func init() {
	m = dep.Map{ebiten.NewImage(resolWidth, resolHeight)}
	game = &Game{
		&dep.Player{0, 0, 10, 10, resolWidth, resolHeight, &m},
		0,
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Nok")
	ebiten.SetMaxTPS(60)
	ebiten.SetWindowResizable(true)
	ebiten.SetFullscreen(false)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
