
#include "BeginEnd.h"
#include "ImPlot.hpp"


// implot.BeginPlot() [BeginEnd.go]
bool igpBeginPlot(const char *title, igpVec2 size, igpFlags flags) {
	return ImPlot::BeginPlot(title, Vec2(size), flags);
}

// implot.EndPlot() [BeginEnd.go]
void igpEndPlot() {
	ImPlot::EndPlot();
}

// implot.BeginSubplots() [BeginEnd.go]
bool igpBeginSubplots(const char *title, int rows, int cols, igpVec2 size, igpSubplotFlags flags, float *row_ratios, float *col_ratios) {
	return ImPlot::BeginSubplots(title, rows, cols, Vec2(size), flags, row_ratios, col_ratios);
}

// implot.EndSubplots() [BeginEnd.go]
void igpEndSubplots() {
	ImPlot::EndSubplots();
}
