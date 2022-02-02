package dep

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var ImgBase = LoadImg("data/img/sprites.png")

var AllImg []*ebiten.Image = []*ebiten.Image{
	ImgBase.SubImage(image.Rect(128, 16, 144, 32)).(*ebiten.Image), //Mono Wall ....
	ImgBase.SubImage(image.Rect(128, 48, 144, 64)).(*ebiten.Image), // ..|.
}

var AllStructure []*Structure = []*Structure{
	&Structure{[]*Tile{
		&Tile{
			true,
			&RigidBody{
				0, 0, 16, 16,
			},
			AllImg[0],
		},
	}},
}

type Structure struct {
	Tile []*Tile
}

func (s *Structure) Draw(screen *ebiten.Image) {
	for _, i := range s.Tile {
		i.Draw(screen)
	}
}
