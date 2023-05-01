package pixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawPixel(screen *ebiten.Image, x, y int, color color.RGBA) {

	if x > screen.Bounds().Max.X {
		x = screen.Bounds().Max.X
	}

	if x < 0 {
		x = 0
	}

	if y > screen.Bounds().Max.Y {
		y = screen.Bounds().Max.Y
	}

	if y < 0 {
		y = 0
	}

	screen.Set(x, y, color)
}
