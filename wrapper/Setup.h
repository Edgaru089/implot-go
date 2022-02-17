#pragma once

#include <stdint.h>
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


#ifdef __cplusplus
}
#endif
