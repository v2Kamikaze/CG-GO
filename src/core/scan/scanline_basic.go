package scan

import (
	"cg-go/src/core/pixel"
	"cg-go/src/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScanlineBasic(screen *ebiten.Image, s *shapes.GeometricShape, color color.RGBA) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []int

		pi := s.Vertices[0]

		for p := 1; p < len(s.Vertices); p++ {
			pf := s.Vertices[p]

			point := Intersection(y, pi, pf)

			if point.X >= 0 {
				i = append(i, point.X)
			}

			pi, pf = pf, pi
		}

		pf := s.Vertices[0]

		point := Intersection(y, pi, pf)

		if point.X >= 0 {
			i = append(i, point.X)
		}

		for pi := 0; pi < len(i); pi += 2 {
			x1, x2 := i[pi], i[pi+1]

			if x2 < x1 {
				x1, x2 = x2, x1
			}

			for xk := x1; xk <= x2; xk++ {
				pixel.DrawPixel(screen, xk, y, color)
			}
		}

	}

}
