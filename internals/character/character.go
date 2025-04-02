package character

import (
	"clash2D/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Character struct {
	Name     string
	Sheet    *ebiten.Image
	Cutout   *utils.Cutout
	Position struct{ X, Y, Dx, Dy float64 }
	frame    int
}

type AnimationSequence struct {
	Start int
	End   int
	Step  int
	Speed int // in ticks per change
}

func NewCharacter(
	name string,
	sheet *ebiten.Image,
	cutout *utils.Cutout,
	frame int,
) *Character {
	return &Character{
		Name:     name,
		Sheet:    sheet,
		Cutout:   cutout,
		Position: struct{ X, Y, Dx, Dy float64 }{X: 0, Y: 0, Dx: 0, Dy: 0},
		frame:    frame,
	}
}

func (c *Character) Frame() int {
	return c.frame
}

func (c *Character) SetFrame(id int) {
	c.frame = id
}

func (c *Character) RunSequence(tickIndex int, seq AnimationSequence) {
	if tickIndex%seq.Speed != 0 {
		return
	}

	c.frame += seq.Step
	if c.frame > seq.End {
		c.frame = seq.Start
	}
}
