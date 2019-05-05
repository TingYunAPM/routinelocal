package routinelocal

import (
	"runtime"
)

type Storage interface {
	Get() uintptr
	Set(p uintptr)
}

type byNative struct {
	Storage
}

func (g *byNative) Get() uintptr {
	return runtime.GetRoutineLocal()
}
func (g *byNative) Set(p uintptr) {
	runtime.SetRoutineLocal()
}

var _inst = byNative{}

func Get() Storage {
	return &_inst
}
