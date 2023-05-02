package main

import (
	"cg-go/src/core/window"
	"cg-go/src/render"
)

func main() {
	window.NewWindow().
		SetWidth(1240).
		SetHeight(860).
		SetTitle("Term").
		SetOnUpdate(render.Update).
		Run()
}
