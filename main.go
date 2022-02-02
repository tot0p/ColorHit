package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"

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
	game  *Game
	debug bool
	temp  int
)

type Game struct {
	gamebody *dep.GameBody
	count    int
	start    bool
}

func (g *Game) Update() error {
	if g.start {
		g.start = g.gamebody.Update()
	}
	if ebiten.IsKeyPressed(ebiten.KeyF3) && temp+15 < g.count {
		debug = !debug
		temp = g.count
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.start {
		g.gamebody.Draw(screen)
	}
	//debug
	if debug {
		msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, msg)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return resolWidth, resolHeight
}

func init() {
	rand.Seed(time.Now().Unix())
	debug = false
	t := &dep.Map{ebiten.NewImage(resolWidth, resolHeight), 0, nil, dep.AllStructure}
	img := dep.LoadImg("data/img/tank.png")
	game = &Game{
		&dep.GameBody{
			&dep.Player{&dep.RigidBody{resolWidth/2 - 4, resolHeight/2 - 4, 16, 16}, resolWidth, resolHeight, t, img.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), img.SubImage(image.Rect(16, 0, 32, 16)).(*ebiten.Image), 0, 0},
			t,
			0,
		},
		0,
		true,
	}
	dep.Chen = dep.Pal[rand.Intn(len(dep.Pal))]
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Color Hit")
	ebiten.SetMaxTPS(60)
	ebiten.SetWindowResizable(true)
	ebiten.SetFullscreen(false)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
