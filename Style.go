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
type Style struct {
	handle C.igpStyle
}

// CurrentStyle returns a handle to the global Style.
func CurrentStyle() Style {
	return Style{handle: C.igpGetStyle()}
}

// Color returns one of the Colors of the style.
func (s Style) Color(color StyleCol) imgui.Vec4 {
	return unwrapVec4(C.igpStyleGetColor(s.handle, C.igpStyleCol(color)))
}

// SetColor sets one of the Colors of the style.
func (s Style) SetColor(color StyleCol, to imgui.Vec4) {
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
func (s Style) Var(v StyleVar) interface{} {
	switch v {
	case // Float
		StyleVar_LineWeight,
		StyleVar_MarkerSize,
		StyleVar_MarkerWeight,
		StyleVar_FillAlpha,
		StyleVar_ErrorBarSize,
		StyleVar_ErrorBarWeight,
		StyleVar_DigitalBitHeight,
		StyleVar_DigitalBitGap,
		StyleVar_PlotBorderSize,
		StyleVar_MinorAlpha:
		return s.VarFloat(v)
	case // ImVec2
		StyleVar_MajorTickLen,
		StyleVar_MinorTickLen,
		StyleVar_MajorTickSize,
		StyleVar_MinorTickSize,
		StyleVar_MajorGridSize,
		StyleVar_MinorGridSize,
		StyleVar_PlotPadding,
		StyleVar_LabelPadding,
		StyleVar_LegendPadding,
		StyleVar_LegendInnerPadding,
		StyleVar_LegendSpacing,
		StyleVar_MousePosPadding,
		StyleVar_AnnotationPadding,
		StyleVar_FitPadding,
		StyleVar_PlotDefaultSize,
		StyleVar_PlotMinSize:
		return s.VarVec2(v)
	case // Int
		StyleVar_Marker:
		return s.VarInt(v)
	}
	panic(fmt.Errorf("unknown StyleVar %d", v))
}

// SetVar sets one of the variables of the style.
//
// #to must be either a float32, a imgui.Vec2, or a int.
// If the types mismatch, it panics.
func (s Style) SetVar(v StyleVar, to interface{}) {
	switch v {
	case // Float
		StyleVar_LineWeight,
		StyleVar_MarkerSize,
		StyleVar_MarkerWeight,
		StyleVar_FillAlpha,
		StyleVar_ErrorBarSize,
		StyleVar_ErrorBarWeight,
		StyleVar_DigitalBitHeight,
		StyleVar_DigitalBitGap,
		StyleVar_PlotBorderSize,
		StyleVar_MinorAlpha:
		s.SetVarFloat(v, convFloat32(to, "StyleVar/value type mismatch"))
	case // ImVec2
		StyleVar_MajorTickLen,
		StyleVar_MinorTickLen,
		StyleVar_MajorTickSize,
		StyleVar_MinorTickSize,
		StyleVar_MajorGridSize,
		StyleVar_MinorGridSize,
		StyleVar_PlotPadding,
		StyleVar_LabelPadding,
		StyleVar_LegendPadding,
		StyleVar_LegendInnerPadding,
		StyleVar_LegendSpacing,
		StyleVar_MousePosPadding,
		StyleVar_AnnotationPadding,
		StyleVar_FitPadding,
		StyleVar_PlotDefaultSize,
		StyleVar_PlotMinSize:
		if vt, ok := to.(imgui.Vec2); ok {
			s.SetVarVec2(v, vt)
		} else {
			panic("StyleVar/value type mismatch")
		}
	case // Int
		StyleVar_Marker:
		if vt, ok := to.(int); ok {
			s.SetVarInt(v, vt)
		} else {
			panic("StyleVar/value type mismatch")
		}
	default:
		panic(fmt.Errorf("unknown StyleVar %d", v))
	}
}

// VarFloat returns a float style variable.
// If v is not a float, it panics.
func (s Style) VarFloat(v StyleVar) float32 {
	return float32(C.igpStyleGetVarFloat(s.handle, C.igpStyleVar(v)))
}

// VarVec2 returns a ImVec2 style variable.
// If v is not a float, it panics.
func (s Style) VarVec2(v StyleVar) imgui.Vec2 {
	return unwrapVec2(C.igpStyleGetVarVec2(s.handle, C.igpStyleVar(v)))
}

// VarInt returns a int style variable.
// If v is not a float, it panics.
func (s Style) VarInt(v StyleVar) int {
	return int(C.igpStyleGetVarInt(s.handle, C.igpStyleVar(v)))
}

// SetVarFloat sets a float style variable.
func (s Style) SetVarFloat(v StyleVar, to float32) {
	C.igpStyleSetVarFloat(s.handle, C.igpStyleVar(v), C.float(to))
}

// SetVarVec2 sets a ImVec2 style variable.
func (s Style) SetVarVec2(v StyleVar, to imgui.Vec2) {
	C.igpStyleSetVarVec2(s.handle, C.igpStyleVar(v), wrapVec2(to))
}

// SetVarInt sets a int style variable.
func (s Style) SetVarInt(v StyleVar, to int) {
	C.igpStyleSetVarInt(s.handle, C.igpStyleVar(v), C.int(to))
}

// StyleColorsAuto sets all global style colors to be automatically deduced
// from the current ImGui style.
func StyleColorsAuto() {
	C.igpStyleColorsAuto(nil)
}

// StyleColorsClassic sets the global style colors to mimic
// the ImGui "Classic" style.
func StyleColorsClassic() {
	C.igpStyleColorsClassic(nil)
}

// StyleColorsDark sets the global style colors to mimic
// the ImGui "Dark" style.
func StyleColorsDark() {
	C.igpStyleColorsDark(nil)
}

// StyleColorsLight sets the global style colors to mimic
// the ImGui "Light" style.
func StyleColorsLight() {
	C.igpStyleColorsLight(nil)
}

// PushStyleColor pushes the given color onto the stack.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushStyleColor(id StyleCol, color imgui.Vec4) {
	C.igpPushStyleColor(C.igpStyleCol(id), wrapVec4(color))
}

// PopStyleColor pops one color off the style stack.
// It calls PopStyleColorV(1).
func PopStyleColor() {
	C.igpPopStyleColor(1)
}

// PopStyleColorV pops #count colors off the stack.
func PopStyleColorV(count int) {
	C.igpPopStyleColor(C.int(count))
}

// PushStyleVar pushes a given style variable onto the stack.
//
// If the types of #v and #val mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushStyleVar(v StyleVar, val interface{}) {
	switch v {
	case // Float
		StyleVar_LineWeight,
		StyleVar_MarkerSize,
		StyleVar_MarkerWeight,
		StyleVar_FillAlpha,
		StyleVar_ErrorBarSize,
		StyleVar_ErrorBarWeight,
		StyleVar_DigitalBitHeight,
		StyleVar_DigitalBitGap,
		StyleVar_PlotBorderSize,
		StyleVar_MinorAlpha:
		PushStyleVarFloat(v, convFloat32(val, "StyleVar/value type mismatch"))
	case // ImVec2
		StyleVar_MajorTickLen,
		StyleVar_MinorTickLen,
		StyleVar_MajorTickSize,
		StyleVar_MinorTickSize,
		StyleVar_MajorGridSize,
		StyleVar_MinorGridSize,
		StyleVar_PlotPadding,
		StyleVar_LabelPadding,
		StyleVar_LegendPadding,
		StyleVar_LegendInnerPadding,
		StyleVar_LegendSpacing,
		StyleVar_MousePosPadding,
		StyleVar_AnnotationPadding,
		StyleVar_FitPadding,
		StyleVar_PlotDefaultSize,
		StyleVar_PlotMinSize:
		if valtype, ok := val.(imgui.Vec2); ok {
			PushStyleVarVec2(v, valtype)
		} else {
			panic("StyleVar/value type mismatch")
		}
	case // Int
		StyleVar_Marker:
		if valtype, ok := val.(int); ok {
			PushStyleVarInt(v, valtype)
		} else {
			panic("StyleVar/value type mismatch")
		}
	default:
		panic(fmt.Errorf("unknown StyleVar %d", v))
	}
}

// PushStyleVarFloat pushes a given style variable of float onto the stack.
//
// If the type of #v mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushStyleVarFloat(v StyleVar, val float32) {
	C.igpPushStyleVarFloat(C.igpStyleVar(v), C.float(val))
}

// PushStyleVarVec2 pushes a given style variable of ImVec2 onto the stack.
//
// If the type of #v mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushStyleVarVec2(v StyleVar, val imgui.Vec2) {
	C.igpPushStyleVarVec2(C.igpStyleVar(v), wrapVec2(val))
}

// PushStyleVarInt pushes a given style variable of int onto the stack.
//
// If the type of #v mismatch, it panics.
//
// You MUST call a pop for every push, otherwise you will leak memory!
// This behaves just like ImGui itself.
func PushStyleVarInt(v StyleVar, val int) {
	C.igpPushStyleVarInt(C.igpStyleVar(v), C.int(val))
}

// PopStyleVar pops one variable off the stack.
// It calls PopStyleVarV(1).
func PopStyleVar() {
	C.igpPopStyleVar(1)
}

// PopStyleVarV pops #count variables off the stack.
func PopStyleVarV(count int) {
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

var styleColNames = [StyleCol_Count]string{
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
func GetStyleColorName(id StyleCol) string {
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
