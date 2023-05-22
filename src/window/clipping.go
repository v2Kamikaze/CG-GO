package window

import (
	"cg-go/src/geo"
)

const (
	top    = 1 << 0
	bottom = 1 << 1
	right  = 1 << 2
	left   = 1 << 3
)

func ClipPolygon(original *geo.GeometricShape, win *Window) *geo.GeometricShape {
	poly := geo.Copy(original)

	return poly
}
