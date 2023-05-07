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
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var img, _ = image.ReadImage("./resources/gato.png")
var player = geo.NewRect(1, 1, vec.NewVec2(3, 3)).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(1, 0),
		vec.NewVec2(1, 1),
		vec.NewVec2(0, 1),
	}).WithTexture(img)

var rectColor = geo.NewRect(20, 20, vec.NewVec2(150, 100)).WithColors(
	[]color.RGBA{
		colors.HexToRGBA(colors.Yellow),
		colors.HexToRGBA(colors.Purple),
		colors.HexToRGBA(colors.Purple),
		colors.HexToRGBA(colors.Yellow),
	},
)

var rect = geo.NewRect(1.5, 1.5, vec.NewVec2(3, 3))
var tri = geo.NewTriangle(2, 3, vec.NewVec2(3, 3))

var vp = window.NewViewport(500, 400)
var win = window.New(vec.NewVec2(0, 0), vec.NewVec2(6, 6))

var mem = memory.New(int(vp.Width), int(vp.Height))

var sqr = geo.NewRect(20, 20, vec.NewVec2(50, 50))

func Update(ctx *ebiten.Image) {
	mem.Clear()

	sqr.DrawBounds(mem)

	scan.ScanlineBasic(mem, rect, colors.HexToRGBA(colors.Yellow))
	rect.DrawBounds(mem)

	scan.ScanlineBasic(mem, tri, colors.HexToRGBA(colors.Pink))
	tri.DrawBounds(mem)

	scan.ScanlineTexture(mem, player, player.Texture)
	player.DrawBounds(mem)

	scan.ScanlineGradient(mem, rectColor)
	rectColor.DrawBounds(mem)

	mem.Draw(ctx)

	transform.TranslateVertices(vec.NewVec2(0.5, 0), tri)
	transform.RotateVerticesOnPivot(-4, tri.Center(), tri)
	transform.RotateVerticesOnPivot(-4, player.Center(), player)
	transform.RotateVerticesOnPivot(4, player.Center(), rect)
	transform.RotateVerticesOnPivot(6, player.Center(), sqr)

}

func main() {
	win.MapPoints(player, vp)
	win.MapPoints(rect, vp)
	win.MapPoints(tri, vp)

	screen.New().
		SetWidth(int(math.Round(vp.Width))).
		SetHeight(int(math.Round(vp.Height))).
		SetTitle("Term").
		SetOnUpdate(Update).
		Build().
		Run()
}
