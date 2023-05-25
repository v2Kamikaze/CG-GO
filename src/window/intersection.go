package window

import (
	"cg-go/src/vec"
)

func Intersection(pi, pf, linePi, linePf vec.Vec2D) (vec.Vec2D, float64, float64) {
	// Cálculo dos vetores de direção dos segmentos
	dir1 := vec.NewVec2D(pf.X-pi.X, pf.Y-pi.Y)
	dir2 := vec.NewVec2D(linePf.X-linePi.X, linePf.Y-linePi.Y)

	// Cálculo do determinante
	det := dir1.X*dir2.Y - dir1.Y*dir2.X

	// Verifica se os segmentos são paralelos ou coincidentes
	if pi.Y == pf.Y {
		return vec.NewVec2D(-1, -1), 0, 0
	}

	// Cálculo dos vetores entre os pontos iniciais dos segmentos
	startDiff := linePi.Sub(pi)

	// Cálculo dos parâmetros de interseção
	t := (startDiff.X*dir2.Y - startDiff.Y*dir2.X) / det
	u := (startDiff.X*dir1.Y - startDiff.Y*dir1.X) / det

	// Verifica se a interseção ocorre dentro dos segmentos
	if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
		// Cálculo das coordenadas do ponto de interseção
		intersectionX := pi.X + t*dir1.X
		intersectionY := pi.Y + t*dir1.Y
		return vec.NewVec2D(intersectionX, intersectionY), t, u
	}

	// Não há interseção dentro dos segmentos
	return vec.NewVec2D(-1, -1), 0, 0
}

func CohenSutherland() {

}
