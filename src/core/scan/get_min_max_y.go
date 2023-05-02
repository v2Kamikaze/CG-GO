package scan

import (
	"cg-go/src/core/vec"
	"math"
)

func GetMinMaxY(vertices []vec.Vec2D) (ymin int, ymax int) {
	ymin = math.MaxUint32
	ymax = 0

	for _, p := range vertices {
		if p.Y < ymin {
			ymin = p.Y
		}

		if p.Y > ymax {
			ymax = p.Y
		}
	}

	return
}
