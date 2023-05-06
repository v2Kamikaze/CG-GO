package window

import (
	"cg-go/src/core/matrix"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
)

type Window struct {
	pi, pf vec.Vec2D
}

func New(pi, pf vec.Vec2D) *Window {
	return &Window{pi, pf}
}

func (w *Window) MapPoints(s *shapes.GeometricShape, vp *Viewport) {
	MapPointsToWindow(s, w.pi, w.pf, vp)
}

func MapPointToWindow(point vec.Vec2D, wi, wf vec.Vec2D, vp *Viewport) vec.Vec2D {
	mtx := [][]float64{
		{vp.Width / (wf.X - wi.X), 0, (wi.X * vp.Width / (wf.X - wi.X))},
		{0, vp.Height / (wf.Y - wi.Y), (wi.Y * vp.Height / (wf.Y - wi.Y))},
		{0, 0, 1},
	}

	newPoint := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2(newPoint[0][0], newPoint[1][0])
}

func MapPointsToWindow(s *shapes.GeometricShape, wi, wf vec.Vec2D, vp *Viewport) {
	for i, p := range s.Vertices {
		s.Vertices[i] = MapPointToWindow(p, wi, wf, vp)
	}
}
