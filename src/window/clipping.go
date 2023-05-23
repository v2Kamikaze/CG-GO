package window

import (
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/scan"
	"cg-go/src/vec"
	"fmt"
)

func ClipPolygon(mem memory.Memory, shape *geo.GeometricShape, vp *Viewport) *geo.GeometricShape {

	var (
		topLeft     = vp.pi
		topRight    = vec.NewVec2D(vp.pf.X, vp.pi.Y)
		bottomLeft  = vec.NewVec2D(vp.pi.X, vp.pf.Y)
		bottomRight = vp.pf
		s           = geo.Copy(shape)

		point vec.Vec2D
		pi    vec.Vec2D
		pf    vec.Vec2D
	)

	check := func() {
		// Checando se intersecta a parte superior da viewport
		point, _, _ = Intersection(pi, pf, topLeft, topRight)
		if point.X > 0 {
			fmt.Printf("Se chocou com a parte superior %+v %+v\n", point, pi)
			scan.ScanlineBasic(mem, geo.NewRect(5, 5, point), colors.ColorRed)
		}

		// Checando se intersecta a parte lateral direita da viewport
		point, _, _ = Intersection(pi, pf, topRight, bottomRight)
		if point.X > 0 {
			fmt.Printf("Se chocou com a parte lateral direita %+v %+v\n", point, pf)
			scan.ScanlineBasic(mem, geo.NewRect(5, 5, point), colors.ColorRed)

		}
		// Checando se intersecta a parte inferior da viewport
		point, _, _ = Intersection(pi, pf, bottomRight, bottomLeft)
		if point.X > 0 {
			fmt.Printf("Se chocou com a parte inferior %+v %+v\n", point, pf)
			scan.ScanlineBasic(mem, geo.NewRect(5, 5, point), colors.ColorRed)

		}
		// Checando se intersecta a parte lateral esquerda da viewport
		point, _, _ = Intersection(pi, pf, bottomLeft, topLeft)
		if point.X > 0 {
			fmt.Printf("Se chocou com a parte lateral direita %+v %+v\n", point, pf)
			scan.ScanlineBasic(mem, geo.NewRect(5, 5, point.Plus(vec.NewVec2D(10, 0))), colors.ColorRed)
		}
	}

	pi = s.Vertices[0]

	for i := 1; i < len(s.Vertices); i++ {
		pf = s.Vertices[i]

		check()

		pi = pf
	}

	pf = s.Vertices[0]

	check()

	return s
}
