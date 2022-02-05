package dep

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var ImgBase = LoadImg("data/img/sprites.png")

var AllImg []*ebiten.Image = []*ebiten.Image{
	ImgBase.SubImage(image.Rect(128, 16, 144, 32)).(*ebiten.Image), //Mono Wall .... 0
	ImgBase.SubImage(image.Rect(128, 48, 144, 64)).(*ebiten.Image), // ..|. 1
	ImgBase.SubImage(image.Rect(128, 80, 144, 96)).(*ebiten.Image), // |... 2
	ImgBase.SubImage(image.Rect(160, 16, 176, 32)).(*ebiten.Image), // .|.. 3
	ImgBase.SubImage(image.Rect(176, 16, 192, 32)).(*ebiten.Image), // .|.| 4
	ImgBase.SubImage(image.Rect(192, 16, 208, 32)).(*ebiten.Image), // ...| 5
	ImgBase.SubImage(image.Rect(160, 48, 176, 64)).(*ebiten.Image), // .||. 6
	ImgBase.SubImage(image.Rect(176, 48, 192, 64)).(*ebiten.Image), // .||| 7
	ImgBase.SubImage(image.Rect(192, 48, 208, 64)).(*ebiten.Image), // ..|| 8
	ImgBase.SubImage(image.Rect(160, 64, 176, 80)).(*ebiten.Image), // |.|| 9
	ImgBase.SubImage(image.Rect(176, 64, 192, 80)).(*ebiten.Image), // .... 10
	ImgBase.SubImage(image.Rect(192, 64, 208, 80)).(*ebiten.Image), // |.|| 11
	ImgBase.SubImage(image.Rect(160, 80, 176, 96)).(*ebiten.Image), // ||.. 12
	ImgBase.SubImage(image.Rect(176, 80, 192, 96)).(*ebiten.Image), // ||.| 13
	ImgBase.SubImage(image.Rect(192, 80, 208, 96)).(*ebiten.Image), // |..| 14
	ImgBase.SubImage(image.Rect(224, 16, 240, 32)).(*ebiten.Image), // sp 15
	ImgBase.SubImage(image.Rect(224, 48, 240, 64)).(*ebiten.Image), // sp 16
}

var AllMap = [][]*Structure{
	{
		{
			CreateTile(64, 64, true, AllImg[6]),
			CreateTile(80, 64, true, AllImg[7]),
			{
				true,
				&RigidBody{
					96, 64, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					112, 64, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					128, 64, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					144, 64, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					160, 64, 16, 16,
				},
				AllImg[8],
			},
			{
				true,
				&RigidBody{
					64, 80, 16, 16,
				},
				AllImg[12],
			},
			{
				true,
				&RigidBody{
					80, 80, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					96, 80, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					112, 80, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					128, 80, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					144, 80, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					160, 80, 16, 16,
				},
				AllImg[14],
			},
		},
		{
			{
				true,
				&RigidBody{
					224, 48, 16, 16,
				},
				AllImg[6],
			},
			{
				true,
				&RigidBody{
					240, 48, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					256, 48, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					272, 48, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					288, 48, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					304, 48, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					320, 48, 16, 16,
				},
				AllImg[8],
			},
			{
				true,
				&RigidBody{
					224, 64, 16, 16,
				},
				AllImg[12],
			},
			{
				true,
				&RigidBody{
					240, 64, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					256, 64, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					272, 64, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					288, 64, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					304, 64, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					320, 64, 16, 16,
				},
				AllImg[14],
			},
		},
		{
			{
				true,
				&RigidBody{
					48, 144, 16, 16,
				},
				AllImg[6],
			},
			{
				true,
				&RigidBody{
					64, 144, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					80, 144, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					96, 144, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					112, 144, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					128, 144, 16, 16,
				},
				AllImg[7],
			},
			{
				true,
				&RigidBody{
					144, 144, 16, 16,
				},
				AllImg[8],
			},
			{
				true,
				&RigidBody{
					48, 160, 16, 16,
				},
				AllImg[12],
			},
			{
				true,
				&RigidBody{
					64, 160, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					80, 160, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					96, 160, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					112, 160, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					128, 160, 16, 16,
				},
				AllImg[13],
			},
			{
				true,
				&RigidBody{
					144, 160, 16, 16,
				},
				AllImg[14],
			},
		},
	},
}

func CreateTile(x, y int, isSolid bool, img *ebiten.Image) Tile {
	return Tile{isSolid, &RigidBody{x, y, 16, 16}, img}
}

type Structure []Tile

func (s *Structure) Draw(screen *ebiten.Image) {
	for _, i := range *s {
		if i.Colide {
			i.Draw(screen)
		}
	}
}

func (s *Structure) DrawBefore(screen *ebiten.Image) {
	for _, i := range *s {
		if !i.Colide {
			i.Draw(screen)
		}
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
