package geom

type Point struct {
	X, Y float32
}

func IsPointInTriangle(p, A, B, C Point) bool {

	px, py := p.X, p.Y
	ax, ay := A.X, A.Y
	bx, by := B.X, B.Y
	cx, cy := C.X, C.Y

	denominator := (by-cy)*(ax-cx) + (cx-bx)*(ay-cy)

	alpha := ((by-cy)*(px-cx) + (cx-bx)*(py-cy)) / denominator
	beta := ((cy-ay)*(px-cx) + (ax-cx)*(py-cy)) / denominator
	gamma := 1 - alpha - beta

	return alpha >= 0 && beta >= 0 && gamma >= 0
}
