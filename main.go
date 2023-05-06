package main

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/core/screen"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"cg-go/src/window"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var img, _ = image.ReadImage("./resources/gato.png")
var player = shapes.NewRect(1, 1, vec.NewVec2(1.5, 1.5)).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(1, 0),
		vec.NewVec2(1, 1),
		vec.NewVec2(0, 1),
	}).WithTexture(img)

var rect = shapes.NewRect(1.5, 1.5, vec.NewVec2(1.5, 1.5))

var vp = window.NewViewport(400, 200)
var win = window.New(vec.NewVec2(0, 0), vec.NewVec2(3, 3))

func Update(ctx *ebiten.Image) {

	scan.ScanlineBasic(ctx, rect, colors.HexToRGBA(colors.Yellow))
	rect.DrawMesh(ctx)
	scan.ScanlineTexture(ctx, player, player.Texture)
	player.DrawMesh(ctx)

}

func main() {
	win.MapPoints(player, vp)
	win.MapPoints(rect, vp)

	screen.New().
		SetWidth(int(math.Round(vp.Width))).
		SetHeight(int(math.Round(vp.Height))).
		SetTitle("Term").
		SetOnUpdate(Update).
		Build().
		Run()
}
