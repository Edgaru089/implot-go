
#include "Types.h"
#include "ImPlot.hpp"

namespace {

inline igpVec2 wrapVec2(const ImVec2 &from) { return igpVec2{from.x, from.y}; }
inline igpVec4 wrapVec4(const ImVec4 &from) { return igpVec4{from.x, from.y, from.z, from.w}; }

inline ImVec2 unwrapVec2(const igpVec2 &from) { return ImVec2{from.x, from.y}; }
inline ImVec4 unwrapVec4(const igpVec4 &from) { return ImVec4{from.x, from.y, from.z, from.w}; }

} // namespace
