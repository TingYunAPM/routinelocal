package routinelocal

import (
	"runtime"
	"unsafe"
)

type Storage interface {
	Get() unsafe.Pointer
	Set(p unsafe.Pointer)
}

type byNative struct {
	Storage
}

func (g *byNative) Get() unsafe.Pointer {
	return runtime.GetRoutineLocal()
}
func (g *byNative) Set(p unsafe.Pointer) {
	runtime.SetRoutineLocal(p)
}

var _inst = byNative{}

func Get() Storage {
	return &_inst
}
