package pixel

import (
	"cg-go/src/core/vec"
	"cg-go/src/memory"
	"image/color"
	"math"
)

func DrawLine(mem memory.Memory, p0 vec.Vec2D, p1 vec.Vec2D, c color.RGBA) {
	// Convertendo as coordenadas das Vec2D para inteiros
	x0, y0 := int(p0.X), int(p0.Y)
	x1, y1 := int(p1.X), int(p1.Y)

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
	err := deltaX - deltaY

	// Loop principal
	for {
		// Desenha o pixel atual
		mem.SetPixel(x0, y0, c)

		// Verifica se chegou ao fim da linha
		if x0 == x1 && y0 == y1 {
			break
		}

		// Calcula o erro duplo
		error2 := err * 2

		// Verifica se deve mover em x
		if error2 > -deltaY {
			err -= deltaY
			x0 += dirX
		}

		// Verifica se deve mover em y
		if error2 < deltaX {
			err += deltaX
			y0 += dirY
		}
	}
}
