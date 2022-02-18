package implot

// #include <stdlib.h>
// #include <string.h>
// #include "wrapper/Types.h"
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/inkyblackness/imgui-go/v4"
)

func wrapVec2(v imgui.Vec2) C.igpVec2 {
	return C.igpVec2{x: C.float(v.X), y: C.float(v.Y)}
}

func wrapVec4(v imgui.Vec4) C.igpVec4 {
	return C.igpVec4{x: C.float(v.X), y: C.float(v.Y), z: C.float(v.Z), w: C.float(v.W)}
}

func (p Point) wrap() C.igpPoint {
	return C.igpPoint{x: C.double(p.X), y: C.double(p.Y)}
}

func wrapString(str string) *C.char {
	buf := make([]byte, len(str)+1)
	copy(buf, str)
	return (*C.char)(unsafe.Pointer(&buf[0]))
}

func wrapDoubleSlice(slice []float64) *C.double {
	return (*C.double)(unsafe.Pointer(&slice[0]))
}

func wrapPointSlice(slice []Point) (xp, yp *C.double, count, stride C.int) {
	if len(slice) == 0 {
		return
	}
	xp = (*C.double)(unsafe.Pointer(&slice[0].X))
	yp = (*C.double)(unsafe.Pointer(&slice[0].Y))
	count = (C.int)(len(slice))
	stride = (C.int)(unsafe.Sizeof(slice[0]))
	return
}

func wrapXYSlice(xs, ys []float64) (xp, yp *C.double, count, stride C.int) {
	count = (C.int)(minint(len(xs), len(ys)))
	if count == 0 {
		return
	}
	xp = (*C.double)(unsafe.Pointer(&xs[0]))
	yp = (*C.double)(unsafe.Pointer(&ys[0]))
	stride = (C.int)(unsafe.Sizeof(xs[0]))
	return
}

// wraps a []float64 allocating C memory for it.
func wrapDoubleSliceAlloc(slice []float64) (sp *C.double, fin func()) {
	n := len(slice)
	if n == 0 {
		return nil, func() {}
	}
	nbytes := n * int(unsafe.Sizeof(C.double(0)))
	sp = (*C.double)(C.malloc(C.size_t(nbytes)))
	C.memcpy(unsafe.Pointer(sp), unsafe.Pointer(&slice[0]), C.size_t(nbytes))
	return sp, func() {
		C.free(unsafe.Pointer(sp))
	}
}

func wrapStringSlice(slice []string) (sp **C.char, fin func()) {
	if len(slice) == 0 {
		return nil, func() {}
	}

	n := len(slice)
	rsp := make([]uintptr, n+1)
	for i := 0; i < n; i++ {
		nbytes := len(slice[i]) + 1
		rs := C.malloc(C.size_t(nbytes))
		//copy(C.GoBytes(rs, C.int(len(slice[i]))), slice[i])
		bhead := &reflect.SliceHeader{
			Data: uintptr(rs),
			Len:  nbytes + 1,
			Cap:  nbytes + 1,
		}
		b := *((*[]byte)(unsafe.Pointer(bhead)))
		b[copy(b, slice[i])] = 0
		rsp[i] = uintptr(rs)
	}

	// So this fin() keeps the rsp from getting GCed
	return (**C.char)(unsafe.Pointer(&rsp[0])), func() {
		for _, p := range rsp {
			C.free(unsafe.Pointer(p))
		}
	}
}
