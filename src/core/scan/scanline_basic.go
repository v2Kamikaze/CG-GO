package scan

import (
	"cg-go/src/core/pixel"
	"cg-go/src/shapes"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScanlineBasic(screen *ebiten.Image, s *shapes.GeometricShape, color color.RGBA) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []int

		pix := int(s.Vertices[0][0])
		piy := int(s.Vertices[0][1])

		for p := 1; p < len(s.Vertices); p++ {
			pfx, pfy := int(s.Vertices[p][0]), int(s.Vertices[p][1])

			xPoint, _ := Intersection(y, [2][2]int{{(pix), (piy)}, {(pfx), (pfy)}})

			xi := int(math.Round(xPoint))

			if xi >= 0 {
				i = append(i, xi)
			}

			pix, piy = pfx, pfy
		}

		lastPix, lastPiy := int(s.Vertices[0][0]), int(s.Vertices[0][1])
		pfx, pfy := lastPix, lastPiy

		xPoint, _ := Intersection(y, [2][2]int{{(pix), (piy)}, {(pfx), (pfy)}})

		xi := int(math.Round(xPoint))

		if xi >= 0 {
			i = append(i, xi)
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
