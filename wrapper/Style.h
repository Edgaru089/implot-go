#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

// Get the global style
igpStyle igpGetStyle();

// Access style variables
igpVec4 igpStyleGetColor(igpStyle style, igpStyleCol id);
void    igpStyleSetColor(igpStyle style, igpStyleCol id, igpVec4 color);

float   igpStyleGetVarFloat(igpStyle style, igpStyleVar id);
int     igpStyleGetVarInt(igpStyle style, igpStyleVar id);
igpVec2 igpStyleGetVarVec2(igpStyle style, igpStyleVar id);
void    igpStyleSetVarFloat(igpStyle style, igpStyleVar id, float var);
void    igpStyleSetVarInt(igpStyle style, igpStyleVar id, int var);
void    igpStyleSetVarVec2(igpStyle style, igpStyleVar id, igpVec2 var);


// Set style colors a Style or the current one if NULL
void igpStyleColorsAuto(igpStyle dest);
void igpStyleColorsClassic(igpStyle dest);
void igpStyleColorsDark(igpStyle dest);
void igpStyleColorsLight(igpStyle dest);

// Temporarily modify a style color
void igpPushStyleColor(igpStyleCol idx, igpVec4 col);
void igpPopStyleColor(int count);

void igpPushStyleVarFloat(igpStyleVar idx, float val);
void igpPushStyleVarVec2(igpStyleVar idx, igpVec2 val);
void igpPushStyleVarInt(igpStyleVar idx, int val);
void igpPopStyleVar(int count);

// Set styles for the next item only
void igpSetNextLineStyle(igpVec4 color, float weight);
void igpSetNextFillStyle(igpVec4 color, float alpha_mod);
void igpSetNextMarkerStyle(igpMarker marker, float size, igpVec4 fill, float weight, igpVec4 outline);
void igpSetNextErrorBarStyle(igpVec4 color, float size, float weight);

// Gets the last item primary color (i.e. its legend icon color)
igpVec4 igpGetLastItemColor();

// GetStyleColorName & GetMarkerName implemented in Go


#ifdef __cplusplus
}
#endif
