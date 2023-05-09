package main

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/screen"
	"cg-go/src/core/transform"
	"cg-go/src/core/vec"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/scan"
	"cg-go/src/window"

	"github.com/hajimehoshi/ebiten/v2"
)

var img, _ = image.ReadImage("./resources/gato.png")
var player = geo.NewRect(1, 1, vec.NewVec2(6, 6)).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(1, 0),
		vec.NewVec2(1, 1),
		vec.NewVec2(0, 1),
	})

const Width, Height = 400, 400

var rect = geo.NewRect(1.5, 1.5, vec.NewVec2(6, 6))
var tri = geo.NewTriangle(2, 3, vec.NewVec2(6, 6))
var sqr = geo.NewRect(20, 20, vec.NewVec2(150, 150))

var vp1 = window.NewViewport(vec.Zeros(), vec.NewVec2((Width-1)/2, (Height-1)/2))

var vp2 = window.NewViewport(vec.NewVec2((Width-1)/2, 0), vec.NewVec2(Width-1, (Height-1)/2))

var vp3 = window.NewViewport(vec.NewVec2(0, (Height-1)/2), vec.NewVec2((Width-1)/2, Height-1))

var vp4 = window.NewViewport(vec.NewVec2((Width-1)/2, (Height-1)/2), vec.NewVec2(Width-1, Height-1))

var win = window.New(vec.NewVec2(0, 0), vec.NewVec2(12, 12))
var mem = memory.New(Width, Height)

func Update(ctx *ebiten.Image) {
	mem.Clear()

	vp1.DrawBounds(mem)
	vp2.DrawBounds(mem)
	vp3.DrawBounds(mem)
	vp4.DrawBounds(mem)

	sqr.DrawBounds(mem)

	mapToVP(vp1)
	mapToVP(vp2)
	mapToVP(vp3)
	mapToVP(vp4)

	mem.Draw(ctx)

	transform.TranslateVertices(vec.NewVec2(0.5, 0), tri)
	transform.RotateVerticesOnPivot(-4, tri.Center(), tri)
	transform.RotateVerticesOnPivot(-4, player.Center(), player)
	transform.RotateVerticesOnPivot(4, player.Center(), rect)
	transform.RotateVerticesOnPivot(6, vec.NewVec2(Height/2, Width/2), sqr)
}

func main() {

	screen.New().
		SetWidth(mem.Width()).
		SetHeight(mem.Height()).
		SetTitle("Term").
		SetOnUpdate(Update).
		Build().
		Run()
}

func mapToVP(vp *window.Viewport) {
	rect.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineBasic(mem, s, colors.HexToRGBA(colors.Yellow))
	})

	tri.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineBasic(mem, s, colors.HexToRGBA(colors.Pink))

	})

	player.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, img)
	})
}
