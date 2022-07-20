#pragma once

#ifdef __cplusplus
extern "C" {
#endif

typedef int igpAxis;
typedef int igpFlags;
typedef int igpAxisFlags;
typedef int igpSubplotFlags;
typedef int igpLegendFlags;
typedef int igpMouseTextFlags;
typedef int igpDragToolFlags;
typedef int igpBarGroupsFlags;

typedef int igpCondition;
typedef int igpStyleCol;
typedef int igpStyleVar;
typedef int igpMarker;
typedef int igpColormap;
typedef int igpLocation;
typedef int igpBin;

typedef void *igpContext, *iggContext, *igpStyle;

typedef struct {
	float x, y;
} igpVec2;

typedef struct {
	float x, y, z, w;
} igpVec4;

typedef struct {
	double x, y;
} igpPoint;


#ifdef __cplusplus
}
#endif
