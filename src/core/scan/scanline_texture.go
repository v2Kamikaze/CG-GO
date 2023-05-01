package scan

import (
	"cg-go/src/core/pixel"
	"cg-go/src/shapes"

	imgcolor "image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScanlineTexture(screen *ebiten.Image, s *shapes.GeometricShape, texture [][]imgcolor.RGBA) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var intersections []ScanlinePointTexture

		pi := s.Vertices[0]

		for p := 1; p < len(s.Vertices); p++ {
			pf := s.Vertices[p]

			point := IntersectionForTexture(y, [][]uint32{pi, pf})

			if point.X >= 0 {
				intersections = append(intersections, point)
			}

			pi = pf
		}

		pf := s.Vertices[0]

		point := IntersectionForTexture(y, [][]uint32{pi, pf})

		if point.X >= 0 {
			intersections = append(intersections, point)
		}

		for pi := 0; pi < len(intersections); pi += 2 {
			p1, p2 := intersections[pi], intersections[pi+1]

			if p2.X < p1.X {
				p1, p2 = p2, p1
			}

			for xk := p1.X; xk <= p2.X; xk++ {
				pc := float64(xk-p1.X) / float64(p2.X-p1.X)
				tx := p1.Tx + pc*(p2.Tx-p1.Tx)
				ty := p1.Ty + pc*(p2.Ty-p1.Ty)
				color := pixel.GetPixel(texture, tx, ty)
				pixel.DrawPixel(screen, xk, y, color)
			}
		}

	}

}
