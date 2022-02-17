
#include "Plot.h"
#include "ImPlot.hpp"


void igpPlotLine(const char *label, const double *values, int count, double xscale, double x0) {
	ImPlot::PlotLine<double>(label, values, count, xscale, x0, 0);
}
void igpPlotLineXY(const char *label, const double *xs, const double *ys, int count, int stride) {
	ImPlot::PlotLine<double>(label, xs, ys, count, 0, stride);
}

void igpPlotScatter(const char *label, const double *values, int count, double xscale, double x0) {
	ImPlot::PlotScatter<double>(label, values, count, xscale, x0, 0);
}
void igpPlotScatterXY(const char *label, const double *xs, const double *ys, int count, int stride) {
	ImPlot::PlotScatter<double>(label, xs, ys, count, 0, stride);
}

void igpPlotStairs(const char *label, const double *values, int count, double xscale, double x0) {
	ImPlot::PlotStairs<double>(label, values, count, xscale, x0, 0);
}
void igpPlotStairsXY(const char *label, const double *xs, const double *ys, int count, int stride) {
	ImPlot::PlotStairs<double>(label, xs, ys, count, 0, stride);
}

void igpPlotShadedRef(const char *label, const double *values, int count, double yref, double xscale, double x0) {
	ImPlot::PlotShaded<double>(label, values, count, yref, xscale, x0, 0);
}
void igpPlotShadedRefXY(const char *label, const double *xs, const double *ys, double yref, int count, int stride) {
	ImPlot::PlotShaded<double>(label, xs, ys, count, yref, 0, stride);
}
void igpPlotShadedLinesXY(const char *label, const double *xs, const double *ys1, const double *ys2, int count, int stride) {
	ImPlot::PlotShaded(label, xs, ys1, ys2, count, 0, stride);
}
