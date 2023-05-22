package main

import (
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/scan"
	"cg-go/src/screen"
	"cg-go/src/tex"
	"cg-go/src/transform"
	"cg-go/src/vec"
	"cg-go/src/window"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const Width, Height = 600, 400
const NumMeteors = 50
const MeteorsMinDist = 30
const WindowFactor = 2
const MeteorSize = 20

const WindowVelocity = 5

var win = window.New(vec.NewVec2D(0, 0), vec.NewVec2D(Width*WindowFactor, Height*WindowFactor))
var mem = memory.New(Width, Height)
var center = win.Center()
var mainViewport = window.NewViewport(vec.Zeros(), vec.NewVec2D(Width, Height))
var miniMap = window.NewViewport(vec.NewVec2D((Width-50), 0), vec.NewVec2D(Width-1, 50))

var meteorTex = tex.ReadImage("./resources/meteor.png")
var meteors = GenerateMeteors()

var blackHoleTex = tex.ReadImage("./resources/Black_hole.png")
var blackHole = geo.NewRect(100, 100, vec.NewVec2D(Width, Height)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var terranTex = tex.ReadImage("./resources/Terran.png")
var terran = geo.NewRect(60, 60, vec.NewVec2D(150, 80)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var lavaTex = tex.ReadImage("./resources/Lava.png")
var lava = geo.NewRect(60, 60, vec.NewVec2D(250, 400)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var barenTex = tex.ReadImage("./resources/Baren.png")
var baren = geo.NewRect(20, 20, lava.Center().ScalarSum(90)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var iceTex = tex.ReadImage("./resources/Ice.png")
var ice = geo.NewRect(45, 45, vec.NewVec2D(900, 320)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var starTex = tex.ReadImage("./resources/star.png")
var stars = GenerateStars()

var gopherTex = tex.ReadImage("./resources/gopher-astronaut.png")
var gopher = geo.NewRect(20, 20, center).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2D(0, 0),
		vec.NewVec2D(1, 0),
		vec.NewVec2D(1, 1),
		vec.NewVec2D(0, 1),
	})

func Update(ctx *ebiten.Image) {

	mem.Clear(colors.ColorBlack)

	DrawStars(mem)
	DrawMeteors(mem)
	MapObjectsToVP(mainViewport)
	MapObjectsToVP(miniMap)

	miniMap.DrawBounds(mem)

	mem.Draw(ctx)

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		win.Zoom(0.95)
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) {
		win.Zoom(1.05)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		win.Translate(vec.NewVec2D(-WindowVelocity, 0))
		transform.TranslateVertices(win.Center().Sub(gopher.Center()), gopher)
		transform.RotateVerticesOnPivot(-4, gopher.Center(), gopher)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		win.Translate(vec.NewVec2D(WindowVelocity, 0))
		transform.TranslateVertices(win.Center().Sub(gopher.Center()), gopher)
		transform.RotateVerticesOnPivot(4, gopher.Center(), gopher)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		win.Translate(vec.NewVec2D(0, -WindowVelocity))
		transform.TranslateVertices(win.Center().Sub(gopher.Center()), gopher)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		win.Translate(vec.NewVec2D(0, WindowVelocity))
		transform.TranslateVertices(win.Center().Sub(gopher.Center()), gopher)
	}

	RotateMeteors()
	transform.RotateVerticesOnPivot(10, blackHole.Center(), blackHole)
	transform.RotateVerticesOnPivot(0.1, terran.Center(), terran)
	transform.RotateVerticesOnPivot(-0.5, lava.Center(), lava)
	transform.RotateVerticesOnPivot(0.5, ice.Center(), ice)
	transform.RotateVerticesOnPivot(2, lava.Center(), baren)

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

func MapObjectsToVP(vp *window.Viewport) {

	blackHole.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, blackHoleTex)
	})

	terran.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, terranTex)
	})

	lava.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, lavaTex)
	})

	ice.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, iceTex)
	})

	baren.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, barenTex)
	})

	gopher.Apply(func(s *geo.GeometricShape) {
		win.MapPoints(s, vp)
		scan.ScanlineTexture(mem, s, gopherTex)
	})

}

func GenerateStars() (stars []*geo.GeometricShape) {
	for i := 0.0; i < Width*WindowFactor; i += 20 {
		star := geo.NewRect(5, 5, vec.NewVec2D(i, float64(rand.Intn(Height*WindowFactor)))).
			WithTextureVertices([]vec.Vec2D{
				vec.Zeros(),
				vec.NewVec2D(1, 0),
				vec.Ones(),
				vec.NewVec2D(0, 1),
			})
		stars = append(stars, star)
	}
	return
}

func DrawStars(mem memory.Memory) {
	for _, star := range stars {
		star.Apply(func(s *geo.GeometricShape) {
			win.MapPoints(s, mainViewport)
			scan.ScanlineTexture(mem, s, starTex)
		})
	}
}

func MeteorsPositions() []vec.Vec2D {

	points := make([]vec.Vec2D, NumMeteors)

	for i := 0; i < NumMeteors; i++ {
		var x, y float64
		isValid := false

		for !isValid {
			isValid = true
			x = rand.Float64() * Width * WindowFactor
			y = rand.Float64() * Height * WindowFactor

			for j := 0; j < i; j++ {
				if vec.Distance(points[j], vec.NewVec2D(x, y)) < MeteorsMinDist {
					isValid = false
					break
				}
			}
		}

		points[i] = vec.NewVec2D(x, y)
	}

	return points
}

func GenerateMeteors() (meteors []*geo.GeometricShape) {
	for _, pos := range MeteorsPositions() {
		meteors = append(meteors, geo.NewRect(MeteorSize, MeteorSize, pos).WithTextureVertices([]vec.Vec2D{
			vec.Zeros(),
			vec.NewVec2D(1, 0),
			vec.Ones(),
			vec.NewVec2D(0, 1),
		}))
	}
	return
}

func DrawMeteors(mem memory.Memory) {
	for _, meteor := range meteors {
		meteor.Apply(func(s *geo.GeometricShape) {
			win.MapPoints(s, mainViewport)
			scan.ScanlineTexture(mem, s, meteorTex)
		})

		meteor.Apply(func(s *geo.GeometricShape) {
			win.MapPoints(s, miniMap)
			scan.ScanlineTexture(mem, s, meteorTex)
		})
	}
}

func RotateMeteors() {
	for i := range meteors {
		if i%2 == 0 {
			transform.RotateVerticesOnPivot(-1, meteors[i].Center(), meteors[i])
		} else {
			transform.RotateVerticesOnPivot(1, meteors[i].Center(), meteors[i])
		}

	}
}
