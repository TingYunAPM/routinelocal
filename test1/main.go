package main

import (
	"fmt"
	"time"
	"unsafe"
	"sync/atomic"
	"github.com/TingYunAPM/routinelocal"
)

var _interface routinelocal.Storage

func test() {
	x := _interface.Get()
	p := (*MyType)(x)
	if p.value != 1234 {
		fmt.Println("wrong value", p.value)
	}
}
type MyType struct {
	value  int
}
func main() {
	var count int32 = 0
	var runtimes int32 = 0
	_interface = routinelocal.Get()
	begin := time.Now()
	for i := 0; i < 1000000; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
			atomic.AddInt32(&runtimes,1)
			value := &MyType{}
			value.value = 1234
			_interface.Set(unsafe.Pointer(value))
			test()
			atomic.AddInt32(&count, -1)
			_interface.Set(nil)
		}()
	}
	for count > 0 && runtimes < 1000000 {
		time.Sleep(1000 * time.Microsecond)
	}
	end := time.Now()
	fmt.Println("use ", end.Sub(begin))
}
