package shapes

import (
	"cg-go/src/core/pixel"
	"cg-go/src/core/vec"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GeometricShape struct {
	Vertices        []vec.Vec2D
	ColorVertices   []color.RGBA
	TextureVertices []vec.VecTexture
	Texture         [][]color.RGBA
}

func (s *GeometricShape) DrawMesh(screen *ebiten.Image) {
	if len(s.Vertices) < 3 {
		return
	}

	xi := s.Vertices[0].X
	yi := s.Vertices[0].Y
	for i := 1; i < len(s.Vertices); i++ {
		xf := s.Vertices[i].X
		yf := s.Vertices[i].Y
		pixel.DrawLine(screen, xi, yi, xf, yf, color.RGBA{255, 255, 255, 255})
		xi = xf
		yi = yf
	}
	xf := s.Vertices[0].X
	yf := s.Vertices[0].Y
	pixel.DrawLine(screen, xi, yi, xf, yf, color.RGBA{255, 255, 255, 255})
}

func NewSquare(width, height int, center vec.Vec2D) *GeometricShape {
	return &GeometricShape{
		Vertices: []vec.Vec2D{
			vec.NewVec2(center.X-width, center.Y-height),
			vec.NewVec2(center.X+width, center.Y-height),
			vec.NewVec2(center.X+width, center.Y+height),
			vec.NewVec2(center.X-width, center.Y+height),
		},
	}
}
