package render

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var polImg = &shapes.GeometricShape{
	Vertices: []vec.Vec2D{
		vec.NewVec2D(0, 0),
		vec.NewVec2D(400, 0),
		vec.NewVec2D(380, 200),
		vec.NewVec2D(50, 200),
	},

	TextureVertices: []vec.VecTexture{
		vec.NewVecTexture(0, 0),
		vec.NewVecTexture(1, 0),
		vec.NewVecTexture(1, 1),
		vec.NewVecTexture(0, 1),
	},
}

var triangle = &shapes.GeometricShape{
	Vertices: []vec.Vec2D{
		vec.NewVec2D(200, 200),
		vec.NewVec2D(300, 300),
		vec.NewVec2D(100, 300),
	},
}

var square = &shapes.GeometricShape{
	Vertices: []vec.Vec2D{
		vec.NewVec2D(400, 400),
		vec.NewVec2D(1000, 400),
		vec.NewVec2D(1000, 600),
		vec.NewVec2D(400, 600),
	},
	ColorVertices: []color.RGBA{
		colors.HexToRGBA(colors.Purple),
		colors.HexToRGBA(colors.Yellow),
		colors.HexToRGBA(colors.Yellow),
		colors.HexToRGBA(colors.Purple),
	},
}

var img, _ = image.ReadImage("./resources/cat.jpg")

func Update(screen *ebiten.Image) {
	screen.Clear()

	scan.ScanlineTexture(screen, polImg, img)
	polImg.DrawMesh(screen)

	scan.ScanlineBasic(screen, triangle, colors.HexToRGBA(colors.Teal))
	triangle.DrawMesh(screen)

	scan.ScanlineGradient(screen, square)
	square.DrawMesh(screen)

}
