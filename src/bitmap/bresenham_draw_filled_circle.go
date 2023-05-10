package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

func BresenhamDrawFilledCircle(mem memory.Memory, center vec.Vec2D, radius float64, c color.RGBA) {
	x := int(radius)
	y := 0
	xChange := 1 - 2*x
	yChange := 1
	radiusError := 0

	for x >= y {
		for i := int(center.X) - x; i <= int(center.X)+x; i++ {
			mem.SetPixel(i, int(center.Y)+y, c)
			mem.SetPixel(i, int(center.Y)-y, c)
		}
		for i := int(center.X) - y; i <= int(center.X)+y; i++ {
			mem.SetPixel(i, int(center.Y)+x, c)
			mem.SetPixel(i, int(center.Y)-x, c)
		}

		y++
		radiusError += yChange
		yChange += 2

		if 2*radiusError+xChange > 0 {
			x--
			radiusError += xChange
			xChange += 2
		}
	}
}
