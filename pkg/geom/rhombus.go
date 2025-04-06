package geom

import "image"

type Rhombus struct {
	A, B, C, D image.Point
}

func (r *Rhombus) IsPointInside(p image.Point) bool {
	return IsPointInTriangle(p, r.A, r.B, r.D) || IsPointInTriangle(p, r.B, r.C, r.D)
}

