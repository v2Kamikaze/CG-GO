package core

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawLine(screen *ebiten.Image, x0, y0, x1, y1 int, c color.RGBA) {
	// Calcula as diferenças em x e y
	deltaX := math.Abs(float64(x1 - x0))
	deltaY := math.Abs(float64(y1 - y0))
	// Determina as direções da linha em x e y
	dirX := 1
	if x0 > x1 {
		dirX = -1
	}
	dirY := 1
	if y0 > y1 {
		dirY = -1
	}
	// Inicializa o erro
	error := deltaX - deltaY
	// Loop principal
	for x0 != x1 || y0 != y1 {
		// Desenha o pixel atual
		DrawPixel(screen, x0, y0, c)
		// Calcula o erro duplo
		error2 := error * 2
		// Verifica se deve mover em x
		if error2 > -deltaY {
			error -= deltaY
			x0 += dirX
		}
		// Verifica se deve mover em y
		if error2 < deltaX {
			error += deltaX
			y0 += dirY
		}
	}
}
