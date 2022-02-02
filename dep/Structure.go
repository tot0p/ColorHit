package dep

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var ImgBase = LoadImg("data/img/sprites.png")

var AllImg []*ebiten.Image = []*ebiten.Image{
	ImgBase.SubImage(image.Rect(128, 16, 144, 32)).(*ebiten.Image), //Mono Wall ....
	ImgBase.SubImage(image.Rect(128, 48, 144, 64)).(*ebiten.Image), // ..|.
	ImgBase.SubImage(image.Rect(128, 80, 144, 96)).(*ebiten.Image), // |...
}

var AllStructure = []*Structure{
	{
		Tile{
			true,
			&RigidBody{
				32, 32, 16, 16,
			},
			AllImg[1],
		},
		Tile{
			true,
			&RigidBody{
				32, 48, 16, 16,
			},
			AllImg[2],
		},
	},
}

type Structure []Tile

func (s *Structure) Draw(screen *ebiten.Image) {
	for _, i := range *s {
		i.Draw(screen)
	}
}

func (s *Structure) Collide(r *RigidBody) bool {
	for _, i := range *s {
		if i.Colide {
			if i.RigidBody.Collide(r) {
				return true
			}
		}
	}
	return false
}
