package scan

type ScanlinePoint struct {
	Xi int
	T  float64
}

type ScanlinePointGradient struct {
	ScanlinePoint
	Hex uint32
}

type ScanlinePointTexture struct {
	X, Y   int
	Tx, Ty float64
}

func NewScanlinePoint(xi int, t float64) ScanlinePoint {
	return ScanlinePoint{xi, t}
}

func NewScanlinePointGradient(xi int, t float64, hex uint32) ScanlinePointGradient {
	return ScanlinePointGradient{NewScanlinePoint(xi, t), hex}
}

func NewScanlinePointTexture(x, y int, tx, ty float64) ScanlinePointTexture {
	return ScanlinePointTexture{x, y, tx, ty}
}
