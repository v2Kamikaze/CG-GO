package window

import (
	"cg-go/src/core/matrix"
	"cg-go/src/core/transform"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
)

type window struct {
	pi, pf vec.Vec2D
}

func New(pi, pf vec.Vec2D) *window {
	return &window{pi, pf}
}

func (w *window) Translate(dx, dy float64) {
	mtx := transform.NewTranslateMatrix(dx, dy)
	w.pi = transform.TranslatePoint(mtx, w.pi)
	w.pf = transform.TranslatePoint(mtx, w.pf)
}

func (w *window) MapPoints(s *shapes.GeometricShape, vpw, vph float64) {
	MapPointsToWindow(s, w.pi, w.pf, vpw, vph)
}

func MapPointToWindow(point vec.Vec2D, wi, wf vec.Vec2D, vpw, vph float64) vec.Vec2D {

	mtx := [][]float64{
		{vpw / (wf.X - wi.X), 0, (wi.X * vpw / (wf.X - wi.X))},
		{0, vph / (wf.Y - wi.Y), (wi.Y * vph / (wf.Y - wi.Y))},
		{0, 0, 1},
	}

	newPoint := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2(newPoint[0][0], newPoint[1][0])
}

func MapPointsToWindow(s *shapes.GeometricShape, wi, wf vec.Vec2D, vpw, vph float64) {
	for i, p := range s.Vertices {
		s.Vertices[i] = MapPointToWindow(p, wi, wf, vpw, vph)
	}
}
