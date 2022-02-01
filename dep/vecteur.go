package dep

type Mouv struct {
	SpeedX, SpeedY int
	DestX, DestY   int
}

func (v *Mouv) Apply(x, y *int) {
	if *x < v.DestX-v.SpeedX/2 {
		*x += v.SpeedX
	} else {
		*x -= v.SpeedX
	}
	if *y < v.DestY-v.SpeedX/2 {
		*y += v.SpeedY
	} else {
		*y -= v.SpeedY
	}
}
