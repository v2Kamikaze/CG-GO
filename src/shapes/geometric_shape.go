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
		pixel.DrawLine(screen, int(xi), int(yi), int(xf), int(yf), color.RGBA{255, 255, 255, 255})
		xi = xf
		yi = yf
	}
	xf := s.Vertices[0].X
	yf := s.Vertices[0].Y
	pixel.DrawLine(screen, int(xi), int(yi), int(xf), int(yf), color.RGBA{255, 255, 255, 255})
}

func (s *GeometricShape) WithColors(colors []color.RGBA) *GeometricShape {
	s.ColorVertices = colors
	return s
}

func (s *GeometricShape) WithTexture(texture [][]color.RGBA) *GeometricShape {
	s.Texture = texture
	return s
}

func (s *GeometricShape) WithTextureVertices(textureVertices []vec.VecTexture) *GeometricShape {
	s.TextureVertices = textureVertices
	return s
}

func NewSquare(width, height float64, center vec.Vec2D) *GeometricShape {
	return &GeometricShape{
		Vertices: []vec.Vec2D{
			vec.NewVec2(center.X-width, center.Y-height),
			vec.NewVec2(center.X+width, center.Y-height),
			vec.NewVec2(center.X+width, center.Y+height),
			vec.NewVec2(center.X-width, center.Y+height),
		},
	}
}

func NewTriangle(base, height float64, center vec.Vec2D) *GeometricShape {
	return &GeometricShape{
		Vertices: []vec.Vec2D{
			vec.NewVec2(center.X, center.Y-height/2),
			vec.NewVec2(center.X+base/2, center.Y+height/2),
			vec.NewVec2(center.X-base/2, center.Y+height/2),
		},
	}
}

func Copy(s *GeometricShape) *GeometricShape {
	return &GeometricShape{
		Vertices:        s.Vertices,
		ColorVertices:   s.ColorVertices,
		TextureVertices: s.TextureVertices,
		Texture:         s.Texture,
	}
}
