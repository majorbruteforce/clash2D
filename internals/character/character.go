package character

import (
	"clash2D/internals/core"
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
	originX, originY := core.Gb.Origin()
	c.Position.X += c.Position.Dx
	c.Position.Y += c.Position.Dy

	op.GeoM.Translate(c.Position.X, c.Position.Y)
	op.GeoM.Translate(float64(originX), float64(originY))
	op.GeoM.Translate(-float64(c.Cutout.TileWidth)*scaleFactor/2, -float64(c.Cutout.TileHeight)*scaleFactor/2)

	c.Position.Dx, c.Position.Dy = 0, 0

	// frame
	characterSubImage := c.Sheet.SubImage(c.Cutout.GetTileRectangleById(c.frame)).(*ebiten.Image)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(characterSubImage, op)
}
