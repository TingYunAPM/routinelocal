package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/TingYunAPM/routinelocal"
)

var _interface routinelocal.Storage

func test() {
	x := _interface.Get()
	p := unsafe.Pointer(x)
	fmt.Println("Set:", p)
	fmt.Println(*(*uint)(p))
}
func main() {
	_interface = routinelocal.Get()
	go func() {
		var data uint = 1234
		p := unsafe.Pointer(&data)
		fmt.Println("Set:", p)
		_interface.Set(uintptr(p))
		test()
	}()
	time.Sleep(1000 * time.Microsecond)
}
