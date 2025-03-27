package main

import "github.com/hajimehoshi/ebiten/v2"

type Camera struct {
	X float64
	Y float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

func (c *Camera) Move(key ebiten.Key, offset float64) {
	movement := map[ebiten.Key]struct{ dx, dy float64 }{
		ebiten.KeyArrowUp:    {0, +offset},
		ebiten.KeyArrowDown:  {0, -offset},
		ebiten.KeyArrowLeft:  {offset, 0},
		ebiten.KeyArrowRight: {-offset, 0},
	}

	if move, ok := movement[key]; ok {
		c.X += move.dx
		c.Y += move.dy
	}
}

func (c *Camera) MouseMove(dx, dy, rate float64) {
	c.X += dx * rate
	c.Y += dy * rate
}
