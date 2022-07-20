package implot

// #include "wrapper/Style.h"
import "C"
import (
	"fmt"
	_ "unsafe"

	"github.com/inkyblackness/imgui-go/v4"
)

// Style contains plotting style data.
//
// Right now this struct doesn't have much point, since there's only one
// style (the global one) accessible, and you can't create any.
type PlotStyle struct {
	handle C.igpStyle
}

// CurrentStyle returns a handle to the global Style.
func CurrentPlotStyle() PlotStyle {
	return PlotStyle{handle: C.igpGetStyle()}
}

// Color returns one of the Colors of the style.
func (s PlotStyle) PlotColor(color PlotStyleCol) imgui.Vec4 {
	return unwrapVec4(C.igpStyleGetColor(s.handle, C.igpStyleCol(color)))
}

// SetColor sets one of the Colors of the style.
func (s PlotStyle) SetColor(color PlotStyleCol, to imgui.Vec4) {
	C.igpStyleSetColor(s.handle, C.igpStyleCol(color), wrapVec4(to))
}

// convFloat32 converts a interface to float32/float64, or panics.
func convFloat32(val interface{}, panicmsg interface{}) float32 {
	switch val.(type) {
	case float32:
		return val.(float32)
	case float64:
		return float32(val.(float64))
	}
	panic(panicmsg)
}

// Var returns one of the variables of the style.
// Returns either a float32, a imgui.Vec2, or a int.
func (s PlotStyle) Var(v PlotStyleVar) interface{} {
	switch v {
	case // Float
		PlotStyleVar_LineWeight,
		PlotStyleVar_MarkerSize,
		PlotStyleVar_MarkerWeight,
		PlotStyleVar_FillAlpha,
		PlotStyleVar_ErrorBarSize,
		PlotStyleVar_ErrorBarWeight,
		PlotStyleVar_DigitalBitHeight,
		PlotStyleVar_DigitalBitGap,
		PlotStyleVar_PlotBorderSize,
		PlotStyleVar_MinorAlpha:
		return s.VarFloat(v)
	case // ImVec2
		PlotStyleVar_MajorTickLen,
		PlotStyleVar_MinorTickLen,
		PlotStyleVar_MajorTickSize,
		PlotStyleVar_MinorTickSize,
		PlotStyleVar_MajorGridSize,
		PlotStyleVar_MinorGridSize,
		PlotStyleVar_PlotPadding,
		PlotStyleVar_LabelPadding,
		PlotStyleVar_LegendPadding,
		PlotStyleVar_LegendInnerPadding,
		PlotStyleVar_LegendSpacing,
		PlotStyleVar_MousePosPadding,
		PlotStyleVar_AnnotationPadding,
		PlotStyleVar_FitPadding,
		PlotStyleVar_PlotDefaultSize,
		PlotStyleVar_PlotMinSize:
		return s.VarVec2(v)
	case // Int
		PlotStyleVar_Marker:
		return s.VarInt(v)
	}
	panic(fmt.Errorf("unknown PlotStyleVar %d", v))
}

// SetVar sets one of the variables of the style.
//
// #to must be either a float32, a imgui.Vec2, or a int.
// If the types mismatch, it panics.
func (s PlotStyle) SetVar(v PlotStyleVar, to interface{}) {
	switch v {
	case // Float
		PlotStyleVar_LineWeight,
		PlotStyleVar_MarkerSize,
		PlotStyleVar_MarkerWeight,
		PlotStyleVar_FillAlpha,
		PlotStyleVar_ErrorBarSize,
		PlotStyleVar_ErrorBarWeight,
		PlotStyleVar_DigitalBitHeight,
		PlotStyleVar_DigitalBitGap,
		PlotStyleVar_PlotBorderSize,
		PlotStyleVar_MinorAlpha:
		s.SetVarFloat(v, convFloat32(to, "PlotStyleVar/value type mismatch"))
	case // ImVec2
		PlotStyleVar_MajorTickLen,
		PlotStyleVar_MinorTickLen,
		PlotStyleVar_MajorTickSize,
		PlotStyleVar_MinorTickSize,
		PlotStyleVar_MajorGridSize,
		PlotStyleVar_MinorGridSize,
		PlotStyleVar_PlotPadding,
		PlotStyleVar_LabelPadding,
		PlotStyleVar_LegendPadding,
		PlotStyleVar_LegendInnerPadding,
		PlotStyleVar_LegendSpacing,
		PlotStyleVar_MousePosPadding,
		PlotStyleVar_AnnotationPadding,
		PlotStyleVar_FitPadding,
		PlotStyleVar_PlotDefaultSize,
		PlotStyleVar_PlotMinSize:
		if vt, ok := to.(imgui.Vec2); ok {
			s.SetVarVec2(v, vt)
		} else {
			panic("PlotStyleVar/value type mismatch")
		}
	case // Int
		PlotStyleVar_Marker:
		if vt, ok := to.(int); ok {
			s.SetVarInt(v, vt)
		} else {
			panic("PlotStyleVar/value type mismatch")
		}
	default:
		panic(fmt.Errorf("unknown PlotStyleVar %d", v))
	}
}

// VarFloat returns a float style variable.
// If v is not a float, it panics.
func (s PlotStyle) VarFloat(v PlotStyleVar) float32 {
	return float32(C.igpStyleGetVarFloat(s.handle, C.igpStyleVar(v)))
}

// VarVec2 returns a ImVec2 style variable.
// If v is not a float, it panics.
func (s PlotStyle) VarVec2(v PlotStyleVar) imgui.Vec2 {
	return unwrapVec2(C.igpStyleGetVarVec2(s.handle, C.igpStyleVar(v)))
}

// VarInt returns a int style variable.
// If v is not a float, it panics.
func (s PlotStyle) VarInt(v PlotStyleVar) int {
	return int(C.igpStyleGetVarInt(s.handle, C.igpStyleVar(v)))
}

// SetVarFloat sets a float style variable.
func (s PlotStyle) SetVarFloat(v PlotStyleVar, to float32) {
	C.igpStyleSetVarFloat(s.handle, C.igpStyleVar(v), C.float(to))
}

// SetVarVec2 sets a ImVec2 style variable.
func (s PlotStyle) SetVarVec2(v PlotStyleVar, to imgui.Vec2) {
	C.igpStyleSetVarVec2(s.handle, C.igpStyleVar(v), wrapVec2(to))
}

// SetVarInt sets a int style variable.
func (s PlotStyle) SetVarInt(v PlotStyleVar, to int) {
	C.igpStyleSetVarInt(s.handle, C.igpStyleVar(v), C.int(to))
}

// PlotStyleColorsAuto sets all global style colors to be automatically deduced
// from the current ImGui style.
func PlotStyleColorsAuto() {
	C.igpStyleColorsAuto(nil)
}

// PlotStyleColorsClassic sets the global style colors to mimic
// the ImGui "Classic" style.
func PlotStyleColorsClassic() {
	C.igpStyleColorsClassic(nil)
}

// PlotStyleColorsDark sets the global style colors to mimic
// the ImGui "Dark" style.
func PlotStyleColorsDark() {
	C.igpStyleColorsDark(nil)
}

// PlotStyleColorsLight sets the global style colors to mimic
// the ImGui "Light" style.
func PlotStyleColorsLight() {
	C.igpStyleColorsLight(nil)
}

// PushPlotStyleColor pushes the given color onto the stack.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushPlotStyleColor(id PlotStyleCol, color imgui.Vec4) {
	C.igpPushStyleColor(C.igpStyleCol(id), wrapVec4(color))
}

// PopPlotStyleColor pops one color off the style stack.
// It calls PopStyleColorV(1).
func PopPlotStyleColor() {
	C.igpPopStyleColor(1)
}

// PopPlotStyleColorV pops #count colors off the stack.
func PopPlotStyleColorV(count int) {
	C.igpPopStyleColor(C.int(count))
}

// PushPlotStyleVar pushes a given style variable onto the stack.
//
// If the types of #v and #val mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushPlotStyleVar(v PlotStyleVar, val interface{}) {
	switch v {
	case // Float
		PlotStyleVar_LineWeight,
		PlotStyleVar_MarkerSize,
		PlotStyleVar_MarkerWeight,
		PlotStyleVar_FillAlpha,
		PlotStyleVar_ErrorBarSize,
		PlotStyleVar_ErrorBarWeight,
		PlotStyleVar_DigitalBitHeight,
		PlotStyleVar_DigitalBitGap,
		PlotStyleVar_PlotBorderSize,
		PlotStyleVar_MinorAlpha:
		PushPlotStyleVarFloat(v, convFloat32(val, "StyleVar/value type mismatch"))
	case // ImVec2
		PlotStyleVar_MajorTickLen,
		PlotStyleVar_MinorTickLen,
		PlotStyleVar_MajorTickSize,
		PlotStyleVar_MinorTickSize,
		PlotStyleVar_MajorGridSize,
		PlotStyleVar_MinorGridSize,
		PlotStyleVar_PlotPadding,
		PlotStyleVar_LabelPadding,
		PlotStyleVar_LegendPadding,
		PlotStyleVar_LegendInnerPadding,
		PlotStyleVar_LegendSpacing,
		PlotStyleVar_MousePosPadding,
		PlotStyleVar_AnnotationPadding,
		PlotStyleVar_FitPadding,
		PlotStyleVar_PlotDefaultSize,
		PlotStyleVar_PlotMinSize:
		if valtype, ok := val.(imgui.Vec2); ok {
			PushPlotStyleVarVec2(v, valtype)
		} else {
			panic("StyleVar/value type mismatch")
		}
	case // Int
		PlotStyleVar_Marker:
		if valtype, ok := val.(int); ok {
			PushPlotStyleVarInt(v, valtype)
		} else {
			panic("StyleVar/value type mismatch")
		}
	default:
		panic(fmt.Errorf("unknown StyleVar %d", v))
	}
}

// PushPlotStyleVarFloat pushes a given style variable of float onto the stack.
//
// If the type of #v mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushPlotStyleVarFloat(v PlotStyleVar, val float32) {
	C.igpPushStyleVarFloat(C.igpStyleVar(v), C.float(val))
}

// PushPlotStyleVarVec2 pushes a given style variable of ImVec2 onto the stack.
//
// If the type of #v mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushPlotStyleVarVec2(v PlotStyleVar, val imgui.Vec2) {
	C.igpPushStyleVarVec2(C.igpStyleVar(v), wrapVec2(val))
}

// PushPlotStyleVarInt pushes a given style variable of int onto the stack.
//
// If the type of #v mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushPlotStyleVarInt(v PlotStyleVar, val int) {
	C.igpPushStyleVarInt(C.igpStyleVar(v), C.int(val))
}

// PopPlotStyleVar pops one variable off the stack.
// It calls PopPlotStyleVarV(1).
func PopPlotStyleVar() {
	C.igpPopStyleVar(1)
}

// PopPlotStyleVarV pops #count variables off the stack.
func PopPlotStyleVarV(count int) {
	C.igpPopStyleVar(C.int(count))
}

// SetNextLineStyle set the line color and weight for the next item only.
//
// All the values can be set to Auto/AutoColor to deduce from
// the current Style and Colormap data.
func SetNextLineStyle(color imgui.Vec4, weight float32) {
	C.igpSetNextLineStyle(wrapVec4(color), C.float(weight))
}

// SetNextLineStyle set the the fill color for the next item only.
//
// All the values can be set to Auto/AutoColor to deduce from
// the current Style and Colormap data.
func SetNextFillStyle(color imgui.Vec4, alpha float32) {
	C.igpSetNextFillStyle(wrapVec4(color), C.float(alpha))
}

// SetNextMarkerStyle set the marker style for the next item only.
//
// All the values can be set to Auto/AutoColor to deduce from
// the current Style and Colormap data.
func SetNextMarkerStyle(marker Marker, size float32, fillColor imgui.Vec4, outlineWeight float32, outlineColor imgui.Vec4) {
	C.igpSetNextMarkerStyle(C.igpMarker(marker), C.float(size), wrapVec4(fillColor), C.float(outlineWeight), wrapVec4(outlineColor))
}

// SetNextErrorBarStyle set the error bar style for the next item only.
//
// All the values can be set to Auto/AutoColor to deduce from
// the current Style and Colormap data.
func SetNextErrorBarStyle(color imgui.Vec4, size, weight float32) {
	C.igpSetNextErrorBarStyle(wrapVec4(color), C.float(size), C.float(weight))
}

// GetLastItemColor returns the primary color of the last item (i.e. its legend icon color)
func GetLastItemColor() imgui.Vec4 {
	return unwrapVec4((C.igpGetLastItemColor)())
}

var styleColNames = [PlotStyleCol_Count]string{
	"Line",
	"Fill",
	"MarkerOutline",
	"MarkerFill",
	"ErrorBar",
	"FrameBg",
	"PlotBg",
	"PlotBorder",
	"LegendBg",
	"LegendBorder",
	"LegendText",
	"TitleText",
	"InlayText",
	"AxisText",
	"AxisGrid",
	"AxisTick",
	"AxisBg",
	"AxisBgHovered",
	"AxisBgActive",
	"Selection",
	"Crosshairs",
}

// GetStyleColorName returns the name of a style color.
func GetPlotStyleColorName(id PlotStyleCol) string {
	return styleColNames[id]
}

var markerNames = [Marker_Count]string{
	Marker_Circle:   "Circle",
	Marker_Square:   "Square",
	Marker_Diamond:  "Diamond",
	Marker_Up:       "Up",
	Marker_Down:     "Down",
	Marker_Left:     "Left",
	Marker_Right:    "Right",
	Marker_Cross:    "Cross",
	Marker_Plus:     "Plus",
	Marker_Asterisk: "Asterisk",
}

// GetMarkerName returns the name of a marker.
func GetMarkerName(id Marker) string {
	if id == Marker_None {
		return "None"
	}
	return markerNames[id]
}
