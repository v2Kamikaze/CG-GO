package main

import (
	"cg-go/src/core"
	"cg-go/src/lifecycle"
)

func main() {
	core.Run(lifecycle.Update)
}
