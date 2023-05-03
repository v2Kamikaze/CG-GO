package window

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Window struct {
	update        func(screen *ebiten.Image)
	title         string
	width, height int
}

func (w *Window) Update() error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
	w.update(screen)
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.width, w.height
}

func (w *Window) SetWidth(width int) *Window {
	w.width = width
	return w
}

func (w *Window) SetHeight(height int) *Window {
	w.height = height
	return w
}

func (w *Window) SetOnUpdate(update func(screen *ebiten.Image)) *Window {
	w.update = update
	return w
}

func (w *Window) SetTitle(title string) *Window {
	w.title = title
	return w
}

func (w *Window) Build() *Window {
	ebiten.SetWindowTitle(w.title)
	ebiten.SetWindowSize(w.width, w.height)
	return w
}

func NewWindow() *Window {
	return &Window{}
}

func (w *Window) Run() {
	if err := ebiten.RunGameWithOptions(w, &ebiten.RunGameOptions{}); err != nil {
		log.Fatal(err)
	}
}
