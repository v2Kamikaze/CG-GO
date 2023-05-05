package render

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/core/transform"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var polImg = &shapes.GeometricShape{
	Vertices: []vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(200, 0),
		vec.NewVec2(180, 200),
		vec.NewVec2(50, 200),
	},

	TextureVertices: []vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(1, 0),
		vec.NewVec2(1, 1),
		vec.NewVec2(0, 1),
	},
}

var triangle = &shapes.GeometricShape{
	Vertices: []vec.Vec2D{
		vec.NewVec2(200, 200),
		vec.NewVec2(300, 300),
		vec.NewVec2(100, 300),
	},
}

var square = &shapes.GeometricShape{
	Vertices: []vec.Vec2D{
		vec.NewVec2(300, 300),
		vec.NewVec2(400, 300),
		vec.NewVec2(400, 400),
		vec.NewVec2(300, 400),
	},
	ColorVertices: []color.RGBA{
		colors.HexToRGBA(colors.Yellow),
		colors.HexToRGBA(colors.Purple),
		colors.HexToRGBA(colors.Purple),
		colors.HexToRGBA(colors.Yellow),
	},
}

var square2 = shapes.NewRect(100, 100, vec.NewVec2(50, 100))

var img, _ = image.ReadImage("./resources/cat.jpg")

const velocity = 5

func Update(screen *ebiten.Image) {
	screen.Clear()

	scan.ScanlineBasic(screen, triangle, colors.HexToRGBA(colors.Teal))
	triangle.DrawMesh(screen)

	scan.ScanlineGradient(screen, square)
	square.DrawMesh(screen)

	square2.DrawMesh(screen)

	scan.ScanlineTexture(screen, polImg, img)
	polImg.DrawMesh(screen)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		transform.ScaleVertices(1.01, 1.01, polImg)
		scan.ScanlineTexture(screen, polImg, img)
		polImg.DrawMesh(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		transform.TranslateVertices(vec.NewVec2(-velocity, 0), polImg)
		scan.ScanlineTexture(screen, polImg, img)
		polImg.DrawMesh(screen)

	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		transform.TranslateVertices(vec.NewVec2(velocity, 0), polImg)
		scan.ScanlineTexture(screen, polImg, img)
		polImg.DrawMesh(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		transform.TranslateVertices(vec.NewVec2(0, -velocity), polImg)
		scan.ScanlineTexture(screen, polImg, img)
		polImg.DrawMesh(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		transform.TranslateVertices(vec.NewVec2(0, velocity), polImg)
		scan.ScanlineTexture(screen, polImg, img)
		polImg.DrawMesh(screen)
	}

}
