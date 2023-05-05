package pixel

import (
	"cg-go/src/core/vec"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func FloodFill(ctx *ebiten.Image, point vec.Vec2D, colorBg, colorBorder color.RGBA) {
	width, height := ctx.Bounds().Max.X, ctx.Bounds().Max.Y

	// Cria uma matriz booleana para controlar os pixels preenchidos
	filled := make([][]bool, height)
	for i := range filled {
		filled[i] = make([]bool, width)
	}

	// Define uma função auxiliar para preencher um pixel com a cor de fundo
	fillPixel := func(p vec.Vec2D) {
		x, y := int(p.X), int(p.Y)
		if x < 0 || x >= width || y < 0 || y >= height {
			return
		}
		if !filled[y][x] {
			DrawPixel(ctx, x, y, colorBg)
			filled[y][x] = true
		}
	}

	// Define uma função para verificar se um pixel é a cor de borda
	isBorderColor := func(p vec.Vec2D) bool {
		x, y := int(p.X), int(p.Y)
		if x < 0 || x >= width || y < 0 || y >= height {
			return true
		}
		r, g, b, a := ctx.At(x, y).RGBA()
		return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)} == colorBorder
	}

	// Inicia a pilha com o ponto inicial
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
		fillPixel(p)
		if !isBorderColor(p.Plus(vec.NewVec2(-1, 0))) {
			stack = append(stack, p.Plus(vec.NewVec2(-1, 0)))
		}
		if !isBorderColor(p.Plus(vec.NewVec2(1, 0))) {
			stack = append(stack, p.Plus(vec.NewVec2(1, 0)))
		}
		if !isBorderColor(p.Plus(vec.NewVec2(0, -1))) {
			stack = append(stack, p.Plus(vec.NewVec2(0, -1)))
		}
		if !isBorderColor(p.Plus(vec.NewVec2(0, 1))) {
			stack = append(stack, p.Plus(vec.NewVec2(0, 1)))
		}
	}
}
