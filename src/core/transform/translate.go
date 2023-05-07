package transform

import (
	"cg-go/src/core/vec"
	"cg-go/src/geo"
)

func NewTranslateMatrix[T int | float64](dx, dy T) [][]T {
	return [][]T{
		{1, 0, dx},
		{0, 1, dy},
		{0, 0, 1},
	}
}

func TranslateVertices(delta vec.Vec2D, shape *geo.GeometricShape) {
	translateMat := NewTranslateMatrix(delta.X, delta.Y)

	var translated []vec.Vec2D

	for _, point := range shape.Vertices {
		translated = append(translated, TransformPoint(translateMat, point))
	}

	shape.Vertices = translated
}
