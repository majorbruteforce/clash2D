package core

import (
	"clash2D/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	Name       string
	Sheet      *ebiten.Image
	Cutout     *utils.Cutout
	Position   struct{ X, Y, Dx, Dy int }
	Frame      int
	Visible    bool
	Properties struct{}
	// figure out how to absorb camera offsets
}

func NewEntity(
	name string,
	sheet *ebiten.Image,
	cutout *utils.Cutout,
) *Entity {
	return &Entity{
		Name:   name,
		Sheet:  sheet,
		Cutout: cutout,
		Position: struct {
			X  int
			Y  int
			Dx int
			Dy int
		}{X: 0, Y: 0, Dx: 0, Dy: 0},
		Frame:      0,
		Visible:    true,
		Properties: struct{}{},
	}
}

// start and end setup sequence range
// step is the step size for selecting frame ids
// if a vertical sequence is desired, use an non-unit step size
// speed is in ticks per change
func (e *Entity) UpdateFrame(start, end, step, speed, tickIndex int) {
	if tickIndex%speed != 0 {
		return
	}
	e.Frame += step
	if e.Frame > end {
		e.Frame = start
	}
}


