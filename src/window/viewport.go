package window

type Viewport struct {
	width, height float64
}

func NewViewport(width, height float64) *Viewport {
	return &Viewport{width, height}
}
