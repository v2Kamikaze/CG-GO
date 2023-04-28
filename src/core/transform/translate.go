package transform

import (
	"cg-go/src/core/matrix"
	"cg-go/src/shapes"
	"math"
)

func TranslatePoint(translateMat [][]uint32, point []uint32) []uint32 {

	pointTransposed := matrix.Transpose1D([]uint32{point[0], point[1], 1})

	newPointTransposed := matrix.MatrixMult(translateMat, pointTransposed)
	rawPoint := []uint32{}

	for _, row := range newPointTransposed {
		rawPoint = append(rawPoint, row...)
	}

	newPoint := []uint32{0, 0, point[2]}
	x, y := float64(rawPoint[0]), float64(rawPoint[1])

	newPoint[0] = uint32(math.Round(x))
	newPoint[1] = uint32(math.Round(y))

	return newPoint
}

func TranslatePolygon(dx, dy int, shape *shapes.GeometricShape) {
	translateMat := [][]uint32{
		{1, 0, uint32(dx)},
		{0, 1, uint32(dy)},
		{0, 0, 1},
	}

	var translated [][]uint32
	for _, point := range shape.Vertices {
		translated = append(translated, TranslatePoint(translateMat, point))
	}
	shape.Vertices = translated
}
