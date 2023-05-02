package window

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type window struct {
	update        func(screen *ebiten.Image)
	width, height int
}

func (w *window) Update() error {
	return nil
}

func (w *window) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
	w.update(screen)
}

func (w *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.width, w.height
}

func (w *window) SetWidth(width int) *window {
	w.width = width
	return w
}

func (w *window) SetHeight(height int) *window {
	w.height = height
	return w
}

func (w *window) SetOnUpdate(update func(screen *ebiten.Image)) *window {
	w.update = update
	return w
}

func (w *window) SetTitle(title string) *window {
	ebiten.SetWindowTitle(title)
	return w
}

func NewWindow() *window {
	return &window{}
}

func (w *window) Run() {
	if err := ebiten.RunGameWithOptions(w, &ebiten.RunGameOptions{}); err != nil {
		log.Fatal(err)
	}
}
