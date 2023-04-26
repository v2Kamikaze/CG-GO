package core

func Intersection(scan int, seg [][]int) (float64, float64) {
	xi := float64(seg[0][0])
	yi := float64(seg[0][1])
	xf := float64(seg[1][0])
	yf := float64(seg[1][1])

	// se segmento horizontal, não tem interseção
	if yi == yf {
		return -1, -1
	}

	// troca para garantir ponto inicial em cima
	if yi > yf {
		xi, xf = xf, xi
		yi, yf = yf, yi
	}

	// calcula t
	t := (float64(scan) - yi) / (yf - yi)

	// calcula x
	if t > 0 && t <= 1 {
		return xi + t*(xf-xi), t
	}

	// sem interseção
	return -1, -1
}
