package scan

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/pixel"
	"cg-go/src/shapes"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type ScanlinePointType struct {
	Xi  int
	T   float64
	Hex uint32
}

func Scanline(screen *ebiten.Image, s *shapes.GeometricShape) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin := math.MaxUint32
	ymax := 0
	for _, p := range s.Vertices {
		if int(p[1]) < ymin {
			ymin = int(p[1])
		}
		if int(p[1]) > ymax {
			ymax = int(p[1])
		}
	}

	for y := ymin; y <= ymax; y++ {
		var i []ScanlinePointType

		pix := int(s.Vertices[0][0])
		piy := int(s.Vertices[0][1])

		for p := 1; p < len(s.Vertices); p++ {
			pfx := int(s.Vertices[p][0])
			pfy := int(s.Vertices[p][1])
			hex := s.Vertices[p][2]

			xPoint, t := Intersection(y, [][]int{{(pix), (piy)}, {(pfx), (pfy)}})

			xi := int(math.Round(xPoint))

			if xi >= 0 {
				i = append(i, ScanlinePointType{xi, t, hex})
			}

			pix = pfx
			piy = pfy
		}

		lastPix := int(s.Vertices[0][0])
		lastPiy := int(s.Vertices[0][1])
		hex := s.Vertices[0][2]
		pfx := lastPix
		pfy := lastPiy

		xPoint, t := Intersection(y, [][]int{{(pix), (piy)}, {(pfx), (pfy)}})

		xi := int(math.Round(xPoint))

		if xi >= 0 {
			i = append(i, ScanlinePointType{xi, t, hex})
		}

		// ordenando os valores baseados na posição xi calculada
		for pi := 0; pi < len(i); pi += 2 {
			x1 := i[pi].Xi
			t1 := i[pi].T
			startColor := i[pi].Hex
			x2 := i[pi+1].Xi
			t2 := i[pi+1].T
			endColor := i[pi+1].Hex

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
				//fmt.Printf("x: %d | y: %d | color: %x | ratio: %f\n", xk, y, c, smoothRatio)
			}
		}

	}

}
