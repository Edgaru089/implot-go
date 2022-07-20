
#include "Contexts.h"
#include "ImPlot.hpp"
#include "Types.h"


igpContext igpCreateContext() {
	return reinterpret_cast<igpContext>(ImPlot::CreateContext());
}

void igpDestroyContext(igpContext ctx) {
	ImPlot::DestroyContext(reinterpret_cast<ImPlotContext *>(ctx));
}

// implot.CurrentContext() [Contexts.go]
igpContext igpCurrentContext() {
	return reinterpret_cast<igpContext>(ImPlot::GetCurrentContext());
}

// implot.SetCurrentContext() [Contexts.go]
void igpSetCurrentContext(igpContext ctx) {
	ImPlot::SetCurrentContext(reinterpret_cast<ImPlotContext *>(ctx));
}

// implot.SetImGUIContext() [Contexts.go]
void igpSetImGUIContext(iggContext ctx) {
	ImPlot::SetImGuiContext(reinterpret_cast<ImGuiContext *>(ctx));
}
