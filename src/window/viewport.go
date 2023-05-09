package window

import (
	"cg-go/src/core/vec"
	"cg-go/src/memory"
	"cg-go/src/pixel"
	"image/color"
)

type Viewport struct {
	pi, pf vec.Vec2D
}

func NewViewport(pi, pf vec.Vec2D) *Viewport {
	return &Viewport{pi, pf}
}

func (vp *Viewport) DrawBounds(mem memory.Memory) {
	pixel.DrawLine(mem, vp.pi, vec.NewVec2(vp.pf.X, vp.pi.Y), color.RGBA{255, 255, 255, 255})
	pixel.DrawLine(mem, vec.NewVec2(vp.pf.X, vp.pi.Y), vp.pf, color.RGBA{255, 255, 255, 255})
	pixel.DrawLine(mem, vp.pf, vec.NewVec2(vp.pi.X, vp.pf.Y), color.RGBA{255, 255, 255, 255})
	pixel.DrawLine(mem, vec.NewVec2(vp.pi.X, vp.pf.Y), vp.pi, color.RGBA{255, 255, 255, 255})
}
