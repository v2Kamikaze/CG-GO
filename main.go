package main

import (
	"cg-go/src/core"
	"cg-go/src/render"
)

func main() {
	core.Run(render.Update)
}
