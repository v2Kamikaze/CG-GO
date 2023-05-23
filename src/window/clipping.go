package window

import (
	"cg-go/src/bitmap"
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/vec"
	"fmt"
	"math"
)

const (
	top    = 1 << 0
	bottom = 1 << 1
	right  = 1 << 2
	left   = 1 << 3
)

func ClipPolygon(mem memory.Memory, original *geo.GeometricShape, vp *Viewport) *geo.GeometricShape {
	var (
		point vec.Vec2D
		pi    vec.Vec2D
		pf    vec.Vec2D
		t     float64
		u     float64
	)

	poly := geo.Copy(original)

	pi = poly.Vertices[1]

	pf = poly.Vertices[2]

	vpPoint := vec.NewVec2D(vp.pf.X, vp.pi.Y)
	bitmap.BresenhamLine(mem, vp.pi, vpPoint, colors.ColorWhite)

	point, t, u = Intersection(pi, pf, vp.pi, vpPoint)
	fmt.Printf("Intersection: (%f, %f), t: %f u: %f Reta: (%f, %f)\r ", point.X, point.Y, t, u, vpPoint.X, vpPoint.Y)

	return poly
}

func Intersection(seg1Start, seg1End, seg2Start, seg2End vec.Vec2D) (vec.Vec2D, float64, float64) {
	// Cálculo dos vetores de direção dos segmentos
	dir1 := vec.NewVec2D(seg1End.X-seg1Start.X, seg1End.Y-seg1Start.Y)
	dir2 := vec.NewVec2D(seg2End.X-seg2Start.X, seg2End.Y-seg2Start.Y)

	// Cálculo do determinante
	det := dir1.X*dir2.Y - dir1.Y*dir2.X

	// Verifica se os segmentos são paralelos ou coincidentes
	if math.Abs(det) < 1e-10 {
		return vec.NewVec2D(-1, -1), 0, 0
	}

	// Cálculo dos vetores entre os pontos iniciais dos segmentos
	startDiff := seg2Start.Sub(seg1Start)

	// Cálculo dos parâmetros de interseção
	t := (startDiff.X*dir2.Y - startDiff.Y*dir2.X) / det
	u := (startDiff.X*dir1.Y - startDiff.Y*dir1.X) / det

	// Verifica se a interseção ocorre dentro dos segmentos
	if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
		// Cálculo das coordenadas do ponto de interseção
		intersectionX := seg1Start.X + t*dir1.X
		intersectionY := seg1Start.Y + t*dir1.Y
		return vec.NewVec2D(intersectionX, intersectionY), t, u
	}

	// Não há interseção dentro dos segmentos
	return vec.NewVec2D(-1, -1), 0, 0
}
