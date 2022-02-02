package dep

import "github.com/hajimehoshi/ebiten/v2"

type Acc struct {
	ImgLogo        *ebiten.Image
	ImgButtonStart Button
	ImgButtonQuit  Button
}

func (a *Acc) Draw(screen *ebiten.Image) {
	a.ImgButtonStart.Draw(screen)
	a.ImgButtonQuit.Draw(screen)
}

func (a *Acc) Update(temp, count *int) {

}
