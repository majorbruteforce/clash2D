package core

type Shared struct {
	originX int
	originY int
}

func (s *Shared) Origin() (int, int) {
	return s.originX, s.originY
}

func (s *Shared) UpdateOrigin(x, y int) {
	s.originX += x
	s.originY += y
}

var Sh = Shared{
	originX: 0,
	originY: 0,
}