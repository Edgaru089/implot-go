#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif


// implot.PlotLine() [Plot.go]
void igpPlotLine(const char *label, const double *values, int count, double xscale, double x0);
void igpPlotLineXY(const char *label, const double *xs, const double *ys, int count, int stride);


#ifdef __cplusplus
}
#endif
