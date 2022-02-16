package implot

// #include "wrapper/Types.h"
import "C"
import "github.com/inkyblackness/imgui-go/v4"

func wrapVec2(v imgui.Vec2) C.igpVec2 {
	return C.igpVec2{x: C.float(v.X), y: C.float(v.Y)}
}

func wrapVec4(v imgui.Vec4) C.igpVec4 {
	return C.igpVec4{x: C.float(v.X), y: C.float(v.Y), z: C.float(v.Z), w: C.float(v.W)}
}

func (p Point) wrap() C.igpPoint {
	return C.igpPoint{x: C.double(p.X), y: C.double(p.Y)}
}
