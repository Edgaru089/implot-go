package implot

// #include "wrapper/Main.h"
import "C"

// Version returns a version string, e.g., "0.13 WIP".
func Version() string {
	return C.GoString(C.igpMain_Version())
}

// ShowDemoWindow shows the ImPlot demo window.
func ShowDemoWindow(open *bool) {
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
