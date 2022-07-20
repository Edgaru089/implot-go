
#include "Main.h"
#include "ImPlot.hpp"


const char *igpMain_Version() {
	return IMPLOT_VERSION;
}

void igpShowDemoWindow(bool *open) {
	ImPlot::ShowDemoWindow(open);
}
