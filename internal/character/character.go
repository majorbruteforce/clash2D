package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/internal/core"
	"github.com/majorbruteforce/clash2D/pkg/sprite"
)

type Character struct {
	Name string
	*sprite.Cutout
	Pos   struct{ X, Y, Dx, Dy float64 }
	frame int
}

type AnimationSequence struct {
	Start int
	End   int
	Step  int
	Speed int // in ticks per change
}

func NewCharacter(
	name,
	imagePath string,
	framesPerRow,
	numberOfRows int,
) *Character {
	return &Character{
		Name:   name,
		Cutout: sprite.NewCutout(imagePath, framesPerRow, numberOfRows),
		Pos:    struct{ X, Y, Dx, Dy float64 }{X: 0, Y: 0, Dx: 0, Dy: 0},
		frame:  0,
	}
}

func (c *Character) Frame() int {
	return c.frame
}

func (c *Character) SetFrame(idx int) {
	c.frame = idx
}

func (c *Character) RunSequence(seq AnimationSequence) {

	if core.Gb.TickIndex()%seq.Speed != 0 {
		return
	}

	c.frame += seq.Step
	if c.frame > seq.End {
		c.frame = seq.Start
	}
}

func (c *Character) Render(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	// scale
	scaleFactor := float64(core.Gb.UnitSize()) / float64(c.Cutout.TileHeight)
	op.GeoM.Scale(scaleFactor, scaleFactor)

	// transform
	originX, originY := core.Sh.Origin()
	c.Pos.X += c.Pos.Dx
	c.Pos.Y += c.Pos.Dy

	op.GeoM.Translate(c.Pos.X, c.Pos.Y)
	op.GeoM.Translate(float64(originX), float64(originY))
	op.GeoM.Translate(-float64(c.Cutout.TileWidth)*scaleFactor/2, -float64(c.Cutout.TileHeight)*scaleFactor/2)

	c.Pos.Dx, c.Pos.Dy = 0, 0

	// frame
	characterSubImage := c.Cutout.GetSubImageByIndex(c.frame)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(characterSubImage, op)
}
