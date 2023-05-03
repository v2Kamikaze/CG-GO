package vec

type Vec2D struct {
	X, Y int
}

type VecTexture struct {
	Tx, Ty float64
}

func NewVec2(x, y int) Vec2D {
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

func (v Vec2D) ToXYZ() []int {
	return []int{v.X, v.Y, 1}
}

func (v Vec2D) ToTransposedXY1() [][]int {
	return [][]int{{v.X}, {v.Y}, {1}}
}

func (v Vec2D) ToTransposedFXY1() [][]float64 {
	return [][]float64{{float64(v.X)}, {float64(v.Y)}, {1}}
}
