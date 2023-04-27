package main

import (
	"cg-go/src/core"
	"cg-go/src/lifecycle"
	"fmt"
	"time"
)

func main() {

	fmt.Println(core.ReadImage("./teste.jpg"))

	time.Sleep(time.Hour)
	core.Run(lifecycle.Update)
}
