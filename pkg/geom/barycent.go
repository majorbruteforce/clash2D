package geom

import "image"

func IsPointInTriangle(p, A, B, C image.Point) bool {

	px, py := float64(p.X), float64(p.Y)
	ax, ay := float64(A.X), float64(A.Y)
	bx, by := float64(B.X), float64(B.Y)
	cx, cy := float64(C.X), float64(C.Y)

	denominator := (by-cy)*(ax-cx) + (cx-bx)*(ay-cy)

	alpha := ((by-cy)*(px-cx) + (cx-bx)*(py-cy)) / denominator
	beta := ((cy-ay)*(px-cx) + (ax-cx)*(py-cy)) / denominator
	gamma := 1 - alpha - beta

	return alpha >= 0 && beta >= 0 && gamma >= 0
}
