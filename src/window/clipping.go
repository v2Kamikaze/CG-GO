package window

import (
	"cg-go/src/geo"
)

func ClipPolygon(original *geo.GeometricShape, vp *Viewport, win *Window) *geo.GeometricShape {
	poly := geo.Copy(original)

	return poly
}
