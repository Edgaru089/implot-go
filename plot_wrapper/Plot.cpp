
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

void igpPlotBars(const char *label, const double *values, int count, double bar_width, double x0) {
	ImPlot::PlotBars(label, values, count, bar_width, x0, 0);
}
void igpPlotBarsXY(const char *label, const double *xs, const double *ys, int count, double bar_width, int stride) {
	ImPlot::PlotBars(label, xs, ys, count, bar_width, stride);
}
void igpPlotBarsH(const char *label, const double *values, int count, double bar_height, double y0) {
	ImPlot::PlotBarsH(label, values, count, bar_height, y0, 0);
}
void igpPlotBarsHXY(const char *label, const double *xs, const double *ys, int count, double bar_height, int stride) {
	ImPlot::PlotBarsH(label, xs, ys, count, bar_height, 0, stride);
}

void igpPlotBarGroups(const char **labels, const double *values, int items_per_group, int groups, double group_width, double x0, igpBarGroupsFlags flags) {
	ImPlot::PlotBarGroups(labels, values, items_per_group, groups, group_width, x0, flags);
}
void igpPlotBarGroupsH(const char **labels, const double *values, int items_per_group, int groups, double group_height, double y0, igpBarGroupsFlags flags) {
	ImPlot::PlotBarGroupsH(labels, values, items_per_group, groups, group_height, y0, flags);
}
