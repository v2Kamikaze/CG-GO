package world

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/core/vec"
	"cg-go/src/core/window"
	"cg-go/src/shapes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type world struct {
	screen  *window.Window
	objects []*shapes.GeometricShape
}

func New() *world {
	return &world{}
}

func (w *world) Create() {
	w.objects = []*shapes.GeometricShape{
		{
			Vertices: []vec.Vec2D{
				vec.NewVec2(0, 0),
				vec.NewVec2(200, 0),
				vec.NewVec2(180, 200),
				vec.NewVec2(50, 200),
			},

			TextureVertices: []vec.VecTexture{
				vec.NewVecTexture(0, 0),
				vec.NewVecTexture(1, 0),
				vec.NewVecTexture(1, 1),
				vec.NewVecTexture(0, 1),
			},
			Texture: img,
		},
		{
			Vertices: []vec.Vec2D{
				vec.NewVec2(200, 200),
				vec.NewVec2(300, 300),
				vec.NewVec2(100, 300),
			},
		},
		{
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
		},
		shapes.NewSquare(100, 100, vec.NewVec2(50, 100)),
	}

	w.screen = window.NewWindow().
		SetWidth(540).
		SetHeight(360).
		SetTitle("Term").
		SetOnUpdate(w.Update).
		Build()

	w.screen.Run()
}

var img, _ = image.ReadImage("./resources/cat.jpg")

func (w *world) Update(screen *ebiten.Image) {
	screen.Clear()

	for i := range w.objects {
		if len(w.objects[i].Texture) != 0 {
			scan.ScanlineTexture(screen, w.objects[i], w.objects[i].Texture)
			w.objects[i].DrawMesh(screen)
		} else if len(w.objects[i].ColorVertices) != 0 {
			scan.ScanlineGradient(screen, w.objects[i])
		} else {
			scan.ScanlineBasic(screen, w.objects[i], colors.HexToRGBA(colors.Purple))
		}
	}
	/*
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			transform.ScaleVertices(1.01, 1.01, polImg)
			scan.ScanlineTexture(screen, polImg, img)
			polImg.DrawMesh(screen)
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) {
			transform.TranslateVertices(-velocity, 0, polImg)
			scan.ScanlineTexture(screen, polImg, img)
			polImg.DrawMesh(screen)

		}

		if ebiten.IsKeyPressed(ebiten.KeyD) {
			transform.TranslateVertices(velocity, 0, polImg)
			scan.ScanlineTexture(screen, polImg, img)
			polImg.DrawMesh(screen)
		}

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			transform.TranslateVertices(0, -velocity, polImg)
			scan.ScanlineTexture(screen, polImg, img)
			polImg.DrawMesh(screen)
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) {
			transform.TranslateVertices(0, velocity, polImg)
			scan.ScanlineTexture(screen, polImg, img)
			polImg.DrawMesh(screen)
		} */

}
