package dep

type RigidBody struct {
	X, Y, Width, Height int
}

func (h *RigidBody) CollideCenter(h2 *RigidBody) bool {
	X, Y := h2.X+h2.Width/2, h2.Y+h2.Height/2
	return h.X <= X && X <= h.X+h.Width && h.Y <= Y && Y <= h.Y+h.Height
}

func (h *RigidBody) Collide(h2 *RigidBody) bool {
	var t [4][2]int
	t[0][0], t[0][1] = h2.X, h2.Y
	t[1][0], t[1][1] = h2.X+h2.Width-3, h2.Y
	t[2][0], t[2][1] = h2.X+h2.Width-3, h2.Y+h2.Height-2
	t[3][0], t[3][1] = h2.X, h2.Y+h2.Height-2
	for _, i := range t {
		if h.X <= i[0] && i[0] <= h.X+h.Width-3 && h.Y <= i[1] && i[1] <= h.Y+h.Height-2 {
			return true
		}
	}
	return false
}

func (h *RigidBody) Update(X, Y float64) {
	h.X, h.Y = int(X), int(Y)
}
