package render

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/shapes"

	"github.com/hajimehoshi/ebiten/v2"
)

var pol = &shapes.GeometricShape{
	Vertices: [][]uint32{
		{50, 50, 0, 0},
		{200, 60, 1, 0},
		{200, 200, 1, 1},
		{50, 200, 0, 1},
	},
}

var pol2 = &shapes.GeometricShape{
	Vertices: [][]uint32{
		{0, 0, colors.Blue},
		{50, 0, colors.Blue},
		{50, 50, colors.Red},
		{0, 50, colors.Red},
	},
}

var img, _ = image.ReadImage("./resources/cat.jpg")

func Update(screen *ebiten.Image) {
	screen.Clear()
	scan.ScanlineTexture(screen, pol, img)
	pol.DrawMesh(screen)

	scan.ScanlineGradient(screen, pol2)
	pol2.DrawMesh(screen)
}
