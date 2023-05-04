package transform

import (
	"cg-go/src/core/matrix"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
)

func TranslatePoint(mtx [][]float64, point vec.Vec2D) vec.Vec2D {
	pointTranslated := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2(pointTranslated[0][0], pointTranslated[1][0])
}

func NewTranslateMatrix[T int | float64](dx, dy T) [][]T {
	return [][]T{
		{1, 0, dx},
		{0, 1, dy},
		{0, 0, 1},
	}
}

func TranslateVertices(delta vec.Vec2D, shape *shapes.GeometricShape) {
	translateMat := NewTranslateMatrix(delta.X, delta.Y)

	var translated []vec.Vec2D

	for _, point := range shape.Vertices {
		translated = append(translated, TranslatePoint(translateMat, point))
	}

	shape.Vertices = translated
}
