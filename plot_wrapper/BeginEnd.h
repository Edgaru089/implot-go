#pragma once

#include <stdbool.h>
#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif


// implot.BeginPlot() [BeginEnd.go]
bool igpBeginPlot(const char *title, igpVec2 size, igpFlags flags);

// implot.EndPlot() [BeginEnd.go]
void igpEndPlot();

// implot.BeginSubplots() [BeginEnd.go]
bool igpBeginSubplots(const char *title, int rows, int cols, igpVec2 size, igpSubplotFlags flags, float *row_ratios, float *col_ratios);

// implot.EndSubplots() [BeginEnd.go]
void igpEndSubplots();


#ifdef __cplusplus
}
#endif
