package lifecycle

import (
	"cg-go/src/core"

	"github.com/hajimehoshi/ebiten/v2"
)

var pol = &core.GeometricShape{
	Vertices: [][]uint32{
		{5, 5, core.Blue},
		{50, 50, core.Red},
		{5, 50, core.Red},
	},
}

func Update(screen *ebiten.Image) {
	screen.Clear()
	core.Scanline(screen, pol)

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		core.TranslatePolygon(2, 0, pol)
		core.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		core.TranslatePolygon(-2, 0, pol)
		core.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		core.TranslatePolygon(0, -2, pol)
		core.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		core.TranslatePolygon(0, 2, pol)
		core.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		core.ScalePolygon(1.05, 1.05, pol)
		core.Scanline(screen, pol)
	}
}
