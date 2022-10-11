package main

import (
	"image"
	"log"
	"math/rand"
	"time"

	"ColorHit/dep"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 512
	screenHeight = 384
	resolWidth   = 512
	resolHeight  = 384
)

var (
	cursorGame = dep.LoadImg("data/img/icons.png").SubImage(image.Rect(112, 80, 127, 95)).(*ebiten.Image)
	cursorMenu = dep.LoadImg("data/img/icons.png").SubImage(image.Rect(128, 80, 144, 96)).(*ebiten.Image)
	game       *Game
	temp       int
	start      bool = true
)

type Game struct {
	gamebody *dep.GameBody
	count    int
	start    bool
	Acc      *dep.Acc
	Sco      *dep.Sco
}

func (g *Game) Update() error {
	if g.start {
		g.start = g.gamebody.Update()
		if ebiten.IsKeyPressed(ebiten.KeyEscape) {
			g.start = false
			start = true
		}
		if !g.start {
			dep.Chen = dep.Pal[rand.Intn(len(dep.Pal))]
			g.Sco.SetScore(g.gamebody.M.Point)
			t := &dep.Map{ebiten.NewImage(resolWidth, resolHeight), 0, nil, dep.AllMap[rand.Intn(len(dep.AllMap))], []*dep.Coin{}, []dep.Add{}, 1, 0}
			img := dep.LoadImg("data/img/tank.png")
			g.gamebody = &dep.GameBody{
				&dep.Player{&dep.RigidBody{resolWidth/2 - 4, resolHeight/2 - 4, 16, 16}, resolWidth, resolHeight, t, img.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), img.SubImage(image.Rect(16, 0, 32, 16)).(*ebiten.Image), 0, 0, 3, 0, 0},
				t,
				0,
				[]int{},
				[]int{},
			}
			g.gamebody.Player.RigidBody.X, g.gamebody.Player.RigidBody.Y = game.gamebody.CreateSpawn()
		}
	} else if start {
		g.start = g.Acc.Update(&temp, &g.count)
		start = !g.start
	} else {
		g.start = g.Sco.Update(&temp, &g.count)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.start {
		g.gamebody.Draw(screen)
	} else if start {
		g.Acc.Draw(screen)
	} else {
		g.Sco.Draw(screen)
	}
	t, x := ebiten.CursorPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t), float64(x))
	if g.start {
		op.GeoM.Translate(-7, -7)
		screen.DrawImage(cursorGame, op)
	} else {
		screen.DrawImage(cursorMenu, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return resolWidth, resolHeight
}

func init() {
	rand.Seed(time.Now().Unix())
	t := &dep.Map{ebiten.NewImage(resolWidth, resolHeight), 0, nil, dep.AllMap[rand.Intn(len(dep.AllMap))], []*dep.Coin{}, []dep.Add{}, 1, 0}
	img := dep.LoadImg("data/img/tank.png")
	game = &Game{
		&dep.GameBody{
			&dep.Player{&dep.RigidBody{resolWidth/2 - 4, resolHeight/2 - 4, 16, 16}, resolWidth, resolHeight, t, img.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), img.SubImage(image.Rect(16, 0, 32, 16)).(*ebiten.Image), 0, 0, 3, 0, 0},
			t,
			0,
			[]int{},
			[]int{},
		},
		0,
		false,
		&dep.Acc{
			dep.LoadImg("data/ui/menu.png"),
			dep.RigidBody{224, 192, 64, 32},
			dep.RigidBody{224, 240, 64, 32},
		},
		&dep.Sco{
			dep.LoadImg("data/ui/score_menu.png"),
			dep.RigidBody{192, 240, 128, 32},
			dep.RigidBody{224, 288, 64, 32},
			dep.LoadScore(),
			0,
		},
	}
	game.gamebody.Player.RigidBody.X, game.gamebody.Player.RigidBody.Y = game.gamebody.CreateSpawn()
	dep.Chen = dep.Pal[rand.Intn(len(dep.Pal))]
}

func main() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Color Hit")
	ebiten.SetMaxTPS(60)
	ebiten.SetWindowResizable(true)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowIcon([]image.Image{dep.LoadImgImage("data/logo16.png"), dep.LoadImgImage("data/logo32.png"), dep.LoadImgImage("data/logo48.png")})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
