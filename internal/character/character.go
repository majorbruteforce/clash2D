package character

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/internal/core"
	"github.com/majorbruteforce/clash2D/pkg/sprite"
)

type Character struct {
	Name string
	*sprite.Cutout
	Pos   struct{ X, Y, Dx, Dy float64 }
	frame int
	Seq   string
	Dist  struct{ X, Y float64 }
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
		Dist: struct {
			X float64
			Y float64
		}{X: 0, Y: 0},
		frame: 0,
		Seq:   "WalkS",
	}
}

func (c *Character) Frame() int {
	return c.frame
}

func (c *Character) SetFrame(idx int) {
	c.frame = idx
}

func (c *Character) RunSequence() {

	if c.Dist.X == 0 && c.Dist.Y == 0 {
		c.frame = LucySequences[c.Seq].Start
		return
	}

	if core.Gb.TickIndex()%LucySequences[c.Seq].Speed != 0 {
		return
	}

	c.frame += LucySequences[c.Seq].Step
	if c.frame > LucySequences[c.Seq].End {
		c.frame = LucySequences[c.Seq].Start
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
	characterSubImage := c.Cutout.GetTileByIndex(c.frame)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(characterSubImage, op)
}

func (c *Character) MoveRemainigDist() {

	if core.Gb.TickIndex()%4 == 0 {
		c.Pos.Dx = 1 * Sign(c.Dist.X)
		c.Pos.Dy = 1 * Sign(c.Dist.Y)

		if c.Dist.X != 0 {
			c.Dist.X = Sign(c.Dist.X) * (math.Abs(c.Dist.X) - 1)
		}

		if c.Dist.Y != 0 {
			c.Dist.Y = Sign(c.Dist.Y) * (math.Abs(c.Dist.Y) - 1)
		}
	}
}

func Sign(a float64) float64 {
	if a == 0 {
		return 0
	}

	return a / math.Abs(a)
}
