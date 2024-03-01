package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type smallStruct struct {
	a, b int64
	c, d float64
}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	debug.SetMaxThreads(1)
	smallAllocation()
}
