package scan

import "math"

func GetMinMaxY(vertices [][]uint32) (ymin int, ymax int) {
	ymin = math.MaxUint32
	ymax = 0

	for _, p := range vertices {
		if int(p[1]) < ymin {
			ymin = int(p[1])
		}
		if int(p[1]) > ymax {
			ymax = int(p[1])
		}
	}

	return
}
