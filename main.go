package main

import (
	"cg-go/src/core/screen"
	"cg-go/src/world"
)

func main() {
	screen.New().
		SetWidth(1000).
		SetHeight(800).
		SetTitle("Term").
		SetOnUpdate(world.Update).
		Build().
		Run()
}
