package dep

import (
	"math"
)

type Mouv struct {
	SpeedX, SpeedY int
	DestX, DestY   int
}

func (v *Mouv) Apply(x, y *int) bool {
	var xstop, ystop bool
	if math.Abs(float64(*x)) > math.Abs(float64(v.DestX-v.SpeedX/2)) {
		*x -= v.SpeedX
		xstop = false
	} else {
		xstop = true
	}
	if math.Abs(float64(*y)) > math.Abs(float64(v.DestX-v.SpeedX/2)) {
		*y -= v.SpeedY
		ystop = false
	} else {
		ystop = true
	}
	return xstop && ystop
}
