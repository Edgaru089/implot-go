
#include "Style.h"
#include "ImPlot.hpp"
#include "Panic.h"
#include "Wraps.hpp"


namespace {
inline ImPlotStyle &unwrapStyle(igpStyle s) {
	return *reinterpret_cast<ImPlotStyle *>(s);
}

inline float &getStylef(igpStyle style, igpStyleVar id) {
	ImPlotStyle &s = unwrapStyle(style);
	switch (id) {
		case ImPlotStyleVar_LineWeight: return s.LineWeight;
		case ImPlotStyleVar_MarkerSize: return s.MarkerSize;
		case ImPlotStyleVar_MarkerWeight: return s.MarkerWeight;
		case ImPlotStyleVar_FillAlpha: return s.FillAlpha;
		case ImPlotStyleVar_ErrorBarSize: return s.ErrorBarSize;
		case ImPlotStyleVar_ErrorBarWeight: return s.ErrorBarWeight;
		case ImPlotStyleVar_DigitalBitHeight: return s.DigitalBitHeight;
		case ImPlotStyleVar_DigitalBitGap: return s.DigitalBitGap;
		case ImPlotStyleVar_PlotBorderSize: return s.PlotBorderSize;
		case ImPlotStyleVar_MinorAlpha: return s.MinorAlpha;
	}
	igpPanic("StyleVar/value type mismatch");
}
inline ImVec2 &getStylev(igpStyle style, igpStyleVar id) {
	ImPlotStyle &s = unwrapStyle(style);
	switch (id) {
		case ImPlotStyleVar_MajorTickLen: return s.MajorTickLen;
		case ImPlotStyleVar_MinorTickLen: return s.MinorTickLen;
		case ImPlotStyleVar_MajorTickSize: return s.MajorTickSize;
		case ImPlotStyleVar_MinorTickSize: return s.MinorTickSize;
		case ImPlotStyleVar_MajorGridSize: return s.MajorGridSize;
		case ImPlotStyleVar_MinorGridSize: return s.MinorGridSize;
		case ImPlotStyleVar_PlotPadding: return s.PlotPadding;
		case ImPlotStyleVar_LabelPadding: return s.LabelPadding;
		case ImPlotStyleVar_LegendPadding: return s.LegendPadding;
		case ImPlotStyleVar_LegendInnerPadding: return s.LegendInnerPadding;
		case ImPlotStyleVar_LegendSpacing: return s.LegendSpacing;
		case ImPlotStyleVar_MousePosPadding: return s.MousePosPadding;
		case ImPlotStyleVar_AnnotationPadding: return s.AnnotationPadding;
		case ImPlotStyleVar_FitPadding: return s.FitPadding;
		case ImPlotStyleVar_PlotDefaultSize: return s.PlotDefaultSize;
		case ImPlotStyleVar_PlotMinSize: return s.PlotMinSize;
	}
	igpPanic("StyleVar/value type mismatch");
}

inline int &getStylei(igpStyle style, igpStyleVar id) {
	ImPlotStyle &s = unwrapStyle(style);
	switch (id) {
		case ImPlotStyleVar_Marker: return s.Marker;
	}
	igpPanic("StyleVar/value type mismatch");
}

} // namespace


// Get the global style
igpStyle igpGetStyle() {
	return reinterpret_cast<igpStyle>(&ImPlot::GetStyle());
}

// Access style variables
igpVec4 igpStyleGetColor(igpStyle style, igpStyleCol id) {
	return wrapVec4(unwrapStyle(style).Colors[id]);
}
void igpStyleSetColor(igpStyle style, igpStyleCol id, igpVec4 color) {
	unwrapStyle(style).Colors[id] = unwrapVec4(color);
}

float   igpStyleGetVarFloat(igpStyle style, igpStyleVar id) { return getStylef(style, id); }
int     igpStyleGetVarInt(igpStyle style, igpStyleVar id) { return getStylei(style, id); }
igpVec2 igpStyleGetVarVec2(igpStyle style, igpStyleVar id) { return wrapVec2(getStylev(style, id)); }
void    igpStyleSetVarFloat(igpStyle style, igpStyleVar id, float var) { getStylef(style, id) = var; }
void    igpStyleSetVarInt(igpStyle style, igpStyleVar id, int var) { getStylei(style, id) = var; }
void    igpStyleSetVarVec2(igpStyle style, igpStyleVar id, igpVec2 var) { getStylev(style, id) = unwrapVec2(var); }


// Set style colors a Style or the current one if NULL
void igpStyleColorsAuto(igpStyle dest) {
	ImPlot::StyleColorsAuto(reinterpret_cast<ImPlotStyle *>(dest));
}
void igpStyleColorsClassic(igpStyle dest) {
	ImPlot::StyleColorsClassic(reinterpret_cast<ImPlotStyle *>(dest));
}
void igpStyleColorsDark(igpStyle dest) {
	ImPlot::StyleColorsDark(reinterpret_cast<ImPlotStyle *>(dest));
}
void igpStyleColorsLight(igpStyle dest) {
	ImPlot::StyleColorsLight(reinterpret_cast<ImPlotStyle *>(dest));
}


// Temporarily modify a style color
void igpPushStyleColor(igpStyleCol idx, igpVec4 col) { ImPlot::PushStyleColor(idx, unwrapVec4(col)); }
void igpPopStyleColor(int count) { ImPlot::PopStyleColor(count); }

void igpPushStyleVarFloat(igpStyleVar idx, float val) { ImPlot::PushStyleVar(idx, val); }
void igpPushStyleVarVec2(igpStyleVar idx, igpVec2 val) { ImPlot::PushStyleVar(idx, unwrapVec2(val)); }
void igpPushStyleVarInt(igpStyleVar idx, int val) { ImPlot::PushStyleVar(idx, val); }
void igpPopStyleVar(int count) { ImPlot::PopStyleVar(count); }


// Set styles for the next item only
void igpSetNextLineStyle(igpVec4 color, float weight) {
	ImPlot::SetNextLineStyle(unwrapVec4(color), weight);
}
void igpSetNextFillStyle(igpVec4 color, float alpha_mod) {
	ImPlot::SetNextFillStyle(unwrapVec4(color), alpha_mod);
}
void igpSetNextMarkerStyle(igpMarker marker, float size, igpVec4 fill, float weight, igpVec4 outline) {
	ImPlot::SetNextMarkerStyle(marker, size, unwrapVec4(fill), weight, unwrapVec4(outline));
}
void igpSetNextErrorBarStyle(igpVec4 color, float size, float weight) {
	ImPlot::SetNextErrorBarStyle(unwrapVec4(color), size, weight);
}

// Gets the last item primary color (i.e. its legend icon color)
igpVec4 igpGetLastItemColor() {
	return wrapVec4(ImPlot::GetLastItemColor());
}
