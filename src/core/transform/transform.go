package transform

import (
	"cg-go/src/core/matrix"
	"cg-go/src/core/vec"
)

func TransformPoint(mtx [][]float64, point vec.Vec2D) vec.Vec2D {
	pointScaled := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2(pointScaled[0][0], pointScaled[1][0])
}
