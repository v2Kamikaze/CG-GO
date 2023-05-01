package main

import (
	"cg-go/src/core/window"
	"cg-go/src/render"
)

func main() {
	window.NewWindow().
		SetHeight(400).
		SetWidth(300).
		SetTitle("Term").
		SetOnUpdate(render.Update).
		Run()
}
