package scan

import (
	"cg-go/src/core/vec"
	"math"
)

func Intersection(y int, pi, pf vec.Vec2D) ScanlinePoint {

	// se segmento horizontal, não tem interseção
	if pi.Y == pf.Y {
		return NewScanlinePoint(-1, -1)
	}

	// troca para garantir ponto inicial em cima
	if pi.Y > pf.Y {
		pi, pf = pf, pi
	}

	// calcula t
	var t float64 = (float64(y) - float64(pi.Y)) / float64(pf.Y-pi.Y)

	// calcula x
	if t > 0 && t <= 1 {
		return NewScanlinePoint(int(math.Round(float64(pi.X)+t*float64(pf.X-pi.X))), t)
	}

	// sem interseção
	return NewScanlinePoint(-1, -1)
}

func IntersectionForTexture(y int, pi, pf vec.Vec2D, texturePi, texturePf vec.VecTexture) ScanlinePointTexture {

	if pi.Y == pf.Y {
		return NewScanlinePointTexture(-1, 0, 0, 0)
	}

	if pi.Y > pf.Y {
		pi, pf = pf, pi
		texturePi, texturePf = texturePf, texturePi
	}

	t := (float64(y) - float64(pi.Y)) / float64(pf.Y-pi.Y)

	if t > 0 && t <= 1 {
		x := math.Round(float64(pi.X) + t*float64(float64(pf.X)-float64(pi.X)))
		tx := texturePi.Tx + t*(texturePf.Tx-texturePi.Tx)
		ty := texturePi.Ty + t*(texturePf.Ty-texturePi.Ty)

		return NewScanlinePointTexture(int(x), y, tx, ty)
	}

	return NewScanlinePointTexture(-1, 0, 0, 0)
}
