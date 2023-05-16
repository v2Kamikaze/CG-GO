package main

import (
	"cg-go/src/bitmap"
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/scan"
	"cg-go/src/screen"
	"cg-go/src/tex"
	"cg-go/src/transform"
	"cg-go/src/vec"
	"cg-go/src/window"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const Width, Height = 400, 400

var win = window.New(vec.NewVec2D(0, 0), vec.NewVec2D(8, 8))
var mem = memory.New(Width, Height)

var center = win.Center()

var img, _ = tex.ReadImage("./resources/gopher.png")
var player = geo.NewRect(1, 1, center).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2D(0, 0),
		vec.NewVec2D(1, 0),
		vec.NewVec2D(1, 1),
		vec.NewVec2D(0, 1),
	})

var rect = geo.NewRect(1.5, 1.5, center).
	WithColors([]color.RGBA{
		colors.HexToRGBA(colors.Blue),
		colors.HexToRGBA(colors.Red),
		colors.HexToRGBA(colors.Red),
		colors.HexToRGBA(colors.Blue),
	})

var tri = geo.NewTriangle(2, 3, center)

var vp1 = window.NewViewport(vec.Zeros(), vec.NewVec2D((Width-1)/2, (Height-1)/2))
var vp2 = window.NewViewport(vec.NewVec2D((Width-1)/2, 0), vec.NewVec2D(Width-1, (Height-1)/2))
var vp3 = window.NewViewport(vec.NewVec2D(0, (Height-1)/2), vec.NewVec2D((Width-1)/2, Height-1))
var vp4 = window.NewViewport(vec.NewVec2D((Width-1)/2, (Height-1)/2), vec.NewVec2D(Width-1, Height-1))

func Update(ctx *ebiten.Image) {
	mem.Clear()

	vp1.DrawBounds(mem)
	vp2.DrawBounds(mem)
	vp3.DrawBounds(mem)
	vp4.DrawBounds(mem)

	mapToVP(vp1)
	mapToVP(vp2)
	mapToVP(vp3)
	mapToVP(vp4)

	bitmap.BresenhamDrawFilledCircle(mem, vec.NewVec2D(Height/2, Width/2), 35.0, colors.HexToRGBA(colors.Red), colors.HexToRGBA(colors.Silver))

	mem.Draw(ctx)

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		win.Zoom(0.95)
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) {
		win.Zoom(1.05)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		win.Translate(vec.NewVec2D(0.5, 0))
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		win.Translate(vec.NewVec2D(-0.5, 0))
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		win.Translate(vec.NewVec2D(0, 0.5))
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		win.Translate(vec.NewVec2D(0, -0.5))
	}

	transform.TranslateVertices(vec.NewVec2D(-0.1, 0), tri)
	transform.RotateVerticesOnPivot(-4, tri.Center(), tri)
	transform.RotateVerticesOnPivot(-4, player.Center(), player)
	transform.RotateVerticesOnPivot(1, player.Center(), rect)
}

func main() {
	screen.New().
		SetWidth(mem.Width()).
		SetHeight(mem.Height()).
		SetTitle("CG").
		SetOnUpdate(Update).
		Build().
		Run()

}

func mapToVP(vp *window.Viewport) {
	rect.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineGradient(mem, s)
	})

	tri.Apply(func(s *geo.GeometricShape) {
		s = window.ClipPolygon(s, vp, win)
		win.MapPoints(s, vp)
		scan.ScanlineBasic(mem, s, colors.HexToRGBA(colors.Pink))

	})

	player.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, img)
	})

}
