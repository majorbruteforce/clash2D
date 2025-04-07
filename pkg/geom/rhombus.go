package geom

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Rhombus struct {
	A, B, C, D Point
}

func NewRhombusFromTile(tileX, tileY, tileWidth, tileHeight, offsetX, offsetY int) *Rhombus {

	tX, tY := float32(tileX), float32(tileY)
	tW, tH := float32(tileWidth), float32(tileHeight)
	offX, offY := float32(offsetX), float32(offsetY)

	a := Point{
		X: tX + offX,
		Y: tY + tH/2 + offY,
	}
	b := Point{
		X: tX + tW/2 + offX,
		Y: tY + tH/4 + offY,
	}
	c := Point{
		X: tX + tW + offX,
		Y: tY + tH/2 + offY,
	}
	d := Point{
		X: tX + tW/2 + offX,
		Y: tY + 3*tH/4 + offY,
	}

	return &Rhombus{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (r *Rhombus) GetPoints() []Point {
	return []Point{
		r.A, r.B, r.C, r.D,
	}
}

func (r *Rhombus) IsPointInside(p Point) bool {
	return IsPointInTriangle(p, r.A, r.B, r.D) || IsPointInTriangle(p, r.B, r.C, r.D)
}

func (r *Rhombus) Stroke(dst *ebiten.Image, strokeWidth float32, clr color.Color, antialias bool) {

	ax, ay := r.A.X, r.A.Y
	bx, by := r.B.X, r.B.Y
	cx, cy := r.C.X, r.C.Y
	dx, dy := r.D.X, r.D.Y

	vector.StrokeLine(dst, ax, ay, bx, by, strokeWidth, clr, antialias)
	vector.StrokeLine(dst, bx, by, cx, cy, strokeWidth, clr, antialias)
	vector.StrokeLine(dst, cx, cy, dx, dy, strokeWidth, clr, antialias)
	vector.StrokeLine(dst, dx, dy, ax, ay, strokeWidth, clr, antialias)
}

func (r *Rhombus) Overlaps(rb *Rhombus) bool {
	rH, rbH := r.D.Y-r.B.Y, rb.D.Y-rb.B.Y

	minRX, minRY, maxRX, maxRY := r.A.X, r.A.Y-rH/2, r.C.X, r.C.Y+rH/2
	minRbX, minRbY, maxRbX, maxRbY := rb.A.X, rb.A.Y-rbH/2, rb.C.X, rb.C.Y+rbH/2

	if minRX > maxRbX || minRbX > maxRX || minRY > maxRbY || minRbY > maxRY {
		return false
	}

	for _, p := range r.GetPoints() {
		if rb.IsPointInside(p) {
			return true
		}
	}

	for _, p := range rb.GetPoints() {
		if r.IsPointInside(p) {
			return true
		}
	}

	return false
}

func (r *Rhombus) BoundsOverlap(rb *Rhombus) bool {
	rH, rbH := r.D.Y-r.B.Y, rb.D.Y-rb.B.Y

	minRX, minRY, maxRX, maxRY := r.A.X, r.A.Y-rH/2, r.C.X, r.C.Y+rH/2
	minRbX, minRbY, maxRbX, maxRbY := rb.A.X, rb.A.Y-rbH/2, rb.C.X, rb.C.Y+rbH/2

	if minRX < maxRbX && minRbX < maxRX && minRY < maxRbY && minRbY < maxRY {
		return true
	}
	
	return false
}
