package world

import (
	"cg-go/src/core/colors"
	"cg-go/src/core/image"
	"cg-go/src/core/scan"
	"cg-go/src/core/screen"
	"cg-go/src/core/transform"
	"cg-go/src/core/vec"
	"cg-go/src/shapes"
	"cg-go/src/window"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type world struct {
	screen  *screen.Screen
	objects []*shapes.GeometricShape
	player  *shapes.GeometricShape
}

func New() *world {
	return &world{}
}

func (w *world) Create() {
	chanImg := make(chan [][]color.RGBA)

	go func(c chan [][]color.RGBA) {
		var img, err = image.ReadImage("./resources/gato.png")
		if err != nil {
			panic(err.Error())
		}
		c <- img
	}(chanImg)

	w.objects = []*shapes.GeometricShape{
		shapes.NewTriangle(50, 100, vec.NewVec2(50, 250)).
			WithColors([]color.RGBA{
				colors.HexToRGBA(colors.Yellow),
				colors.HexToRGBA(colors.Purple),
				colors.HexToRGBA(colors.Purple),
				colors.HexToRGBA(colors.Yellow),
			}),

		shapes.NewSquare(50, 50, vec.NewVec2(200, 200)),
		shapes.NewSquare(100, 100, vec.NewVec2(50, 100)).WithColors([]color.RGBA{
			colors.HexToRGBA(colors.Yellow),
			colors.HexToRGBA(colors.Purple),
			colors.HexToRGBA(colors.Purple),
			colors.HexToRGBA(colors.Yellow),
		}),
	}

	w.player = shapes.NewSquare(200, 100, vec.NewVec2(300, 300)).WithTextureVertices([]vec.VecTexture{
		vec.NewVecTexture(0, 0),
		vec.NewVecTexture(1, 0),
		vec.NewVecTexture(1, 1),
		vec.NewVecTexture(0, 1),
	}).WithTexture(<-chanImg)

	w.screen = screen.New().
		SetWidth(1000).
		SetHeight(800).
		SetTitle("Term").
		SetOnUpdate(w.Update).
		Build()

	w.MapObjects()

	w.screen.Run()
}

const speed = 2

func (w *world) Update(ctx *ebiten.Image) {
	ctx.Clear()

	velocity := vec.Zeros()

	///window.MapPointsToWindow(w.player, vec.NewVec2(0, 0), vec.NewVec2(1000, 800), w.viewportWidth, w.viewportHeight)
	//time.Sleep(time.Minute)

	for i := range w.objects {
		if len(w.objects[i].Texture) != 0 {
			scan.ScanlineTexture(ctx, w.objects[i], w.objects[i].Texture)
			w.objects[i].DrawMesh(ctx)
		} else if len(w.objects[i].ColorVertices) != 0 {
			scan.ScanlineGradient(ctx, w.objects[i])
			w.objects[i].DrawMesh(ctx)
		} else {
			scan.ScanlineBasic(ctx, w.objects[i], colors.HexToRGBA(colors.Purple))
			w.objects[i].DrawMesh(ctx)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) {
		transform.ScaleVertices(0.95, 0.95, w.player)
		scan.ScanlineTexture(ctx, w.player, w.player.Texture)
		w.player.DrawMesh(ctx)
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		transform.ScaleVertices(1.05, 1.05, w.player)
		scan.ScanlineTexture(ctx, w.player, w.player.Texture)
		w.player.DrawMesh(ctx)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		velocity = vec.NewVec2(-speed, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		velocity = vec.NewVec2(speed, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		velocity = vec.NewVec2(0, -speed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		velocity = vec.NewVec2(0, speed)
	}

	if !velocity.IsZero() {
		transform.TranslateVertices(velocity, w.player)
	}

	scan.ScanlineTexture(ctx, w.player, w.player.Texture)
	w.player.DrawMesh(ctx)
}

func (w *world) MapObjects() {
	window := window.New(vec.NewVec2(0, 0), vec.NewVec2(1000, 800))

	window.MapPoints(w.player, 1000, 800)

	for i := range w.objects {
		window.MapPoints(w.objects[i], 1000, 800)
	}
}
