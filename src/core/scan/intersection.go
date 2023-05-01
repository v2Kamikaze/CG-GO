package scan

import (
	"math"
)

func Intersection(y int, seg [2][2]int) (float64, float64) {
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
	var t float64 = (float64(y) - yi) / (yf - yi)

	// calcula x
	if t > 0 && t <= 1 {
		return xi + t*(xf-xi), t
	}

	// sem interseção
	return -1, -1
}

func IntersectionForTexture(y int, seg [][]uint32) ScanlinePointTexture {
	pi := seg[0]
	pf := seg[1]

	if pi[1] == pf[1] {
		return NewScanlinePointTexture(-1, 0, 0, 0)
	}

	if pi[1] > pf[1] {
		pi, pf = pf, pi
	}

	//fmt.Printf("Y %d | pi-X1 %d | pf-X2 %d", y, pi[1], pf[1])

	t := (float64(y) - float64(pi[1])) / float64(pf[1]-pi[1])
	//fmt.Println("t =======> ", t)

	if t > 0 && t <= 1 {
		x := math.Round(float64(pi[0]) + t*float64(float64(pf[0])-float64(pi[0])))
		tx := float64(pi[2]) + t*(float64(pf[2])-float64(pi[2]))
		ty := float64(pi[3]) + t*(float64(pf[3])-float64(pi[3]))

		return NewScanlinePointTexture(int(x), y, tx, ty)
	}

	return NewScanlinePointTexture(-1, 0, 0, 0)
}
