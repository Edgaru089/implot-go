package implot

// #include "wrapper/Types.h"
import "C"
import (
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
