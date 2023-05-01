package scan

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/pixel"
	"cg-go/src/shapes"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScanlineGradient(screen *ebiten.Image, s *shapes.GeometricShape) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []ScanlinePointGradient

		pix := int(s.Vertices[0][0])
		piy := int(s.Vertices[0][1])

		for p := 1; p < len(s.Vertices); p++ {
			pfx, pfy, hex := int(s.Vertices[p][0]), int(s.Vertices[p][1]), s.Vertices[p][2]

			xPoint, t := Intersection(y, [2][2]int{{(pix), (piy)}, {(pfx), (pfy)}})

			xi := int(math.Round(xPoint))

			if xi >= 0 {
				i = append(i, NewScanlinePointGradient(xi, t, hex))
			}

			pix, piy = pfx, pfy
		}

		lastPix, lastPiy, hex := int(s.Vertices[0][0]), int(s.Vertices[0][1]), s.Vertices[0][2]
		pfx, pfy := lastPix, lastPiy

		xPoint, t := Intersection(y, [2][2]int{{(pix), (piy)}, {(pfx), (pfy)}})

		xi := int(math.Round(xPoint))

		if xi >= 0 {
			i = append(i, NewScanlinePointGradient(xi, t, hex))
		}

		// ordenando os valores baseados na posição xi calculada
		for pi := 0; pi < len(i); pi += 2 {
			x1, t1, startColor := i[pi].Xi, i[pi].T, i[pi].Hex
			x2, t2, endColor := i[pi+1].Xi, i[pi+1].T, i[pi+1].Hex

			if x2 < x1 {
				x1, x2 = x2, x1
				t1, t2 = t2, t1
				startColor, endColor = endColor, startColor
			}

			for xk := x1; xk <= x2; xk++ {
				ratio := float64(xk-x1) / float64(x2-x1)
				smoothRatio := t1 + (t2-t1)*ratio

				c := colors.InterpolateColors(colors.HexToRGBA(startColor), colors.HexToRGBA(endColor), smoothRatio)
				pixel.DrawPixel(screen, xk, y, c)
			}
		}

	}

}
