package routinelocal
import (
	"unsafe"
)
type Storage interface {
	Get() unsafe.Pointer
	Set(p unsafe.Pointer)
}
