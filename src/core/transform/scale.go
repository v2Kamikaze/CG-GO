package transform

import (
	"cg-go/src/core/matrix"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"math"
)

func NewScaleMatrix(sx, sy float64) [][]float64 {
	return [][]float64{{sx, 0, 0}, {0, sy, 0}, {0, 0, 1}}
}

func NewScaledTranslatedMatrix(dx, dy, sx, sy float64) [][]float64 {
	translateMatForward := NewTranslateMatrix(dx, dy)
	translateMatBack := NewTranslateMatrix(-dx, -dy)
	rawScaleMat := NewScaleMatrix(sx, sy)
	return matrix.MatrixMult(matrix.MatrixMult(translateMatForward, rawScaleMat), translateMatBack)
}

func ScalePoint(mtx [][]float64, point vec.Vec2D) vec.Vec2D {
	pointScaled := matrix.MatrixMult(mtx, point.ToTransposedFXY1())
	return vec.NewVec2(int(math.Round(pointScaled[0][0])), int(math.Round(pointScaled[1][0])))
}

func ScaleVertices(sx, sy float64, s *shapes.GeometricShape) {
	mtx := NewScaledTranslatedMatrix(float64(s.Vertices[0].X), float64(s.Vertices[0].Y), sx, sy)

	var scaled []vec.Vec2D

	for _, point := range s.Vertices {
		scaled = append(scaled, ScalePoint(mtx, point))
	}

	s.Vertices = scaled
}
