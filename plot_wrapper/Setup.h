#pragma once

#include <stdint.h>
#include <stdbool.h>
#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif


// implot.SetupAxis() [Setup.go]
void igpSetupAxis(igpAxis axis, const char *label, igpAxisFlags flags);
// implot.SetupAxisLimits() [Setup.go]
void igpSetupAxisLimits(igpAxis axis, double vmin, double vmax, igpCondition cond);
// implot.SetupAxisLinks() [Setup.go]
void igpSetupLinks(igpAxis axis); // TODO implot.SetupAxisLinks
// implot.SetupAxisFormat() [Setup.go]
void igpSetupAxisFormat(igpAxis axis, const char *fmt);
// implot.SetupAxisFormatCallback() [Setup.go]
void igpSetupAxisFormatCallback(igpAxis axis, uintptr_t handle); // TODO implot.SetupAxisFormatCallback
// implot.SetupAxisTickValues() [Setup.go]
void igpSetupAxisTickValues(igpAxis axis, const double *values, int n, const char **labels, bool keepDefaults);
// implot.SetupAxisTickRange() [Setup.go]
void igpSetupAxisTickRange(igpAxis axis, double vmin, double vmax, int n, const char **labels, bool keepDefaults);

void igpSetupAxes(const char *xlabel, const char *ylabel, igpAxisFlags xflags, igpAxisFlags yflags);
void igpSetupAxesLimits(double xmin, double xmax, double ymin, double ymax, igpCondition cond);

void igpSetupLegend(igpLocation location, igpLegendFlags flags);
void igpSetupMouseText(igpLocation location, igpMouseTextFlags flags);

void igpSetupFinish();


#ifdef __cplusplus
}
#endif
