package routinelocal

import (
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type byGidMap struct {
	Storage
	_id_map map[int64]uintptr
	lock    *sync.RWMutex
}

func get_gid() int64 {
	buffer := make([]byte, 65536)
	runtime.Stack(buffer, false)
	if parts := Split(string(buffer[:32]), " "); parts[0] == "goroutine" && len(parts) > 1 {
		if goid, err := strconv.ParseInt(parts[1], 10, 0); err == nil {
			return goid
		}
	}
	return -1
}
func (g *byGidMap) init() *byGidMap {
	g._id_map = map[int64]uintptr{}
	g.lock = new(sync.RWMutex)
	return g
}
func (g *byGidMap) Get() uintptr {
	gid := get_gid()
	if gid == -1 {
		return uintptr(0)
	}
	g.lock.RLock()
	result := uintptr(0)
	if r, found := g._id_map[gid]; found {
		result = r
	}
	g.lock.RUnlock()
	return result
}
func (g *byGidMap) Set(p uintptr) {
	gid := get_gid()
	if gid == -1 {
		return
	}
	g.lock.Lock()
	if p != uintptr(0) {
		g._id_map[gid] = p
	} else if _, exist := g._id_map[gid]; exist {
		delete(g._id_map, gid)
	}
	g.lock.Unlock()
}

var _gidmap *byGidMap

func Get() Storage {
	return _gidmap
}

func init() {
	_gidmap = (&byGidMap{}).init()
}

func Split(s, sep string) []string {
	sep_len := len(sep)
	if sep_len == 0 {
		return []string{s}
	}
	count := 0
	index := 0
	for i := strings.Index(s[index:], sep); index < len(s); i = strings.Index(s[index:], sep) {
		if i != 0 {
			count++
		}
		if i < 0 {
			break
		}
		index += i + sep_len
	}
	if count == 0 {
		return []string{}
	}
	r := make([]string, count)
	index = 0
	for i := 0; i < count; i++ {
		sep_index := 0
		for sep_index = strings.Index(s[index:], sep); sep_index == 0; sep_index = strings.Index(s[index:], sep) {
			index += sep_len
		}
		if sep_index > 0 {
			r[i] = string(s[index : index+sep_index])
		} else {
			r[i] = string(s[index:])
		}
		index += sep_index
	}
	return r
}
