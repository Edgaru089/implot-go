#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif


// implot.PlotLine() [Plot.go]
void igpPlotLine(const char *label, const double *values, int count, double xscale, double x0);
void igpPlotLineXY(const char *label, const double *xs, const double *ys, int count, int stride);

// implot.PlotScatter() [Plot.go]
void igpPlotScatter(const char *label, const double *values, int count, double xscale, double x0);
void igpPlotScatterXY(const char *label, const double *xs, const double *ys, int count, int stride);

// implot.PlotStairs() [Plot.go]
void igpPlotStairs(const char *label, const double *values, int count, double xscale, double x0);
void igpPlotStairsXY(const char *label, const double *xs, const double *ys, int count, int stride);

// implot.PlotShadedRef() [Plot.go]
void igpPlotShadedRef(const char *label, const double *values, int count, double yref, double xscale, double x0);
void igpPlotShadedRefXY(const char *label, const double *xs, const double *ys, double yref, int count, int stride);
// implot.PlotShadedLines() [Plot.go]
void igpPlotShadedLinesXY(const char *label, const double *xs, const double *ys1, const double *ys2, int count, int stride);

// Vertical/Horizontal bars
void igpPlotBars(const char *label, const double *values, int count, double bar_width, double x0);
void igpPlotBarsXY(const char *label, const double *xs, const double *ys, int count, double bar_width, int stride);
void igpPlotBarsH(const char *label, const double *values, int count, double bar_height, double y0);
void igpPlotBarsHXY(const char *label, const double *xs, const double *ys, int count, double bar_height, int stride);
// Groups of vertical/horizontal bars
void igpPlotBarGroups(const char **labels, const double *values, int items_per_group, int groups, double group_width, double x0, igpBarGroupsFlags flags);
void igpPlotBarGroupsH(const char **labels, const double *values, int items_per_group, int groups, double group_height, double y0, igpBarGroupsFlags flags);


#ifdef __cplusplus
}
#endif
