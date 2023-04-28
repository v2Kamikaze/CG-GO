package core

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type window struct {
	update func(screen *ebiten.Image)
}

func (g *window) Update() error {
	return nil
}

func (g *window) Draw(screen *ebiten.Image) {
	g.update(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
}

func (g *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 300
}

func Run(update func(screen *ebiten.Image)) {
	ebiten.RunGame(&window{update})
}
