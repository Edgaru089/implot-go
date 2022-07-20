#pragma once

#include "../implot/imgui.h"
#include "../implot/implot.h"
#include "Types.h"


// convertions
namespace {

static inline ImVec2 Vec2(const igpVec2 &v) {
	return ImVec2{v.x, v.y};
}
static inline ImVec4 Vec4(const igpVec4 &v) {
	return ImVec4{v.x, v.y, v.z, v.w};
}
static inline ImPlotPoint Point(const igpPoint &p) {
	return ImPlotPoint{p.x, p.y};
}

} // namespace
