
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
