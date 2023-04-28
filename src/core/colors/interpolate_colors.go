package colors

import "image/color"

func InterpolateColors(startColor color.RGBA, endColor color.RGBA, t float64) color.RGBA {
	r := uint8(float64(startColor.R) + float64(endColor.R-startColor.R)*t)
	g := uint8(float64(startColor.G) + float64(endColor.G-startColor.G)*t)
	b := uint8(float64(startColor.B) + float64(endColor.B-startColor.B)*t)
	a := uint8(float64(startColor.A) + float64(endColor.A-startColor.A)*t)
	return color.RGBA{r, g, b, a}
}
