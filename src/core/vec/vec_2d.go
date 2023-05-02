package vec

type Vec2D struct {
	X, Y int
}

type VecTexture struct {
	Tx, Ty float64
}

func NewVec2D(x, y int) Vec2D {
	return Vec2D{x, y}
}

func NewVecTexture(tx, ty float64) VecTexture {
	return VecTexture{tx, ty}
}

func Zeros() Vec2D {
	return Vec2D{0, 0}
}

func Ones() Vec2D {
	return Vec2D{1, 1}
}
