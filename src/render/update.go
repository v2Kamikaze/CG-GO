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
		{50, 50, colors.Red},
		{5, 50, colors.Red},
	},
}

var img, _ = image.ReadImage("./resources/teste.jpg")
var pos = 50

func Update(screen *ebiten.Image) {
	screen.Clear()

	for i := range img {
		for j := range img[i] {
			pixel.DrawPixel(screen, j+pos, i+pos, img[i][j])
		}

	}

	scan.Scanline(screen, pol)
	pol.DrawMesh(screen)

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

}
