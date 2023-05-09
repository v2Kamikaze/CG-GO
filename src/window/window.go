package window

import (
	"cg-go/src/core/matrix"
	"cg-go/src/core/vec"
	"cg-go/src/geo"
)

type Window struct {
	pi, pf vec.Vec2D
}

func New(pi, pf vec.Vec2D) *Window {
	return &Window{pi, pf}
}

func (w *Window) MapPoints(s *geo.GeometricShape, vp *Viewport) {
	MapPointsToWindow(s, w, vp)
}

func MapPointToWindow(point vec.Vec2D, wi, wf vec.Vec2D, vp *Viewport) vec.Vec2D {

	a := (vp.pf.X - vp.pi.X) / (wf.X - wi.X)
	b := (vp.pf.Y - vp.pi.Y) / (wf.Y - wi.Y)

	mtx := [][]float64{
		{a, 0, vp.pi.X - a*wi.X},
		{0, b, vp.pi.Y - b*wi.Y},
		{0, 0, 1},
	}

	newPoint := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2(newPoint[0][0], newPoint[1][0])
}

func MapPointsToWindow(s *geo.GeometricShape, w *Window, vp *Viewport) {
	for i, p := range s.Vertices {
		s.Vertices[i] = MapPointToWindow(p, w.pi, w.pf, vp)
	}
}
