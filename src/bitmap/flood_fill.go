package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

func FloodFill(mem memory.Memory, point vec.Vec2D, colorBg, colorBorder color.RGBA) {

	// Cria uma matriz booleana para controlar os pixels preenchidos
	filled := make([][]bool, mem.Height())
	for i := range filled {
		filled[i] = make([]bool, mem.Width())
	}

	stack := []vec.Vec2D{point}

	// Preenche os pixels enquanto a pilha não estiver vazia
	for len(stack) > 0 {
		// Remove o último elemento da pilha
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Verifica se o pixel já foi preenchido
		if filled[int(p.Y)][int(p.X)] {
			continue
		}

		// Preenche o pixel e adiciona seus vizinhos à pilha se não forem da cor de borda
		fillPixel(mem, filled, p, colorBg)
		if !isBorderColor(mem, p.Plus(vec.NewVec2D(-1, 0)), colorBorder) {
			stack = append(stack, p.Plus(vec.NewVec2D(-1, 0)))
		}
		if !isBorderColor(mem, p.Plus(vec.NewVec2D(1, 0)), colorBorder) {
			stack = append(stack, p.Plus(vec.NewVec2D(1, 0)))
		}
		if !isBorderColor(mem, p.Plus(vec.NewVec2D(0, -1)), colorBorder) {
			stack = append(stack, p.Plus(vec.NewVec2D(0, -1)))
		}
		if !isBorderColor(mem, p.Plus(vec.NewVec2D(0, 1)), colorBorder) {
			stack = append(stack, p.Plus(vec.NewVec2D(0, 1)))
		}
	}
}

func fillPixel(mem memory.Memory, filled [][]bool, p vec.Vec2D, colorBg color.RGBA) {
	x, y := int(p.X), int(p.Y)
	if x < 0 || x >= mem.Height() || y < 0 || y >= mem.Width() {
		return
	}
	if !filled[y][x] {
		mem.SetPixel(x, y, colorBg)
		filled[y][x] = true
	}
}

func isBorderColor(mem memory.Memory, p vec.Vec2D, colorBorder color.RGBA) bool {
	x, y := int(p.X), int(p.Y)
	if x < 0 || x >= mem.Width() || y < 0 || y >= mem.Height() {
		return true
	}
	r, g, b, a := mem.GetPixel(x, y).RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)} == colorBorder
}
