package dep

import "image/color"

var (
	Pal = color.Palette([]color.Color{
		color.RGBA{255, 117, 117, 255},
		color.RGBA{255, 255, 117, 255},
		color.RGBA{160, 255, 117, 255},
		color.RGBA{80, 117, 255, 255},
		color.RGBA{160, 117, 255, 255},
	})
	none = color.RGBA{0, 0, 0, 0}
)
