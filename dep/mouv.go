package dep

type Mouv struct {
	SpeedX, SpeedY float64
	DestX, DestY   float64
	X, Y           float64
	PosX, PosY     bool
}

func (v *Mouv) Apply(x, y *float64) bool {
	var xstop, ystop bool
	if v.X < v.DestX && v.PosX {
		*x -= v.SpeedX
		v.X -= v.SpeedX
		xstop = false
	} else if v.X > v.DestX && !v.PosX {
		*x -= v.SpeedX
		v.X -= v.SpeedX
		xstop = false
	} else {
		xstop = true
	}
	if v.Y < v.DestY && v.PosY {
		*y -= v.SpeedY
		v.Y -= v.SpeedY
		ystop = false
	} else if v.Y > v.DestY && !v.PosY {
		*y -= v.SpeedY
		v.Y -= v.SpeedY
		ystop = false
	} else {
		ystop = true
	}
	return xstop && ystop
}
