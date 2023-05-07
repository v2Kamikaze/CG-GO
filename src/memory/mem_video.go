package memory

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type memVideo struct {
	width  int
	height int
	mem    []color.RGBA
	wg     *sync.WaitGroup
}

func New(width, height int) *memVideo {
	return &memVideo{
		width:  width,
		height: height,
		mem:    make([]color.RGBA, width*height),
		wg:     &sync.WaitGroup{},
	}
}

func (m *memVideo) SetPixel(x, y int, color color.RGBA) {
	// Para aplicar transparência.
	if color.A == 0 {
		return
	}

	if x >= m.width || x < 0 || y >= m.height || y < 0 {
		return
	}

	idx := m.calcPosition(x, y)

	if idx >= len(m.mem) {
		return
	}

	m.mem[idx] = color
}

func (m *memVideo) GetPixel(x, y int) color.RGBA {
	if x >= m.width || x < 0 || y >= m.height || y < 0 {
		return color.RGBA{}
	}

	idx := m.calcPosition(x, y)

	if idx >= len(m.mem) {
		return color.RGBA{}
	}

	return m.mem[idx]
}

func (m *memVideo) Draw(ctx *ebiten.Image) {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			ctx.Set(i, j, m.GetPixel(i, j))
		}
	}

}

func (m *memVideo) calcPosition(x, y int) int {
	return y*m.width + x
}

func (m *memVideo) clearSection(x int) {
	defer m.wg.Done()
	for y := 0; y < m.height; y++ {
		m.mem[m.calcPosition(x, y)] = color.RGBA{}
	}
}

func (m *memVideo) Clear() {
	m.wg.Add(m.width)
	for x := 0; x < m.width; x++ {
		go m.clearSection(x)
	}

	m.wg.Wait()
}

func (m *memVideo) Width() int {
	return m.width
}

func (m *memVideo) Height() int {
	return m.height
}

// 5 2
// 3, 2 => w * x + j
//
//
// [ 0 1 2 3 4 5 6 7 8 9]
// [
//	[0 1 2 3 4]
//  [5 6 7 8 9]
// ]
