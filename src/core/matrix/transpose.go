package matrix

func Transpose1D[T uint32 | float64](a []T) [][]T {
	transposed := make([][]T, len(a))
	for i, e := range a {
		transposed[i] = []T{e}
	}
	return transposed
}

func Transpose(a [][]uint32) [][]uint32 {
	rows := len(a)
	cols := len(a[0])
	transposed := make([][]uint32, cols)
	for j := 0; j < cols; j++ {
		transposed[j] = make([]uint32, rows)
		for i := 0; i < rows; i++ {
			transposed[j][i] = a[i][j]
		}
	}
	return transposed
}
