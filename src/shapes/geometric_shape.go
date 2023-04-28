package shapes

import (
	"cg-go/src/core/pixel"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GeometricShape struct {
	Vertices [][]uint32
}

func (s *GeometricShape) DrawMesh(screen *ebiten.Image) {
	if len(s.Vertices) < 3 {
		return
	}

	xi := int(s.Vertices[0][0])
	yi := int(s.Vertices[0][1])
	for i := 1; i < len(s.Vertices); i++ {
		xf := int(s.Vertices[i][0])
		yf := int(s.Vertices[i][1])
		pixel.DrawLine(screen, xi, yi, xf, yf, color.RGBA{255, 255, 255, 255})
		xi = xf
		yi = yf
	}
	xf := int(s.Vertices[0][0])
	yf := int(s.Vertices[0][1])
	pixel.DrawLine(screen, xi, yi, xf, yf, color.RGBA{255, 255, 255, 255})
}
