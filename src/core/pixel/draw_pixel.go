package pixel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawPixel(screen *ebiten.Image, x, y int, color color.RGBA) {

	if x > screen.Bounds().Max.X {
		x = screen.Bounds().Max.X
	}

	if x < screen.Bounds().Min.X {
		x = screen.Bounds().Min.X
	}

	if y > screen.Bounds().Max.Y {
		y = screen.Bounds().Max.Y
	}

	if y < screen.Bounds().Min.Y {
		y = screen.Bounds().Min.Y
	}

	screen.Set(x, y, color)
}
