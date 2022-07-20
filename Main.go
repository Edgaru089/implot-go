package implot

// #include "wrapper/Main.h"
import "C"

// PlotVersion returns a version string for ImPlot, e.g., "0.13 WIP".
func PlotVersion() string {
	return C.GoString(C.igpMain_Version())
}

// ShowPlotDemoWindow shows the ImPlot demo window.
func ShowPlotDemoWindow(open *bool) {
	if open == nil {
		C.igpShowDemoWindow(nil)
	} else {
		var copen C.bool
		C.igpShowDemoWindow(&copen)
		*open = bool(copen)
	}
}

//export igpPanic
func igpPanic(msg *C.char) {
	panic(C.GoString(msg))
}
