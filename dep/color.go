package dep

import (
	"image/color"
)

var (
	Pal = color.Palette{
		color.RGBA{161, 42, 244, 255}, //purlple
		color.RGBA{31, 167, 40, 255},  //green
		color.RGBA{197, 76, 76, 255},  //red
		color.RGBA{228, 228, 36, 255}, //yellow
		color.RGBA{76, 126, 211, 255}, //blue
	}
	Chen = Pal[0]

	none = color.RGBA{0, 0, 0, 0}
)
