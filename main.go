package main

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/pixel"
	"cg-go/src/core/scan"
	"cg-go/src/core/screen"
	"cg-go/src/core/transform"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"cg-go/src/window"
	"fmt"
	"image/color"
	"math"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var img, _ = image.ReadImage("./resources/gato.png")
var player = shapes.NewRect(1, 1, vec.NewVec2(3, 3)).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2(0, 0),
		vec.NewVec2(1, 0),
		vec.NewVec2(1, 1),
		vec.NewVec2(0, 1),
	}).WithTexture(img)

var rect = shapes.NewRect(1.5, 1.5, vec.NewVec2(3, 3))
var tri = shapes.NewTriangle(2, 3, vec.NewVec2(3, 3))

var vp = window.NewViewport(300, 300)
var win = window.New(vec.NewVec2(0, 0), vec.NewVec2(6, 6))

var mtx = NewScreenMat(int(vp.Width), int(vp.Height))

func Update(ctx *ebiten.Image) {
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[i]); j++ {
			SetPixel(mtx, j, i, colors.HexToRGBA(colors.Blue))
		}
	}

	//Clear(mtx)

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[i]); j++ {
			pixel.DrawPixel(ctx, i, j, mtx[i][j])
		}
	}

	scan.ScanlineBasic(ctx, rect, colors.HexToRGBA(colors.Yellow))
	rect.DrawMesh(ctx)
	scan.ScanlineBasic(ctx, tri, colors.HexToRGBA(colors.Pink))
	tri.DrawMesh(ctx)
	scan.ScanlineTexture(ctx, player, player.Texture)
	player.DrawMesh(ctx)

	transform.RotateVerticesOnPivot(-4, tri.Center(), tri)
	transform.RotateVerticesOnPivot(2, rect.Center(), rect)
	transform.RotateVerticesOnPivot(-2, player.Center(), player)
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

func NewScreenMat(width, height int) [][]color.RGBA {
	mtx := make([][]color.RGBA, width)

	for i := 0; i < width; i++ {
		mtx[i] = make([]color.RGBA, height)
	}

	return mtx
}

func ToString(mtx [][]color.RGBA) {
	fmt.Println(strings.Repeat("=", len(mtx)*20))
	for i := 0; i < len(mtx); i++ {
		fmt.Print("[")
		for j := 0; j < len(mtx[i]); j++ {
			fmt.Print(" ", mtx[i][j])
		}
		fmt.Println("]")
	}
	fmt.Println(strings.Repeat("=", len(mtx)*20))

}

func SetPixel(mtx [][]color.RGBA, x, y int, color color.RGBA) {
	mtx[y][x] = color
}

func Clear(mtx [][]color.RGBA) {
	wg := &sync.WaitGroup{}

	wg.Add(len(mtx))

	for i := 0; i < len(mtx); i++ {
		go func(x int, wg *sync.WaitGroup) {
			for j := 0; j < len(mtx[x]); j++ {
				SetPixel(mtx, x, j, color.RGBA{})
			}
			wg.Done()
		}(i, wg)
	}

	wg.Wait()
}
