package scan

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/pixel"
	"cg-go/src/geo"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScanlineGradient(screen *ebiten.Image, s *geo.GeometricShape) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []ScanlinePointGradient

		pi := s.Vertices[0]
		lastColor := s.ColorVertices[0]

		for p := 1; p < len(s.Vertices); p++ {
			pf := s.Vertices[p]
			currenColor := s.ColorVertices[p]

			point := Intersection(y, pi, pf)

			if point.X >= 0 {
				colorInterpolated := colors.InterpolateColors(lastColor, currenColor, point.T)
				i = append(i, NewScanlinePointGradient(point.X, point.T, colorInterpolated))
			}

			pi, pf, lastColor = pf, pi, currenColor
		}

		pf, currentColor := s.Vertices[0], s.ColorVertices[0]

		point := Intersection(y, pi, pf)

		if point.X >= 0 {
			colorInterpolated := colors.InterpolateColors(lastColor, currentColor, point.T)
			i = append(i, NewScanlinePointGradient(point.X, point.T, colorInterpolated))
		}

		for pi := 0; pi < len(i); pi += 2 {

			if len(i) == 1 {
				continue
			}

			x1, startColor := i[pi].X, i[pi].Color
			x2, endColor := i[pi+1].X, i[pi+1].Color

			if x2 < x1 {
				x1, x2 = x2, x1
				startColor, endColor = endColor, startColor
			}

			for xk := x1; xk <= x2; xk++ {
				ratio := float64(xk-x1) / float64(x2-x1)

				c := colors.InterpolateColors(startColor, endColor, ratio)
				pixel.DrawPixel(screen, xk, y, c)
			}
		}

	}

}
