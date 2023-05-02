package main

import (
	"cg-go/src/core/window"
	"cg-go/src/render"
)

func main() {
	window.NewWindow().
		SetWidth(540).
		SetHeight(360).
		SetTitle("Term").
		SetOnUpdate(render.Update).
		Run()
}
