package world

import (
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"cg-go/src/window"

	"github.com/hajimehoshi/ebiten/v2"
)

var img, _ = image.ReadImage("./resources/gato.png")
var player = shapes.NewRect(200, 100, vec.NewVec2(300, 300)).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(1, 0),
		vec.NewVec2(1, 1),
		vec.NewVec2(0, 1),
	}).WithTexture(img)

var win = window.New(vec.NewVec2(50, 50), vec.NewVec2(600, 600))

func Update(ctx *ebiten.Image) {

	scan.ScanlineTexture(ctx, player, player.Texture)
	player.DrawMesh(ctx)

	win.DrawCameraBounds(ctx)

}
