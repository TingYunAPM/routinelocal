package routinelocal

type Storage interface {
	Get() uintptr
	Set(p uintptr)
}
