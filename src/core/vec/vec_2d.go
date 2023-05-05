package vec

type Vec2D struct {
	X, Y float64
}

type VecTexture struct {
	Tx, Ty float64
}

func NewVec2(x, y float64) Vec2D {
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

func (v Vec2D) ToXYZ() []float64 {
	return []float64{v.X, v.Y, 1}
}

func (v Vec2D) ToTransposedXY1() [][]float64 {
	return [][]float64{{v.X}, {v.Y}, {1}}
}

func (v Vec2D) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec2D) Plus(vec Vec2D) Vec2D {
	return NewVec2(v.X+vec.X, v.Y+vec.Y)
}
