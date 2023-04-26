package core

func ScalePoint(scaleMat [][]float64, point []uint32) []uint32 {
	x := float64(point[0])
	y := float64(point[1])

	pointTransposed := Transpose1D([]float64{x, y, 1})
	newPointTransposed := MatrixMult(scaleMat, pointTransposed)
	rawPoint := make([]uint32, 3)

	for i, row := range newPointTransposed {
		rawPoint[i] = uint32(row[0])
	}

	newPoint := make([]uint32, len(point))
	xn, yn := rawPoint[0], rawPoint[1]

	newPoint[0] = xn
	newPoint[1] = yn
	newPoint[2] = point[2] // Cor do ponto

	return newPoint
}

func ScalePolygon(sx, sy float64, shape *GeometricShape) {
	x, y := float64(shape.Vertices[0][0]), float64(shape.Vertices[0][1])

	translateMatForward := [][]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}

	translateMatBack := [][]float64{
		{1, 0, -x},
		{0, 1, -y},
		{0, 0, 1},
	}

	rawScaleMat := [][]float64{
		{sx, 0, 0},
		{0, sy, 0},
		{0, 0, 1},
	}

	scaleMat := MatrixMult(MatrixMult(translateMatForward, rawScaleMat), translateMatBack)

	scaled := make([][]uint32, len(shape.Vertices))

	for i, point := range shape.Vertices {
		scaled[i] = ScalePoint(scaleMat, point)
	}

	shape.Vertices = scaled
}
