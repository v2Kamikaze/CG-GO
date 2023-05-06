package window

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/matrix"
	"cg-go/src/core/pixel"
	"cg-go/src/core/transform"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	pi, pf vec.Vec2D
}

func New(pi, pf vec.Vec2D) *Window {
	return &Window{pi, pf}
}

func (w *Window) DrawCameraBounds(ctx *ebiten.Image) {
	width := w.pf.X - w.pi.X
	height := w.pf.Y - w.pi.Y

	pixel.DrawLine(ctx, int(w.pi.X), int(w.pi.Y), int(w.pi.X+width), int(w.pi.Y), colors.HexToRGBA(colors.White))
	pixel.DrawLine(ctx, int(w.pi.X), int(w.pi.Y+height), int(w.pi.X+width), int(w.pi.Y+height), colors.HexToRGBA(colors.White))

	pixel.DrawLine(ctx, int(w.pi.X), int(w.pi.Y), int(w.pi.X), int(w.pi.Y+height), colors.HexToRGBA(colors.White))
	pixel.DrawLine(ctx, int(w.pi.X+width), int(w.pi.Y), int(w.pi.X+width), int(w.pi.Y+height), colors.HexToRGBA(colors.White))
}

func (w *Window) Translate(delta vec.Vec2D) {
	mtx := transform.NewTranslateMatrix(delta.X, delta.Y)
	w.pi = transform.TranslatePoint(mtx, w.pi)
	w.pf = transform.TranslatePoint(mtx, w.pf)
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
