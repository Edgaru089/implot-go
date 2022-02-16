
#include "Plot.h"
#include "ImPlot.hpp"


void igpPlotLine(const char *label, const double *values, int count, double xscale, double x0) {
	ImPlot::PlotLine<double>(label, values, count, xscale, x0, 0, sizeof(double));
}
void igpPlotLineXY(const char *label, const double *xs, const double *ys, int count, int stride) {
	ImPlot::PlotLine<double>(label, xs, ys, count, 0, sizeof(double));
}
