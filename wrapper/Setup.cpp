
#include "Setup.h"
#include "ImPlot.hpp"


void igpSetupAxis(igpAxis axis, const char *label, igpAxisFlags flags) {
	ImPlot::SetupAxis(axis, label, flags);
}

void igpSetupAxisLimits(igpAxis axis, double vmin, double vmax, igpCondition cond) {
	ImPlot::SetupAxisLimits(axis, vmin, vmax, cond);
}

void igpSetupAxisFormat(igpAxis axis, const char *fmt) {
	ImPlot::SetupAxisFormat(axis, fmt);
}

extern "C" void igpgoAxisFormatCb(double value, char *buff, int size, void *userData);
void            igpSetupAxisFormatCallback(igpAxis axis, uintptr_t handle) {
    ImPlot::SetupAxisFormat(axis, &igpgoAxisFormatCb, reinterpret_cast<void *>(handle));
}

void igpSetupAxisTickValues(igpAxis axis, const double *values, int n, const char **labels, bool keepDefaults) {
	ImPlot::SetupAxisTicks(axis, values, n, labels, keepDefaults);
}
void igpSetupAxisTickRange(igpAxis axis, double vmin, double vmax, int n, const char **labels, bool keepDefaults) {
	ImPlot::SetupAxisTicks(axis, vmin, vmax, n, labels, keepDefaults);
}

void igpSetupAxes(const char *xlabel, const char *ylabel, igpAxisFlags xflags, igpAxisFlags yflags) {
	ImPlot::SetupAxes(xlabel, ylabel, xflags, yflags);
}
void igpSetupAxesLimits(double xmin, double xmax, double ymin, double ymax, igpCondition cond) {
	ImPlot::SetupAxesLimits(xmin, xmax, ymin, ymax, cond);
}

void igpSetupLegend(igpLocation location, igpLegendFlags flags) {
	ImPlot::SetupLegend(location, flags);
}
void igpSetupMouseText(igpLocation location, igpMouseTextFlags flags) {
	ImPlot::SetupMouseText(location, flags);
}

void igpSetupFinish() {
	ImPlot::SetupFinish();
}
