package render

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/pixel"
	"cg-go/src/core/scan"
	"cg-go/src/core/transform"
	"cg-go/src/shapes"

	"github.com/hajimehoshi/ebiten/v2"
)

var pol = &shapes.GeometricShape{
	Vertices: [][]uint32{
		{5, 5, colors.Blue},
		{300, 300, colors.Red},
		{5, 300, colors.Red},
	},
}

var img, _ = image.ReadImage("./teste.jpg")

func Update(screen *ebiten.Image) {
	screen.Clear()

	for i := range img {
		for j := range img[i] {
			pixel.DrawPixel(screen, j+50, i+50, img[i][j])
		}

	}

	scan.Scanline(screen, pol)

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		transform.TranslatePolygon(2, 0, pol)
		scan.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		transform.TranslatePolygon(-2, 0, pol)
		scan.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		transform.TranslatePolygon(0, -2, pol)
		scan.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		transform.TranslatePolygon(0, 2, pol)
		scan.Scanline(screen, pol)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		transform.ScalePolygon(1.05, 1.05, pol)
		scan.Scanline(screen, pol)
	}
}
